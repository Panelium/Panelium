package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"io"
	"os"
	"panelium/daemon/internal/db"
	"panelium/daemon/internal/docker"
	"panelium/daemon/internal/model"
	"panelium/proto_gen_go/daemon"
	"path"
)

// TODO: implement storage limiting

func Install(s *model.Server) error {
	blueprint := model.Blueprint{}
	tx := db.Instance().First(&blueprint, "s.BID = ?", s.BID)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return fmt.Errorf("failed to find blueprint with ID %s: %w", s.BID, tx.Error)
	}

	// pull setup script docker image
	rc, err := docker.Instance().ImagePull(context.Background(), blueprint.SetupDockerImage, image.PullOptions{})
	if err != nil {
		return fmt.Errorf("failed to pull setup script docker image %s: %w", blueprint.SetupDockerImage, err)
	}

	_, err = io.Copy(io.Discard, rc) // we could get the progress of the image pull here
	if err != nil {
		return fmt.Errorf("failed to read image pull response: %w", err)
	}
	err = rc.Close()
	if err != nil {
		return fmt.Errorf("failed to close image pull response: %w", err)
	}

	vl, err := docker.Instance().VolumeList(context.Background(), volume.ListOptions{
		Filters: filters.NewArgs(filters.Arg("name", s.SID)),
	})
	if err != nil {
		return fmt.Errorf("failed to list volumes: %w", err)
	}

	var vol *volume.Volume

	if len(vl.Volumes) == 0 {
		v, err := docker.Instance().VolumeCreate(context.Background(), volume.CreateOptions{
			Name:   s.SID,
			Driver: "local",
		})
		if err != nil {
			return fmt.Errorf("failed to create volume for server %s: %w", s.SID, err)
		}
		vol = &v
	} else if len(vl.Volumes) > 1 {
		return fmt.Errorf("found multiple volumes with name %s, expected only one", s.SID)
	} else if len(vl.Volumes) == 1 {
		vol = vl.Volumes[0]
		fmt.Printf("found existing volume for server %s: %s\n", s.SID, vol.Name)
	}

	setupScript, err := base64.StdEncoding.DecodeString(blueprint.SetupScriptBase64)
	if err != nil {
		return fmt.Errorf("failed to decode setup script: %w", err)
	}

	err = os.WriteFile(path.Join(vol.Mountpoint, "install"), setupScript, 0777)
	if err != nil {
		return fmt.Errorf("failed to write setup script to volume: %w", err)
	}

	if s.ContainerExists {
		err = docker.Instance().ContainerRemove(context.Background(), s.SID, container.RemoveOptions{
			Force: true,
		})
		if err != nil {
			return fmt.Errorf("failed to remove existing container for server %s: %w", s.SID, err)
		}
		s.ContainerExists = false
		if err := db.Instance().Save(s).Error; err != nil {
			return err
		}
	}

	resources := container.Resources{
		Memory:            int64(s.ResourceLimit.RAM * 1024 * 1024),
		MemoryReservation: int64(s.ResourceLimit.RAM * 1024 * 1024),
		MemorySwap:        int64(s.ResourceLimit.RAM*1024*1024 + s.ResourceLimit.SWAP*1024*1024),
		OomKillDisable:    &[]bool{true}[0], // dumb hack to not have to create a variable for this
	}

	if s.ResourceLimit.CPU > 0 {
		resources.CPUQuota = int64(s.ResourceLimit.CPU * 1_000)
		resources.CPUPeriod = 100_000
		resources.CPUShares = 1024
	}

	// create setup script container
	scr, err := docker.Instance().ContainerCreate(context.Background(), &container.Config{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		OpenStdin:    true,
		Tty:          true,
		Image:        blueprint.SetupDockerImage,
		WorkingDir:   "/data",
		Cmd: []string{
			blueprint.SetupScriptInterpreter,
			"./install",
		},
		Env: []string{
			"SERVER_BINARY=" + blueprint.ServerBinary,
		},
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:     mount.TypeBind,
				Source:   vol.Mountpoint,
				Target:   "/data",
				ReadOnly: false,
			},
		},
		Resources: resources,
	}, &network.NetworkingConfig{}, &v1.Platform{}, s.SID)
	if err != nil {
		return fmt.Errorf("failed to create setup script container: %w", err)
	}

	if err := docker.Instance().ContainerStart(context.Background(), scr.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("failed to start setup script container: %w", err)
	}

	fmt.Printf("setup script container started with ID: %s\n", scr.ID)

	// wait for the setup script container to finish install
	statusCh, errCh := docker.Instance().ContainerWait(context.Background(), scr.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case status := <-statusCh:
		if status.StatusCode != 0 {
			return fmt.Errorf("setup script container exited with status code %d", status.StatusCode)
		}
		fmt.Printf("setup script container finished with status code %d\n", status.StatusCode)

		// remove the setup script container
		if err := docker.Instance().ContainerRemove(context.Background(), scr.ID, container.RemoveOptions{
			Force: true,
		}); err != nil {
			return fmt.Errorf("failed to remove setup script container: %w", err)
		}
	}

	var ports nat.PortSet

	for _, alloc := range s.Allocations {
		if alloc.Port < 1024 || alloc.Port > 65535 {
			return fmt.Errorf("port %d is out of range (1024-65535)", alloc.Port)
		}
		//open both tcp and udp ports
		ports[nat.Port(fmt.Sprintf("%d/tcp", alloc.Port))] = struct{}{}
		ports[nat.Port(fmt.Sprintf("%d/udp", alloc.Port))] = struct{}{}
	}

	// create the server container
	_, err = docker.Instance().ContainerCreate(context.Background(), &container.Config{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		OpenStdin:    true,
		Tty:          true,
		Image:        s.DockerImage,
		WorkingDir:   "/data",
		Cmd:          []string{blueprint.ServerBinary},
		Env:          []string{"SERVER_BINARY=" + blueprint.ServerBinary},
		ExposedPorts: ports,
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:     mount.TypeBind,
				Source:   vol.Mountpoint,
				Target:   "/data",
				ReadOnly: false,
			},
		},
		Resources: resources,
	}, &network.NetworkingConfig{}, &v1.Platform{}, s.SID)
	if err != nil {
		return fmt.Errorf("failed to create server container: %w", err)
	}

	s.ContainerExists = true
	s.OfflineReason = daemon.ServerOfflineReason_SERVER_OFFLINE_REASON_CREATED
	if err := db.Instance().Save(s).Error; err != nil {
		return err
	}

	return nil
}

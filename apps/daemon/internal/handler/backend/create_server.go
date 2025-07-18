package backend

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"log"
	"panelium/daemon/internal/model"
	"panelium/daemon/internal/server"
	"panelium/proto_gen_go"
	"panelium/proto_gen_go/daemon"
)

func (s *BackendServiceHandler) CreateServer(
	ctx context.Context,
	req *connect.Request[daemon.Server],
) (*connect.Response[proto_gen_go.SuccessMessage], error) {
	allocations := make([]model.ServerAllocation, len(req.Msg.Allocations))
	for i, alloc := range req.Msg.Allocations {
		if alloc.Port < 1024 || alloc.Port > 65535 {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("port %d is out of range (1024-65535)", alloc.Port))
		}

		allocations[i] = model.ServerAllocation{
			IP:   alloc.Ip,
			Port: uint16(alloc.Port),
		}
	}

	resourceLimit := model.ResourceLimit{
		CPU:     req.Msg.ResourceLimit.Cpu,
		RAM:     req.Msg.ResourceLimit.Ram,
		SWAP:    req.Msg.ResourceLimit.Swap,
		Storage: req.Msg.ResourceLimit.Storage,
	}

	_, err := server.CreateServer(req.Msg.Sid, req.Msg.OwnerId, req.Msg.UserIds, allocations, resourceLimit, req.Msg.DockerImage, req.Msg.Bid)
	if err != nil {
		log.Printf("Failed to create server: %v", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to create server"))
	}

	res := connect.NewResponse(&proto_gen_go.SuccessMessage{
		Success: true,
	})

	return res, nil
}

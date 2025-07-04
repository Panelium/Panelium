import React from "react";
import {
  FolderOpen,
  type LucideIcon,
  Play,
  Settings,
  Square,
  Terminal,
  Users,
} from "lucide-react";
import { Link, useNavigate } from "react-router-dom";
import { ServerStatusType } from "proto-gen-ts/daemon/Server_pb";
import { cn, formatMemory } from "~/lib/utils";

import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "~/components/ui/card";
import ServerBar from "~/components/bars/ServerBar";
import EntityAvatar from "~/components/avatars/EntityAvatar";
import StatusBadge from "~/components/dashboard/StatusBadge";

//temp until moved to proto
export interface Server {
  id: string;
  name: string;
  status: ServerStatusType;
  description: string;
  icon?: string;
  cpuUsage: number;
  memoryUsage: {
    used: number;
    total: number;
  };
  diskUsage: {
    used: number;
    total: number;
  };
  game: string;
  players?: {
    online: number;
    max: number;
  };
  ip?: string;
  port?: number; //TODO: ip and port will be merged into networkAllocation
  location?: string;
  node?: string;
  console: {time: string, content: string}[]
}

interface ServerCardProps {
  server: Server;
  className?: string;
}

interface ServerCardLinkProps {
  icon: LucideIcon;
  color?: string;
  bgColor?: string;
  className?: string;
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
}

interface ServerCardActionProps {
  server: Server;
  onAction: (e: React.MouseEvent, action: string) => void;
}

const serverCardTransition = "transition-all duration-300 ease-in-out truncate";

const ServerCardButton: React.FC<ServerCardLinkProps> = ({
  icon,
  color,
  bgColor,
  className,
  onClick,
}) => {
  const IconComponent = icon;
  return (
    <div>
    <div className={cn(
      "scale-20 group-hover:scale-100 w-10 h-10 bg-card rounded-full absolute",
      serverCardTransition
    )}/>
    <button
      className={cn(
        "flex items-center justify-center h-10 w-10 rounded-full shadow-sm",
        "scale-20 group-hover:scale-100",
        "cursor-pointer",
        serverCardTransition,
        className
      )}
      style={{ color: color, backgroundColor: bgColor }}
      onClick={onClick}
    >
      <IconComponent className="w-5 h-5" />
    </button>
    </div>
  );
};

const ServerCardHeader: React.FC<{ server: Server }> = ({ server }) => {
  return (
    <CardHeader className="gap-3">
      <div className="flex items-start justify-between gap-3 truncate">
        <EntityAvatar
          src={server.icon}
          alt={server.name}
          title={server.name}
          subTitle={server.game}
          className={serverCardTransition}
        />
        <StatusBadge status={server.status} />
      </div>
      <p className="text-sm text-server-card-foreground line-clamp-2">
        {server.description}
      </p>
    </CardHeader>
  );
};

const ServerCardContent: React.FC<{ server: Server }> = ({ server }) => {
  return (
    <CardContent className="space-y-3">
      <ServerBar
        title="CPU"
        uiValue={server.cpuUsage.toFixed(1) + "%"}
        value={server.cpuUsage}
        max={100}
      />
      <ServerBar
        title="Memory"
        uiValue={
          formatMemory(server.memoryUsage.used) +
          " / " +
          formatMemory(server.memoryUsage.total)
        }
        value={server.memoryUsage.used}
        max={server.memoryUsage.total}
      />
    </CardContent>
  );
};

const ServerCardFooter: React.FC<{ server: Server }> = ({ server }) => {
  return (
    <CardFooter>
      <div className="flex flex-1 items-center justify-between text-xs text-card-muted-foreground">
        {server.ip && (
          <div className="font-mono">
            {server.ip}
            {server.port && ":" + server.port}
          </div>
        )}

        <div className="flex flex-row gap-1">
          {server.players && (
            <>
              <Users className="h-4 w-4 text-card-muted-foreground" />
              <span>
                {server.players.online} / {server.players.max} Players
              </span>
            </>
          )}
        </div>
      </div>
    </CardFooter>
  );
};

const ServerCardAction: React.FC<ServerCardActionProps> = ({
  server,
  onAction,
}) => {
  return (
    <div
      className={cn(
        "flex flex-col items-center justify-center gap-3",
        "border-l border-tag-purple/30 bg-tag-purple-background/20",
        "absolute top-0 bottom-0 right-0 w-[60px] translate-x-[60px] ",
        "group-hover:w-[60px] group-hover:translate-x-[0px]",
        serverCardTransition
      )}
    >
      <ServerCardButton
        icon={server.status === ServerStatusType.ONLINE ? Square : Play}
        className={cn(
          "bg-tag-green-background/30 text-tag-green hover:bg-tag-green-background/70"
        )}
        onClick={(e) => onAction(e, "power")}
      />
      <ServerCardButton
        icon={Terminal}
        className={cn(
          "bg-tag-purple-background/30 text-tag-purple hover:bg-tag-purple-background/70"
        )}
        onClick={(e) => onAction(e, "console")}
      />
      <ServerCardButton
        icon={FolderOpen}
        className={cn(
          "bg-tag-orange-background/30 text-tag-orange hover:bg-tag-orange-background/70"
        )}
        onClick={(e) => onAction(e, "files")}
      />
      <ServerCardButton
        icon={Settings}
        className={cn(
          "bg-tag-gray-background/30 text-tag-gray hover:bg-tag-gray-background/70"
        )}
        onClick={(e) => onAction(e, "settings")}
      />
    </div>
  );
};

const ServerCard: React.FC<ServerCardProps> = ({ server, className }) => {
  const navigate = useNavigate();

  const handleServerCardButtonClick = (e: React.MouseEvent, action: string) => {
    e.preventDefault();
    e.stopPropagation();

    switch (action) {
      case "power":
        navigate(`/server/${server.id}/power`);
        break;
      case "console":
        navigate(`/server/${server.id}/console`);
        break;
      case "files":
        navigate(`/server/${server.id}/files`);
        break;
      case "settings":
        navigate(`/server/${server.id}/settings`);
        break;
      default:
        navigate(`/server/${server.id}`);
    }
  };

  return (
    <Link to={`/server/${server.id}`} className="group">
      <Card
        className={cn(
          "hover:rounded-h-xl relative flex flex-row overflow-hidden shadow-md shadow-black/20",
          "bg-server-card border-border hover:border-border-hover",
          "hover:scale-102 hover:shadow-xl hover:z-2 hover:shadow-black/80",
          serverCardTransition,
          className
        )}
      >
        <div
          className={cn(
            "flex flex-col w-full gap-3",
            "group-hover:w-[calc(100%-60px)]",
            serverCardTransition
          )}
        >
          <ServerCardHeader server={server} />
          <ServerCardContent server={server} />
          <ServerCardFooter server={server} />
        </div>
        <ServerCardAction
          server={server}
          onAction={(e, action) => handleServerCardButtonClick(e, action)}
        />
      </Card>
    </Link>
  );
};

export default ServerCard;

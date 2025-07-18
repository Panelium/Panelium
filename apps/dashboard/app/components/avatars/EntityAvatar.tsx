import React from "react";
import { cn } from "~/lib/utils";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";

interface EntityAvatarProps {
  title: string;
  subTitle: string;
  src?: string;
  alt?: string;
  className?: string;
}

const EntityAvatar: React.FC<EntityAvatarProps> = ({ src, alt, title, subTitle, className }) => {
  return (
    <div className="max-w-[100%] flex items-center text-left gap-3 truncate group/entity-avatar">
      <Avatar>
        <AvatarImage src={src} alt={alt} className="h-full w-full object-cover" />
        <AvatarFallback>{title.charAt(0).toUpperCase()}</AvatarFallback>
      </Avatar>
      <div className="w-1/2 flex flex-col flex-1">
        <h3
          className={cn(
            "font-medium truncate",
            "text-foreground",
            "group-hover/entity-avatar:text-indigo-600 dark:group-hover/entity-avatar:text-indigo-400",
            className
          )}
        >
          {title}
        </h3>
        <p className="text-xs text-muted-foreground truncate">{subTitle}</p>
      </div>
    </div>
  );
};

export default EntityAvatar;

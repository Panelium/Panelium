import { cn } from "~/lib/utils";

interface BarProps {
  percentage: number;
  barColor?: string;
  size: "sm" | "md" | "lg";
  className?: string;
}

const Bar: React.FC<BarProps> = ({
  percentage,
  barColor = "bg-emerald-500",
  className,
  size = "md",
}) => {
  const sizeClasses = {
    sm: "h-1",
    md: "h-1.5",
    lg: "h-2.5",
  };

  return (
    <div
      className={cn(
        sizeClasses[size],
        "w-full overflow-hidden rounded-full",
        "bg-chart-background",
        className
      )}
    >
      <div
        className={cn("h-full rounded-full", barColor)}
        style={{ width: `${percentage}%` }}
      />
    </div>
  );
};

export default Bar;

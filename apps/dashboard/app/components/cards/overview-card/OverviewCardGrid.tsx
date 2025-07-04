import React from "react";
import OverviewCard, { type OverviewCardProps } from "./OverviewCard";

const OverviewCardGrid: React.FC<{ cards: OverviewCardProps[] }> = ({
  cards,
}) => {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mb-6">
      {cards.map((props, idx) => (
        <OverviewCard key={idx} {...props} />
      ))}
    </div>
  );
};

export default OverviewCardGrid;

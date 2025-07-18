import type { Route } from "../../.react-router/types/app/+types/root";

export const meta: Route.MetaFunction = () => [
  { charset: "utf-8" },
  { title: "Panelium" },
  {
    name: "description",
    content:
      "Panelium is an all-in-one server hosting solution that provides a web-based control panel for managing games, servers, applications, and services.",
  },
  { name: "viewport", content: "width=device-width, initial-scale=1" },
  { name: "theme-color", content: "#ffffff" },
];

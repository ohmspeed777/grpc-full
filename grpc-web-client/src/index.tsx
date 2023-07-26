import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Route from "./pages/Route";
import FoodJson from "./pages/json/FoodJson";
import FoodGRPC from "./pages/grpc/FoodGRPC";
import FoodStreamGRPC from "./pages/grpc/FoodStreamGRPC"

const router = createBrowserRouter([
  {
    path: "/",
    element: <Route />,
    children: [
      {
        index: true,
        element: <FoodJson />,
      },
      {
        path: "grpc/food",
        element: <FoodGRPC />,
      },
      {
        path: "grpc/stream/food",
        element: <FoodStreamGRPC />,
      },
      {
        path: "json/food",
        element: <FoodJson />,
      },
    ],
  },
]);

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

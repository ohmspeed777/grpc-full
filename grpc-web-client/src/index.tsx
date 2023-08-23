import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Route from "./pages/Route";
import FoodJson from "./pages/json/FoodJson";
import FoodGRPC from "./pages/grpc/FoodGRPC";
import SignInGRPC from "./pages/grpc/SignInGRPC";
import FoodStreamGRPC from "./pages/grpc/FoodStreamGRPC";
import OrderGRPC from "./pages/grpc/Order";

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
      {
        path: "grpc/sign-in",
        element: <SignInGRPC />,
      },
      {
        path: "/grpc/auth/orders",
        element: <OrderGRPC />,
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

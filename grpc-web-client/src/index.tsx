import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import FoodJson from "./pages/json/FoodJson";
import FoodGRPC from "./pages/grpc/FoodGRPC";

const router = createBrowserRouter([
  {
    path: "/",
    element: <FoodJson />,
  },
  // {
  //   path: "/json/food",
  //   element: <FoodsJson />,
  // },
  {
    path: "/grpc/food",
    element: <FoodGRPC />,
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

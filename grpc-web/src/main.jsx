import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import "./index.css";
import FoodsJson from "./pages/FoodJson";
import FoodGRPC from "./pages/grpc/FoodGRPC";

const router = createBrowserRouter([
  {
    path: "/",
    element: <FoodsJson />,
  },
  {
    path: "/json/food",
    element: <FoodsJson />,
  },
  {
    path: "/grpc/food",
    element: <FoodGRPC />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

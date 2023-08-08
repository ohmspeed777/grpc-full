import React from "react";
import { Link, Outlet } from "react-router-dom";

const Route = () => {
  return (
    <div>
      <div className="navbar bg-base-100 px-6">
        <div className="flex-1">
          <a className="btn btn-ghost normal-case text-xl">daisyUI</a>
        </div>
        <div className="flex-none">
          <ul className="menu menu-horizontal px-1">
            <li>
              <Link to={"/json/food"}>Food JSON</Link>
            </li>
            <li>
              <Link to={"/grpc/food"}>Food GRPC</Link>
            </li>
            <li>
              <Link to={"/grpc/stream/food"}>Food Stream GRPC</Link>
            </li>
            <li>
              <Link to={"/grpc/sign-in"}>Sign In GRPC</Link>
            </li>
            {/* <li>
              <details>
                <summary>Parent</summary>
                <ul className="p-2 bg-base-100">
                  <li>
                    <a>Link 1</a>
                  </li>
                  <li>
                    <a>Link 2</a>
                  </li>
                </ul>
              </details>
            </li> */}
          </ul>
        </div>
      </div>
      <Outlet />
    </div>
  );
};

export default Route;

import React, { useEffect, useState } from "react";
import { SignInReq } from "../../protos/users_pb";
import instants from "../../instants/instants";

const SignInGRPC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const onSubmit = () => {
    if (email && password) {
      callGRPC();
    }
  };

  const callGRPC = async () => {
    const req = new SignInReq();
    req.setEmail(email);
    req.setPassword(password);

    try {
      const resp = await instants.userGRPC.signIn(req, null);

      console.log("token: ", resp.getToken());
      localStorage.setItem("token", resp.getToken());
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div>
      <div className="container max-w-2xl mx-auto pt-12">
        <h1 className="text-center text-secondary text-5xl mb-6">
          Sign in GRPC
        </h1>
        <div className="form-control w-full mb-4">
          <label className="label">
            <span className="label-text">Email</span>
          </label>
          <input
            type="text"
            placeholder="email"
            value={email}
            className="input input-bordered w-full "
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div className="form-control w-full">
          <label className="label">
            <span className="label-text">Password</span>
          </label>
          <input
            type="password"
            placeholder="password"
            className="input input-bordered w-full "
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <button className="btn btn-primary mt-6 text-center" onClick={onSubmit}>
          Submit
        </button>
      </div>
    </div>
  );
};

export default SignInGRPC;

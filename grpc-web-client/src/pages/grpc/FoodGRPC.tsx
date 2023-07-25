import Card from "../../components/card";
import configs from "../../configs/configs";
import React, { useEffect, useState } from "react";
import instants from "../../instants/instants";
import { GetAllRequest } from "../../protos/foods_pb";
import { Query } from "../../protos/common_pb";
import { async } from "q";

const FoodsGRPC = () => {
  const [foods, setFoods] = useState([]);

  useEffect(() => {
    (async() => {
      await callGRPC()
    })()
  }, []);

  const callGRPC = async () => {
    const q = new Query()
    q.setPage(1)
    q.setLimit(10)

    const req = new GetAllRequest();
    req.setQuery(q);
    const resp = await instants.foodGRPC.getAll(req, null);
    // console.log(resp.getEntitiesList().map(el => el.));
  };

  return (
    <div className="container max-w-2xl mx-auto pt-12">
      <h1 className="text-center">Food Json Format HTTP</h1>
      <div className="grid grid-cols-2 mt-6">{foods.map((el) => Card(el))}</div>
    </div>
  );
};

export default FoodsGRPC;

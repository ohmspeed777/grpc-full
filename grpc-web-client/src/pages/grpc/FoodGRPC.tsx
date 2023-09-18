import Card from "../../components/Card";
import React, { useEffect, useState } from "react";
import instants from "../../instants/instants";
import { GetAllRequest } from "../../protos/foods_pb";
import { Query } from "../../protos/common_pb";
import { IFood } from "../../types/food";

const FoodsGRPC = () => {
  const [foods, setFoods] = useState<IFood[]>([]);
  useEffect(() => {
    (async () => {
      await callGRPC();
    })();
  }, []);

  const callGRPC = async () => {
    const q = new Query();
    q.setPage(1);
    q.setLimit(10);

    const req = new GetAllRequest();
    req.setQuery(q);

    const resp = await instants.foodGRPC.getAll(req, null);
    const foods = resp.getEntitiesList().map((el) => {
      const food: IFood = {
        id: el.getId(),
        name: el.getName(),
        price: el.getPrice(),
        created_at: el.getCreatedAt().toDate(),
        updated_at: el.getUpdatedAt().toDate(),
      };
      return food;
    });

    setFoods(foods);
  };

  return (
    <div className="container max-w-2xl mx-auto pt-12">
      <h1 className="text-center">Food Protobuf Format GRPC</h1>
      <div className="grid grid-cols-2 mt-6">{foods.map((el) => Card(el))}</div>
    </div>
  );
};

export default FoodsGRPC;

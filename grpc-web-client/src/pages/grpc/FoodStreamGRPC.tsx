import Card from "../../components/card";
import React, { useState } from "react";
import instants from "../../instants/instants";
import { Empty } from "../../protos/common_pb";
import { IFood } from "../../types/food";

const FoodsStreamGRPC = () => {
  const [foods, setFoods] = useState<IFood[]>([]);

  const getStream = () => {
    const req = new Empty();
    const resp: IFood[] = [];

    const stream = instants.foodGRPC.getAllStream(req);
    stream.on("data", (data) => {
      console.log(data.getName());
      const food = {
        id: data.getId(),
        name: data.getName(),
        price: data.getPrice(),
        created_at: data.getCreatedAt().toDate(),
        updated_at: data.getUpdatedAt().toDate(),
      };

      resp.push(food);
    });

    stream.on("status", function (status) {
      console.log(status.code);
      console.log(status.details);
      console.log(status.metadata);
    });

    stream.on("end", function () {
      // console.log(resp);
      setFoods(resp);
      // stream end signal
      console.log("end");
    });

    // to close the stream
    // stream.cancel();
  };

  return (
    <div className="container max-w-2xl mx-auto pt-12">
      <h1 className="text-center">Food Protobuf Format GRPC</h1>
      <button
        onClick={() => {
          getStream();
        }}
      >
        GET DATA
      </button>
      <div className="grid grid-cols-2 mt-6">{foods.map((el) => Card(el))}</div>
    </div>
  );
};

export default FoodsStreamGRPC;

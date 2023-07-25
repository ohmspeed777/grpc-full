import Card from "../../components/card";
import configs from "../../configs/configs";
import React, { useEffect, useState } from "react";
import axios from "axios";

const FoodsJson = () => {
  const [foods, setFoods] = useState([]);

  useEffect(() => {
    (async () => {
      const data = await callAPI();
      setFoods(data);
    })();
  }, []);

  const callAPI = async () => {
    const resp = await axios.get(`${configs.HOST_API}/api/v1/foods`);
    return resp.data?.entities ?? [];
  };

  return (
    <div className="container max-w-2xl mx-auto pt-12">
      <h1 className="text-center">Food Json Format HTTP</h1>
      <div className="grid grid-cols-2 mt-6">{foods.map((el) => Card(el))}</div>
    </div>
  );
};

export default FoodsJson;

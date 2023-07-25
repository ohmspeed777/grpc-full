import Card from "../../components/card";
import configs from "../../configs/configs";
import React, { useEffect, useState } from "react";
import axios from "axios";

const FoodsJson = () => {
  const [foods, setFoods] = useState([]);

  useEffect(() => {
    (async () => {
      try {
        const data = await callAPI();
        setFoods(data);
      } catch (e) {
        console.log(e);
      }
    })();
  }, []);

  const callAPI = async () => {
    try {
      const resp = await axios.get(`${configs.HOST_API}/api/v1/foods`);
      return resp.data?.entities ?? [];
    } catch (e) {
      throw e
    }
  };

  return (
    <div className="container max-w-2xl mx-auto pt-12">
      <h1 className="text-center">Food Json Format HTTP</h1>
      <div className="grid grid-cols-2 mt-6">{foods.map((el) => Card(el))}</div>
    </div>
  );
};

export default FoodsJson;

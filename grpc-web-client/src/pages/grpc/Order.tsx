import React, { useEffect, useState } from "react";
import { SignInReq, GetMyOrderReq } from "../../protos/users_pb";
import instants from "../../instants/instants";
import { IOrder, Item } from "../../types/food";
import Order from "../../components/Order";

const OrderGRPC = () => {
  const [orders, setOrders] = useState<IOrder[]>([]);
  useEffect(() => {
    (async () => {
      await callGRPC();
    })();
  }, []);

  const callGRPC = async () => {
    const req = new GetMyOrderReq();

    const metaData = {"authorization": `Bearer ${localStorage.getItem("token")}`}

    // remain metadata
    try {
      const resp = await instants.userGRPC.getMyOrder(req, metaData);
      const orders = resp.getEntitiesList();

      const entities: IOrder[] = orders.map((order) => {
        const entity: IOrder = {
          id: order.getId(),
          total: order.getTotal(),
          items: order.getItemsList().map((el) => {
            const product = el.getProduct();
            const item: Item = {
              quantity: el.getQuantity(),
              product: {
                id: product?.getId() ?? "00000",
                name: product?.getName() ?? "not found",
                price: product?.getPrice() ?? 0,
                created_at: product?.getCreatedAt().toDate(),
                updated_at: product?.getUpdatedAt().toDate(),
              },
            };
            return item;
          }),
        };
        return entity;
      });
      setOrders(entities);
    } catch (err) {
      console.log(err);
    }
  };

  const renderOrder = () => {
    return orders.map((el) => Order(el));
  };

  return (
    <div className="container max-w-2xl mx-auto pt-12">
      <h1 className="text-center">Food Protobuf Format GRPC</h1>
      <div className="grid grid-cols-2 mt-6">{renderOrder()}</div>
    </div>
  );
};

export default OrderGRPC;

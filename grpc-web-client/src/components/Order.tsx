import React from "react";

interface Product {
  id: string;
  name: string;
  price: number;
  created_at: Date;
  updated_at: Date;
}

interface Item {
  quantity: number;
  product: Product;
}

interface IProps {
  total: number;
  id: string;
  items: Item[];
}

const Order = (props: IProps) => {
  const { items, total, id } = props;
  return (
    <div className="card  bg-base-100 shadow-xl mx-1 mb-2" key={id}>
      <figure>
        <img
          src="https://plus.unsplash.com/premium_photo-1663852297267-827c73e7529e?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1170&q=80"
          alt="food"
        />
      </figure>
      <div className="card-body">
        <h2 className="card-title">{id}</h2>
        <ul className="menu bg-base-200 w-56 rounded-box">
          {items.map((el) => {
            return (
              <li>
                {el.product.name} : {el.product.price}
              </li>
            );
          })}
        </ul>
        <h4>{total}</h4>
      </div>
    </div>
  );
};

export default Order;

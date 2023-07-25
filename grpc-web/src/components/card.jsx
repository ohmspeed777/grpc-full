import React from "react";

const Card = (props) => {
  const { name, price, id } = props
  return (
    <div className="card  bg-base-100 shadow-xl mx-1 mb-2" key={id}>
      <figure>
        <img
          src="https://plus.unsplash.com/premium_photo-1663852297267-827c73e7529e?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1170&q=80"
          alt="food"
        />
      </figure>
      <div className="card-body">
        <h2 className="card-title">{name}</h2>
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Aut, ipsam?
        </p>
        <div className="card-actions justify-end">
          <button className="btn btn-primary">{price}</button>
        </div>
      </div>
    </div>
  );
};

export default Card;

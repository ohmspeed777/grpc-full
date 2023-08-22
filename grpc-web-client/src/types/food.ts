export interface IFood {
  id: string;
  name: string;
  price: number;
  created_at: Date;
  updated_at: Date;
}

export interface Item {
  quantity: number;
  product: IFood;
}

export interface IOrder {
  total: number;
  id: string;
  items: Item[];
}

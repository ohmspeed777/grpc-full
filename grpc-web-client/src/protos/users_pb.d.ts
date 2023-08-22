import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as src_protos_common_pb from '../../src/protos/common_pb';
import * as src_protos_foods_pb from '../../src/protos/foods_pb';


export class SignInReq extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): SignInReq;

  getPassword(): string;
  setPassword(value: string): SignInReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignInReq.AsObject;
  static toObject(includeInstance: boolean, msg: SignInReq): SignInReq.AsObject;
  static serializeBinaryToWriter(message: SignInReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignInReq;
  static deserializeBinaryFromReader(message: SignInReq, reader: jspb.BinaryReader): SignInReq;
}

export namespace SignInReq {
  export type AsObject = {
    email: string,
    password: string,
  }
}

export class GetMyOrderReq extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMyOrderReq.AsObject;
  static toObject(includeInstance: boolean, msg: GetMyOrderReq): GetMyOrderReq.AsObject;
  static serializeBinaryToWriter(message: GetMyOrderReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMyOrderReq;
  static deserializeBinaryFromReader(message: GetMyOrderReq, reader: jspb.BinaryReader): GetMyOrderReq;
}

export namespace GetMyOrderReq {
  export type AsObject = {
  }
}

export class SignInResp extends jspb.Message {
  getToken(): string;
  setToken(value: string): SignInResp;

  getRefreshToken(): string;
  setRefreshToken(value: string): SignInResp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignInResp.AsObject;
  static toObject(includeInstance: boolean, msg: SignInResp): SignInResp.AsObject;
  static serializeBinaryToWriter(message: SignInResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignInResp;
  static deserializeBinaryFromReader(message: SignInResp, reader: jspb.BinaryReader): SignInResp;
}

export namespace SignInResp {
  export type AsObject = {
    token: string,
    refreshToken: string,
  }
}

export class GetMyOrderResp extends jspb.Message {
  getPageInfo(): src_protos_common_pb.PageInfo | undefined;
  setPageInfo(value?: src_protos_common_pb.PageInfo): GetMyOrderResp;
  hasPageInfo(): boolean;
  clearPageInfo(): GetMyOrderResp;

  getEntitiesList(): Array<Order>;
  setEntitiesList(value: Array<Order>): GetMyOrderResp;
  clearEntitiesList(): GetMyOrderResp;
  addEntities(value?: Order, index?: number): Order;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMyOrderResp.AsObject;
  static toObject(includeInstance: boolean, msg: GetMyOrderResp): GetMyOrderResp.AsObject;
  static serializeBinaryToWriter(message: GetMyOrderResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMyOrderResp;
  static deserializeBinaryFromReader(message: GetMyOrderResp, reader: jspb.BinaryReader): GetMyOrderResp;
}

export namespace GetMyOrderResp {
  export type AsObject = {
    pageInfo?: src_protos_common_pb.PageInfo.AsObject,
    entitiesList: Array<Order.AsObject>,
  }
}

export class Order extends jspb.Message {
  getId(): string;
  setId(value: string): Order;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Order;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): Order;

  getItemsList(): Array<OrderItem>;
  setItemsList(value: Array<OrderItem>): Order;
  clearItemsList(): Order;
  addItems(value?: OrderItem, index?: number): OrderItem;

  getTotal(): number;
  setTotal(value: number): Order;

  getUser(): User | undefined;
  setUser(value?: User): Order;
  hasUser(): boolean;
  clearUser(): Order;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Order.AsObject;
  static toObject(includeInstance: boolean, msg: Order): Order.AsObject;
  static serializeBinaryToWriter(message: Order, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Order;
  static deserializeBinaryFromReader(message: Order, reader: jspb.BinaryReader): Order;
}

export namespace Order {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    itemsList: Array<OrderItem.AsObject>,
    total: number,
    user?: User.AsObject,
  }
}

export class User extends jspb.Message {
  getId(): string;
  setId(value: string): User;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): User;
  hasCreatedAt(): boolean;
  clearCreatedAt(): User;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): User;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): User;

  getEmail(): string;
  setEmail(value: string): User;

  getFirstName(): string;
  setFirstName(value: string): User;

  getLastName(): string;
  setLastName(value: string): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    email: string,
    firstName: string,
    lastName: string,
  }
}

export class OrderItem extends jspb.Message {
  getQuantity(): number;
  setQuantity(value: number): OrderItem;

  getProductId(): string;
  setProductId(value: string): OrderItem;

  getProduct(): src_protos_foods_pb.Food | undefined;
  setProduct(value?: src_protos_foods_pb.Food): OrderItem;
  hasProduct(): boolean;
  clearProduct(): OrderItem;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OrderItem.AsObject;
  static toObject(includeInstance: boolean, msg: OrderItem): OrderItem.AsObject;
  static serializeBinaryToWriter(message: OrderItem, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OrderItem;
  static deserializeBinaryFromReader(message: OrderItem, reader: jspb.BinaryReader): OrderItem;
}

export namespace OrderItem {
  export type AsObject = {
    quantity: number,
    productId: string,
    product?: src_protos_foods_pb.Food.AsObject,
  }
}


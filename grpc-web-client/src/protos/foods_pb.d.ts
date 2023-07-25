import * as jspb from 'google-protobuf'

import * as src_protos_common_pb from '../../src/protos/common_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class GetAllRequest extends jspb.Message {
  getQuery(): src_protos_common_pb.Query | undefined;
  setQuery(value?: src_protos_common_pb.Query): GetAllRequest;
  hasQuery(): boolean;
  clearQuery(): GetAllRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllRequest): GetAllRequest.AsObject;
  static serializeBinaryToWriter(message: GetAllRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllRequest;
  static deserializeBinaryFromReader(message: GetAllRequest, reader: jspb.BinaryReader): GetAllRequest;
}

export namespace GetAllRequest {
  export type AsObject = {
    query?: src_protos_common_pb.Query.AsObject,
  }
}

export class Food extends jspb.Message {
  getId(): string;
  setId(value: string): Food;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Food;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Food;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Food;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): Food;

  getPrice(): number;
  setPrice(value: number): Food;

  getName(): string;
  setName(value: string): Food;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Food.AsObject;
  static toObject(includeInstance: boolean, msg: Food): Food.AsObject;
  static serializeBinaryToWriter(message: Food, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Food;
  static deserializeBinaryFromReader(message: Food, reader: jspb.BinaryReader): Food;
}

export namespace Food {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    price: number,
    name: string,
  }
}

export class GetAllResponse extends jspb.Message {
  getPageInfo(): src_protos_common_pb.PageInfo | undefined;
  setPageInfo(value?: src_protos_common_pb.PageInfo): GetAllResponse;
  hasPageInfo(): boolean;
  clearPageInfo(): GetAllResponse;

  getEntitiesList(): Array<Food>;
  setEntitiesList(value: Array<Food>): GetAllResponse;
  clearEntitiesList(): GetAllResponse;
  addEntities(value?: Food, index?: number): Food;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllResponse): GetAllResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllResponse;
  static deserializeBinaryFromReader(message: GetAllResponse, reader: jspb.BinaryReader): GetAllResponse;
}

export namespace GetAllResponse {
  export type AsObject = {
    pageInfo?: src_protos_common_pb.PageInfo.AsObject,
    entitiesList: Array<Food.AsObject>,
  }
}


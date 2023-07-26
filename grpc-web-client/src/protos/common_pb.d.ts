import * as jspb from 'google-protobuf'



export class Query extends jspb.Message {
  getPage(): number;
  setPage(value: number): Query;

  getLimit(): number;
  setLimit(value: number): Query;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Query.AsObject;
  static toObject(includeInstance: boolean, msg: Query): Query.AsObject;
  static serializeBinaryToWriter(message: Query, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Query;
  static deserializeBinaryFromReader(message: Query, reader: jspb.BinaryReader): Query;
}

export namespace Query {
  export type AsObject = {
    page: number,
    limit: number,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export class PageInfo extends jspb.Message {
  getPage(): number;
  setPage(value: number): PageInfo;

  getSize(): number;
  setSize(value: number): PageInfo;

  getNumOfPages(): number;
  setNumOfPages(value: number): PageInfo;

  getAllOfEntities(): number;
  setAllOfEntities(value: number): PageInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PageInfo.AsObject;
  static toObject(includeInstance: boolean, msg: PageInfo): PageInfo.AsObject;
  static serializeBinaryToWriter(message: PageInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PageInfo;
  static deserializeBinaryFromReader(message: PageInfo, reader: jspb.BinaryReader): PageInfo;
}

export namespace PageInfo {
  export type AsObject = {
    page: number,
    size: number,
    numOfPages: number,
    allOfEntities: number,
  }
}


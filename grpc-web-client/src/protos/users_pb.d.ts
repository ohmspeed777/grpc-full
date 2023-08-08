import * as jspb from 'google-protobuf'



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


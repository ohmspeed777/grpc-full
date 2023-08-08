/**
 * @fileoverview gRPC-Web generated client stub for user
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: src/protos/users.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as src_protos_users_pb from '../../src/protos/users_pb';


export class UserServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorSignIn = new grpcWeb.MethodDescriptor(
    '/user.UserService/SignIn',
    grpcWeb.MethodType.UNARY,
    src_protos_users_pb.SignInReq,
    src_protos_users_pb.SignInResp,
    (request: src_protos_users_pb.SignInReq) => {
      return request.serializeBinary();
    },
    src_protos_users_pb.SignInResp.deserializeBinary
  );

  signIn(
    request: src_protos_users_pb.SignInReq,
    metadata: grpcWeb.Metadata | null): Promise<src_protos_users_pb.SignInResp>;

  signIn(
    request: src_protos_users_pb.SignInReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: src_protos_users_pb.SignInResp) => void): grpcWeb.ClientReadableStream<src_protos_users_pb.SignInResp>;

  signIn(
    request: src_protos_users_pb.SignInReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: src_protos_users_pb.SignInResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.UserService/SignIn',
        request,
        metadata || {},
        this.methodDescriptorSignIn,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.UserService/SignIn',
    request,
    metadata || {},
    this.methodDescriptorSignIn);
  }

}


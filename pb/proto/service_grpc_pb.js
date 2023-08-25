// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_service_pb = require('../proto/service_pb.js');

function serialize_messenger_Req(arg) {
  if (!(arg instanceof proto_service_pb.Req)) {
    throw new Error('Expected argument of type messenger.Req');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_messenger_Req(buffer_arg) {
  return proto_service_pb.Req.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_messenger_Res(arg) {
  if (!(arg instanceof proto_service_pb.Res)) {
    throw new Error('Expected argument of type messenger.Res');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_messenger_Res(buffer_arg) {
  return proto_service_pb.Res.deserializeBinary(new Uint8Array(buffer_arg));
}


var MessengerService = exports.MessengerService = {
  send: {
    path: '/messenger.Messenger/Send',
    requestStream: true,
    responseStream: true,
    requestType: proto_service_pb.Req,
    responseType: proto_service_pb.Res,
    requestSerialize: serialize_messenger_Req,
    requestDeserialize: deserialize_messenger_Req,
    responseSerialize: serialize_messenger_Res,
    responseDeserialize: deserialize_messenger_Res,
  },
};

exports.MessengerClient = grpc.makeGenericClientConstructor(MessengerService);

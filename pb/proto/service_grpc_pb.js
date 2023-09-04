// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_service_pb = require('../proto/service_pb.js');

function serialize_messenger_Request(arg) {
  if (!(arg instanceof proto_service_pb.Request)) {
    throw new Error('Expected argument of type messenger.Request');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_messenger_Request(buffer_arg) {
  return proto_service_pb.Request.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_messenger_Response(arg) {
  if (!(arg instanceof proto_service_pb.Response)) {
    throw new Error('Expected argument of type messenger.Response');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_messenger_Response(buffer_arg) {
  return proto_service_pb.Response.deserializeBinary(new Uint8Array(buffer_arg));
}


var MessengerService = exports.MessengerService = {
  connect: {
    path: '/messenger.Messenger/Connect',
    requestStream: false,
    responseStream: false,
    requestType: proto_service_pb.Request,
    responseType: proto_service_pb.Response,
    requestSerialize: serialize_messenger_Request,
    requestDeserialize: deserialize_messenger_Request,
    responseSerialize: serialize_messenger_Response,
    responseDeserialize: deserialize_messenger_Response,
  },
};

exports.MessengerClient = grpc.makeGenericClientConstructor(MessengerService);

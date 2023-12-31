// source: src/protos/foods.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

var src_protos_common_pb = require('../../src/protos/common_pb.js');
goog.object.extend(proto, src_protos_common_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.foods.Food', null, global);
goog.exportSymbol('proto.foods.GetAllRequest', null, global);
goog.exportSymbol('proto.foods.GetAllResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.foods.GetAllRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.foods.GetAllRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.foods.GetAllRequest.displayName = 'proto.foods.GetAllRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.foods.Food = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.foods.Food, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.foods.Food.displayName = 'proto.foods.Food';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.foods.GetAllResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.foods.GetAllResponse.repeatedFields_, null);
};
goog.inherits(proto.foods.GetAllResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.foods.GetAllResponse.displayName = 'proto.foods.GetAllResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.foods.GetAllRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.foods.GetAllRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.foods.GetAllRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.foods.GetAllRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    query: (f = msg.getQuery()) && src_protos_common_pb.Query.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.foods.GetAllRequest}
 */
proto.foods.GetAllRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.foods.GetAllRequest;
  return proto.foods.GetAllRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.foods.GetAllRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.foods.GetAllRequest}
 */
proto.foods.GetAllRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new src_protos_common_pb.Query;
      reader.readMessage(value,src_protos_common_pb.Query.deserializeBinaryFromReader);
      msg.setQuery(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.foods.GetAllRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.foods.GetAllRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.foods.GetAllRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.foods.GetAllRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getQuery();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      src_protos_common_pb.Query.serializeBinaryToWriter
    );
  }
};


/**
 * optional common.Query query = 1;
 * @return {?proto.common.Query}
 */
proto.foods.GetAllRequest.prototype.getQuery = function() {
  return /** @type{?proto.common.Query} */ (
    jspb.Message.getWrapperField(this, src_protos_common_pb.Query, 1));
};


/**
 * @param {?proto.common.Query|undefined} value
 * @return {!proto.foods.GetAllRequest} returns this
*/
proto.foods.GetAllRequest.prototype.setQuery = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.foods.GetAllRequest} returns this
 */
proto.foods.GetAllRequest.prototype.clearQuery = function() {
  return this.setQuery(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.foods.GetAllRequest.prototype.hasQuery = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.foods.Food.prototype.toObject = function(opt_includeInstance) {
  return proto.foods.Food.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.foods.Food} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.foods.Food.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    createdAt: (f = msg.getCreatedAt()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    updatedAt: (f = msg.getUpdatedAt()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    price: jspb.Message.getFloatingPointFieldWithDefault(msg, 4, 0.0),
    name: jspb.Message.getFieldWithDefault(msg, 5, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.foods.Food}
 */
proto.foods.Food.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.foods.Food;
  return proto.foods.Food.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.foods.Food} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.foods.Food}
 */
proto.foods.Food.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreatedAt(value);
      break;
    case 3:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setUpdatedAt(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setPrice(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.foods.Food.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.foods.Food.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.foods.Food} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.foods.Food.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getCreatedAt();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getUpdatedAt();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getPrice();
  if (f !== 0.0) {
    writer.writeDouble(
      4,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.foods.Food.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.foods.Food} returns this
 */
proto.foods.Food.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional google.protobuf.Timestamp created_at = 2;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.foods.Food.prototype.getCreatedAt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 2));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.foods.Food} returns this
*/
proto.foods.Food.prototype.setCreatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.foods.Food} returns this
 */
proto.foods.Food.prototype.clearCreatedAt = function() {
  return this.setCreatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.foods.Food.prototype.hasCreatedAt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.Timestamp updated_at = 3;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.foods.Food.prototype.getUpdatedAt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.foods.Food} returns this
*/
proto.foods.Food.prototype.setUpdatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.foods.Food} returns this
 */
proto.foods.Food.prototype.clearUpdatedAt = function() {
  return this.setUpdatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.foods.Food.prototype.hasUpdatedAt = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional double price = 4;
 * @return {number}
 */
proto.foods.Food.prototype.getPrice = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 4, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.foods.Food} returns this
 */
proto.foods.Food.prototype.setPrice = function(value) {
  return jspb.Message.setProto3FloatField(this, 4, value);
};


/**
 * optional string name = 5;
 * @return {string}
 */
proto.foods.Food.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.foods.Food} returns this
 */
proto.foods.Food.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.foods.GetAllResponse.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.foods.GetAllResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.foods.GetAllResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.foods.GetAllResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.foods.GetAllResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    pageInfo: (f = msg.getPageInfo()) && src_protos_common_pb.PageInfo.toObject(includeInstance, f),
    entitiesList: jspb.Message.toObjectList(msg.getEntitiesList(),
    proto.foods.Food.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.foods.GetAllResponse}
 */
proto.foods.GetAllResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.foods.GetAllResponse;
  return proto.foods.GetAllResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.foods.GetAllResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.foods.GetAllResponse}
 */
proto.foods.GetAllResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new src_protos_common_pb.PageInfo;
      reader.readMessage(value,src_protos_common_pb.PageInfo.deserializeBinaryFromReader);
      msg.setPageInfo(value);
      break;
    case 2:
      var value = new proto.foods.Food;
      reader.readMessage(value,proto.foods.Food.deserializeBinaryFromReader);
      msg.addEntities(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.foods.GetAllResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.foods.GetAllResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.foods.GetAllResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.foods.GetAllResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPageInfo();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      src_protos_common_pb.PageInfo.serializeBinaryToWriter
    );
  }
  f = message.getEntitiesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.foods.Food.serializeBinaryToWriter
    );
  }
};


/**
 * optional common.PageInfo page_info = 1;
 * @return {?proto.common.PageInfo}
 */
proto.foods.GetAllResponse.prototype.getPageInfo = function() {
  return /** @type{?proto.common.PageInfo} */ (
    jspb.Message.getWrapperField(this, src_protos_common_pb.PageInfo, 1));
};


/**
 * @param {?proto.common.PageInfo|undefined} value
 * @return {!proto.foods.GetAllResponse} returns this
*/
proto.foods.GetAllResponse.prototype.setPageInfo = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.foods.GetAllResponse} returns this
 */
proto.foods.GetAllResponse.prototype.clearPageInfo = function() {
  return this.setPageInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.foods.GetAllResponse.prototype.hasPageInfo = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated Food entities = 2;
 * @return {!Array<!proto.foods.Food>}
 */
proto.foods.GetAllResponse.prototype.getEntitiesList = function() {
  return /** @type{!Array<!proto.foods.Food>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.foods.Food, 2));
};


/**
 * @param {!Array<!proto.foods.Food>} value
 * @return {!proto.foods.GetAllResponse} returns this
*/
proto.foods.GetAllResponse.prototype.setEntitiesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.foods.Food=} opt_value
 * @param {number=} opt_index
 * @return {!proto.foods.Food}
 */
proto.foods.GetAllResponse.prototype.addEntities = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.foods.Food, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.foods.GetAllResponse} returns this
 */
proto.foods.GetAllResponse.prototype.clearEntitiesList = function() {
  return this.setEntitiesList([]);
};


goog.object.extend(exports, proto.foods);

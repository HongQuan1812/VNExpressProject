/**
 * @fileoverview gRPC-Web generated client stub for vnexpress_selector
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.1
// source: VNExpress_selector.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.vnexpress_selector = require('./VNExpress_selector_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.vnexpress_selector.VNExpress_selectorClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.vnexpress_selector.VNExpress_selectorPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.vnexpress_selector.Range,
 *   !proto.vnexpress_selector.News>}
 */
const methodDescriptor_VNExpress_selector_Select_news = new grpc.web.MethodDescriptor(
  '/vnexpress_selector.VNExpress_selector/Select_news',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.vnexpress_selector.Range,
  proto.vnexpress_selector.News,
  /**
   * @param {!proto.vnexpress_selector.Range} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.vnexpress_selector.News.deserializeBinary
);


/**
 * @param {!proto.vnexpress_selector.Range} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.vnexpress_selector.News>}
 *     The XHR Node Readable Stream
 */
proto.vnexpress_selector.VNExpress_selectorClient.prototype.select_news =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/vnexpress_selector.VNExpress_selector/Select_news',
      request,
      metadata || {},
      methodDescriptor_VNExpress_selector_Select_news);
};


/**
 * @param {!proto.vnexpress_selector.Range} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.vnexpress_selector.News>}
 *     The XHR Node Readable Stream
 */
proto.vnexpress_selector.VNExpress_selectorPromiseClient.prototype.select_news =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/vnexpress_selector.VNExpress_selector/Select_news',
      request,
      metadata || {},
      methodDescriptor_VNExpress_selector_Select_news);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.vnexpress_selector.Range,
 *   !proto.vnexpress_selector.Podcast>}
 */
const methodDescriptor_VNExpress_selector_Select_podcast = new grpc.web.MethodDescriptor(
  '/vnexpress_selector.VNExpress_selector/Select_podcast',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.vnexpress_selector.Range,
  proto.vnexpress_selector.Podcast,
  /**
   * @param {!proto.vnexpress_selector.Range} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.vnexpress_selector.Podcast.deserializeBinary
);


/**
 * @param {!proto.vnexpress_selector.Range} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.vnexpress_selector.Podcast>}
 *     The XHR Node Readable Stream
 */
proto.vnexpress_selector.VNExpress_selectorClient.prototype.select_podcast =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/vnexpress_selector.VNExpress_selector/Select_podcast',
      request,
      metadata || {},
      methodDescriptor_VNExpress_selector_Select_podcast);
};


/**
 * @param {!proto.vnexpress_selector.Range} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.vnexpress_selector.Podcast>}
 *     The XHR Node Readable Stream
 */
proto.vnexpress_selector.VNExpress_selectorPromiseClient.prototype.select_podcast =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/vnexpress_selector.VNExpress_selector/Select_podcast',
      request,
      metadata || {},
      methodDescriptor_VNExpress_selector_Select_podcast);
};


module.exports = proto.vnexpress_selector;


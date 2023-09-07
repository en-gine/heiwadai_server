///
//  Generated code. Do not modify.
//  source: v1/user/Post.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'Post.pb.dart' as $7;
import '../../google/protobuf/empty.pb.dart' as $3;
import 'Post.pbjson.dart';

export 'Post.pb.dart';

abstract class PostControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$7.PostResponse> getPostByID($pb.ServerContext ctx, $7.PostRequest request);
  $async.Future<$7.PostsResponse> getPosts($pb.ServerContext ctx, $3.Empty request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'GetPostByID': return $7.PostRequest();
      case 'GetPosts': return $3.Empty();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'GetPostByID': return this.getPostByID(ctx, request as $7.PostRequest);
      case 'GetPosts': return this.getPosts(ctx, request as $3.Empty);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => PostControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => PostControllerServiceBase$messageJson;
}

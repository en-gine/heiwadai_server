///
//  Generated code. Do not modify.
//  source: v1/user/Messages.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'Messages.pb.dart' as $11;
import 'Messages.pbjson.dart';

export 'Messages.pb.dart';

abstract class MessageControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$11.MessagesResponse> getMessagesAfter($pb.ServerContext ctx, $11.MessageRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'GetMessagesAfter': return $11.MessageRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'GetMessagesAfter': return this.getMessagesAfter(ctx, request as $11.MessageRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => MessageControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => MessageControllerServiceBase$messageJson;
}


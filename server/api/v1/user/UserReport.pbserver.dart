///
//  Generated code. Do not modify.
//  source: v1/user/UserReport.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'UserReport.pb.dart' as $16;
import '../../google/protobuf/empty.pb.dart' as $2;
import 'UserReport.pbjson.dart';

export 'UserReport.pb.dart';

abstract class UserReportControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$2.Empty> send($pb.ServerContext ctx, $16.UserReportRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Send': return $16.UserReportRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Send': return this.send(ctx, request as $16.UserReportRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => UserReportControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => UserReportControllerServiceBase$messageJson;
}


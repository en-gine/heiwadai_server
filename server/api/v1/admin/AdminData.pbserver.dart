///
//  Generated code. Do not modify.
//  source: v1/admin/AdminData.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'AdminData.pb.dart' as $0;
import 'AdminData.pbjson.dart';

export 'AdminData.pb.dart';

abstract class AdminDataControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$0.AdminDataResponse> update($pb.ServerContext ctx, $0.AdminUpdateDataRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Update': return $0.AdminUpdateDataRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Update': return this.update(ctx, request as $0.AdminUpdateDataRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => AdminDataControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => AdminDataControllerServiceBase$messageJson;
}


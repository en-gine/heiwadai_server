///
//  Generated code. Do not modify.
//  source: v1/admin/UserData.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'UserData.pb.dart' as $3;
import '../../google/protobuf/empty.pb.dart' as $1;
import 'UserData.pbjson.dart';

export 'UserData.pb.dart';

abstract class UserDataControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$1.Empty> update($pb.ServerContext ctx, $3.UserUpdateDataRequest request);
  $async.Future<$1.Empty> delete($pb.ServerContext ctx, $3.UserDeleteRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Update': return $3.UserUpdateDataRequest();
      case 'Delete': return $3.UserDeleteRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Update': return this.update(ctx, request as $3.UserUpdateDataRequest);
      case 'Delete': return this.delete(ctx, request as $3.UserDeleteRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => UserDataControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => UserDataControllerServiceBase$messageJson;
}


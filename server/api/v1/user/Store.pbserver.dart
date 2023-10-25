///
//  Generated code. Do not modify.
//  source: v1/user/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'Store.pb.dart' as $14;
import '../shared/Store.pb.dart' as $1;
import '../../google/protobuf/empty.pb.dart' as $2;
import 'Store.pbjson.dart';

export 'Store.pb.dart';

abstract class StoreControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$1.Store> getByID($pb.ServerContext ctx, $14.SoreIDRequest request);
  $async.Future<$1.Stores> getAll($pb.ServerContext ctx, $2.Empty request);
  $async.Future<$1.StayableStores> getStayables($pb.ServerContext ctx, $2.Empty request);
  $async.Future<$1.StayableStore> getStayableByID($pb.ServerContext ctx, $14.SoreIDRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'GetByID': return $14.SoreIDRequest();
      case 'GetAll': return $2.Empty();
      case 'GetStayables': return $2.Empty();
      case 'GetStayableByID': return $14.SoreIDRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'GetByID': return this.getByID(ctx, request as $14.SoreIDRequest);
      case 'GetAll': return this.getAll(ctx, request as $2.Empty);
      case 'GetStayables': return this.getStayables(ctx, request as $2.Empty);
      case 'GetStayableByID': return this.getStayableByID(ctx, request as $14.SoreIDRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => StoreControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => StoreControllerServiceBase$messageJson;
}


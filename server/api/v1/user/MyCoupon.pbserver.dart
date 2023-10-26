///
//  Generated code. Do not modify.
//  source: v1/user/MyCoupon.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'MyCoupon.pb.dart' as $12;
import '../shared/Coupon.pb.dart' as $9;
import '../../google/protobuf/empty.pb.dart' as $1;
import 'MyCoupon.pbjson.dart';

export 'MyCoupon.pb.dart';

abstract class MyCouponControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$9.Coupon> getDetail($pb.ServerContext ctx, $12.CouponIDRequest request);
  $async.Future<$12.MyCouponsResponse> getList($pb.ServerContext ctx, $1.Empty request);
  $async.Future<$1.Empty> use($pb.ServerContext ctx, $12.CouponIDRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'GetDetail': return $12.CouponIDRequest();
      case 'GetList': return $1.Empty();
      case 'Use': return $12.CouponIDRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'GetDetail': return this.getDetail(ctx, request as $12.CouponIDRequest);
      case 'GetList': return this.getList(ctx, request as $1.Empty);
      case 'Use': return this.use(ctx, request as $12.CouponIDRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => MyCouponControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => MyCouponControllerServiceBase$messageJson;
}


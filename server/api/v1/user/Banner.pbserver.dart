///
//  Generated code. Do not modify.
//  source: v1/user/Banner.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import '../../google/protobuf/empty.pb.dart' as $1;
import 'Banner.pb.dart' as $5;
import 'Banner.pbjson.dart';

export 'Banner.pb.dart';

abstract class BannerControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$5.BannerResponse> getBanner($pb.ServerContext ctx, $1.Empty request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'GetBanner': return $1.Empty();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'GetBanner': return this.getBanner(ctx, request as $1.Empty);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => BannerControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => BannerControllerServiceBase$messageJson;
}


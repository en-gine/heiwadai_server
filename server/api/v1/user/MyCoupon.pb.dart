///
//  Generated code. Do not modify.
//  source: v1/user/MyCoupon.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../shared/Coupon.pb.dart' as $8;
import '../../google/protobuf/empty.pb.dart' as $1;

class CouponIDRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CouponIDRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..hasRequiredFields = false
  ;

  CouponIDRequest._() : super();
  factory CouponIDRequest({
    $core.String? iD,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    return _result;
  }
  factory CouponIDRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CouponIDRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CouponIDRequest clone() => CouponIDRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CouponIDRequest copyWith(void Function(CouponIDRequest) updates) => super.copyWith((message) => updates(message as CouponIDRequest)) as CouponIDRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CouponIDRequest create() => CouponIDRequest._();
  CouponIDRequest createEmptyInstance() => create();
  static $pb.PbList<CouponIDRequest> createRepeated() => $pb.PbList<CouponIDRequest>();
  @$core.pragma('dart2js:noInline')
  static CouponIDRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CouponIDRequest>(create);
  static CouponIDRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get iD => $_getSZ(0);
  @$pb.TagNumber(1)
  set iD($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);
}

class MyCouponsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MyCouponsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..pc<$8.Coupon>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Coupons', $pb.PbFieldType.PM, protoName: 'Coupons', subBuilder: $8.Coupon.create)
    ..hasRequiredFields = false
  ;

  MyCouponsResponse._() : super();
  factory MyCouponsResponse({
    $core.Iterable<$8.Coupon>? coupons,
  }) {
    final _result = create();
    if (coupons != null) {
      _result.coupons.addAll(coupons);
    }
    return _result;
  }
  factory MyCouponsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MyCouponsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MyCouponsResponse clone() => MyCouponsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MyCouponsResponse copyWith(void Function(MyCouponsResponse) updates) => super.copyWith((message) => updates(message as MyCouponsResponse)) as MyCouponsResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MyCouponsResponse create() => MyCouponsResponse._();
  MyCouponsResponse createEmptyInstance() => create();
  static $pb.PbList<MyCouponsResponse> createRepeated() => $pb.PbList<MyCouponsResponse>();
  @$core.pragma('dart2js:noInline')
  static MyCouponsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MyCouponsResponse>(create);
  static MyCouponsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$8.Coupon> get coupons => $_getList(0);
}

class MyCouponControllerApi {
  $pb.RpcClient _client;
  MyCouponControllerApi(this._client);

  $async.Future<$8.Coupon> getDetail($pb.ClientContext? ctx, CouponIDRequest request) {
    var emptyResponse = $8.Coupon();
    return _client.invoke<$8.Coupon>(ctx, 'MyCouponController', 'GetDetail', request, emptyResponse);
  }
  $async.Future<MyCouponsResponse> getList($pb.ClientContext? ctx, $1.Empty request) {
    var emptyResponse = MyCouponsResponse();
    return _client.invoke<MyCouponsResponse>(ctx, 'MyCouponController', 'GetList', request, emptyResponse);
  }
  $async.Future<$1.Empty> use($pb.ClientContext? ctx, CouponIDRequest request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'MyCouponController', 'Use', request, emptyResponse);
  }
}


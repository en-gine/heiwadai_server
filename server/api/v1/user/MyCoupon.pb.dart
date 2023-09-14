///
//  Generated code. Do not modify.
//  source: v1/user/MyCoupon.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/timestamp.pb.dart' as $6;
import '../shared/Store.pb.dart' as $7;
import '../../google/protobuf/empty.pb.dart' as $1;

import 'MyCoupon.pbenum.dart';

export 'MyCoupon.pbenum.dart';

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

class Coupon extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Coupon', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..e<CouponType>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'CouponType', $pb.PbFieldType.OE, protoName: 'CouponType', defaultOrMaker: CouponType.COUPON_STANDARD, valueOf: CouponType.valueOf, enumValues: CouponType.values)
    ..a<$core.int>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'DiscountAmount', $pb.PbFieldType.OU3, protoName: 'DiscountAmount')
    ..aOM<$6.Timestamp>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ExpireAt', protoName: 'ExpireAt', subBuilder: $6.Timestamp.create)
    ..aOB(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsCombinationable', protoName: 'IsCombinationable')
    ..pPS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Notices', protoName: 'Notices')
    ..pc<$7.Store>(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'TargetStore', $pb.PbFieldType.PM, protoName: 'TargetStore', subBuilder: $7.Store.create)
    ..hasRequiredFields = false
  ;

  Coupon._() : super();
  factory Coupon({
    $core.String? iD,
    $core.String? name,
    CouponType? couponType,
    $core.int? discountAmount,
    $6.Timestamp? expireAt,
    $core.bool? isCombinationable,
    $core.Iterable<$core.String>? notices,
    $core.Iterable<$7.Store>? targetStore,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    if (name != null) {
      _result.name = name;
    }
    if (couponType != null) {
      _result.couponType = couponType;
    }
    if (discountAmount != null) {
      _result.discountAmount = discountAmount;
    }
    if (expireAt != null) {
      _result.expireAt = expireAt;
    }
    if (isCombinationable != null) {
      _result.isCombinationable = isCombinationable;
    }
    if (notices != null) {
      _result.notices.addAll(notices);
    }
    if (targetStore != null) {
      _result.targetStore.addAll(targetStore);
    }
    return _result;
  }
  factory Coupon.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Coupon.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Coupon clone() => Coupon()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Coupon copyWith(void Function(Coupon) updates) => super.copyWith((message) => updates(message as Coupon)) as Coupon; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Coupon create() => Coupon._();
  Coupon createEmptyInstance() => create();
  static $pb.PbList<Coupon> createRepeated() => $pb.PbList<Coupon>();
  @$core.pragma('dart2js:noInline')
  static Coupon getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Coupon>(create);
  static Coupon? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get iD => $_getSZ(0);
  @$pb.TagNumber(1)
  set iD($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  CouponType get couponType => $_getN(2);
  @$pb.TagNumber(3)
  set couponType(CouponType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasCouponType() => $_has(2);
  @$pb.TagNumber(3)
  void clearCouponType() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get discountAmount => $_getIZ(3);
  @$pb.TagNumber(4)
  set discountAmount($core.int v) { $_setUnsignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasDiscountAmount() => $_has(3);
  @$pb.TagNumber(4)
  void clearDiscountAmount() => clearField(4);

  @$pb.TagNumber(5)
  $6.Timestamp get expireAt => $_getN(4);
  @$pb.TagNumber(5)
  set expireAt($6.Timestamp v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasExpireAt() => $_has(4);
  @$pb.TagNumber(5)
  void clearExpireAt() => clearField(5);
  @$pb.TagNumber(5)
  $6.Timestamp ensureExpireAt() => $_ensure(4);

  @$pb.TagNumber(6)
  $core.bool get isCombinationable => $_getBF(5);
  @$pb.TagNumber(6)
  set isCombinationable($core.bool v) { $_setBool(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasIsCombinationable() => $_has(5);
  @$pb.TagNumber(6)
  void clearIsCombinationable() => clearField(6);

  @$pb.TagNumber(7)
  $core.List<$core.String> get notices => $_getList(6);

  @$pb.TagNumber(8)
  $core.List<$7.Store> get targetStore => $_getList(7);
}

class MyCouponsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MyCouponsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..pc<Coupon>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Coupons', $pb.PbFieldType.PM, protoName: 'Coupons', subBuilder: Coupon.create)
    ..hasRequiredFields = false
  ;

  MyCouponsResponse._() : super();
  factory MyCouponsResponse({
    $core.Iterable<Coupon>? coupons,
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
  $core.List<Coupon> get coupons => $_getList(0);
}

class MyCouponControllerApi {
  $pb.RpcClient _client;
  MyCouponControllerApi(this._client);

  $async.Future<Coupon> getDetail($pb.ClientContext? ctx, CouponIDRequest request) {
    var emptyResponse = Coupon();
    return _client.invoke<Coupon>(ctx, 'MyCouponController', 'GetDetail', request, emptyResponse);
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


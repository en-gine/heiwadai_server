///
//  Generated code. Do not modify.
//  source: v1/shared/Coupon.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/timestamp.pb.dart' as $5;
import 'Store.pb.dart' as $1;

import 'Coupon.pbenum.dart';

export 'Coupon.pbenum.dart';

class Coupon extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Coupon', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.shared'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..e<CouponType>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'CouponType', $pb.PbFieldType.OE, protoName: 'CouponType', defaultOrMaker: CouponType.COUPON_STANDARD, valueOf: CouponType.valueOf, enumValues: CouponType.values)
    ..a<$core.int>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'DiscountAmount', $pb.PbFieldType.OU3, protoName: 'DiscountAmount')
    ..aOM<$5.Timestamp>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ExpireAt', protoName: 'ExpireAt', subBuilder: $5.Timestamp.create)
    ..aOB(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsCombinationable', protoName: 'IsCombinationable')
    ..pPS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Notices', protoName: 'Notices')
    ..pc<$1.Store>(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'TargetStore', $pb.PbFieldType.PM, protoName: 'TargetStore', subBuilder: $1.Store.create)
    ..hasRequiredFields = false
  ;

  Coupon._() : super();
  factory Coupon({
    $core.String? iD,
    $core.String? name,
    CouponType? couponType,
    $core.int? discountAmount,
    $5.Timestamp? expireAt,
    $core.bool? isCombinationable,
    $core.Iterable<$core.String>? notices,
    $core.Iterable<$1.Store>? targetStore,
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
  $5.Timestamp get expireAt => $_getN(4);
  @$pb.TagNumber(5)
  set expireAt($5.Timestamp v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasExpireAt() => $_has(4);
  @$pb.TagNumber(5)
  void clearExpireAt() => clearField(5);
  @$pb.TagNumber(5)
  $5.Timestamp ensureExpireAt() => $_ensure(4);

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
  $core.List<$1.Store> get targetStore => $_getList(7);
}


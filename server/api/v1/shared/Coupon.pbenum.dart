///
//  Generated code. Do not modify.
//  source: v1/shared/Coupon.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class CouponType extends $pb.ProtobufEnum {
  static const CouponType COUPON_STANDARD = CouponType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'COUPON_STANDARD');
  static const CouponType COUPON_CUSTOM = CouponType._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'COUPON_CUSTOM');
  static const CouponType COUPON_BIRTHDAY = CouponType._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'COUPON_BIRTHDAY');

  static const $core.List<CouponType> values = <CouponType> [
    COUPON_STANDARD,
    COUPON_CUSTOM,
    COUPON_BIRTHDAY,
  ];

  static final $core.Map<$core.int, CouponType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static CouponType? valueOf($core.int value) => _byValue[value];

  const CouponType._($core.int v, $core.String n) : super(v, n);
}


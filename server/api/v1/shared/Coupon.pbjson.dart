///
//  Generated code. Do not modify.
//  source: v1/shared/Coupon.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use couponTypeDescriptor instead')
const CouponType$json = const {
  '1': 'CouponType',
  '2': const [
    const {'1': 'COUPON_STANDARD', '2': 0},
    const {'1': 'COUPON_CUSTOM', '2': 1},
    const {'1': 'COUPON_BIRTHDAY', '2': 2},
  ],
};

/// Descriptor for `CouponType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List couponTypeDescriptor = $convert.base64Decode('CgpDb3Vwb25UeXBlEhMKD0NPVVBPTl9TVEFOREFSRBAAEhEKDUNPVVBPTl9DVVNUT00QARITCg9DT1VQT05fQklSVEhEQVkQAg==');
@$core.Deprecated('Use couponDescriptor instead')
const Coupon$json = const {
  '1': 'Coupon',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
    const {'1': 'Name', '3': 2, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'CouponType', '3': 3, '4': 1, '5': 14, '6': '.server.shared.CouponType', '10': 'CouponType'},
    const {'1': 'DiscountAmount', '3': 4, '4': 1, '5': 13, '10': 'DiscountAmount'},
    const {'1': 'ExpireAt', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'ExpireAt'},
    const {'1': 'IsCombinationable', '3': 6, '4': 1, '5': 8, '10': 'IsCombinationable'},
    const {'1': 'Notices', '3': 7, '4': 3, '5': 9, '10': 'Notices'},
    const {'1': 'TargetStore', '3': 8, '4': 3, '5': 11, '6': '.server.shared.Store', '10': 'TargetStore'},
  ],
};

/// Descriptor for `Coupon`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List couponDescriptor = $convert.base64Decode('CgZDb3Vwb24SDgoCSUQYASABKAlSAklEEhIKBE5hbWUYAiABKAlSBE5hbWUSOQoKQ291cG9uVHlwZRgDIAEoDjIZLnNlcnZlci5zaGFyZWQuQ291cG9uVHlwZVIKQ291cG9uVHlwZRImCg5EaXNjb3VudEFtb3VudBgEIAEoDVIORGlzY291bnRBbW91bnQSNgoIRXhwaXJlQXQYBSABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUghFeHBpcmVBdBIsChFJc0NvbWJpbmF0aW9uYWJsZRgGIAEoCFIRSXNDb21iaW5hdGlvbmFibGUSGAoHTm90aWNlcxgHIAMoCVIHTm90aWNlcxI2CgtUYXJnZXRTdG9yZRgIIAMoCzIULnNlcnZlci5zaGFyZWQuU3RvcmVSC1RhcmdldFN0b3Jl');

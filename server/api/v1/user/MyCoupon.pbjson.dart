///
//  Generated code. Do not modify.
//  source: v1/user/MyCoupon.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/timestamp.pbjson.dart' as $6;
import '../shared/Store.pbjson.dart' as $7;
import '../../google/protobuf/empty.pbjson.dart' as $1;

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
@$core.Deprecated('Use couponIDRequestDescriptor instead')
const CouponIDRequest$json = const {
  '1': 'CouponIDRequest',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
  ],
};

/// Descriptor for `CouponIDRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List couponIDRequestDescriptor = $convert.base64Decode('Cg9Db3Vwb25JRFJlcXVlc3QSDgoCSUQYASABKAlSAklE');
@$core.Deprecated('Use couponDescriptor instead')
const Coupon$json = const {
  '1': 'Coupon',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
    const {'1': 'Name', '3': 2, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'CouponType', '3': 3, '4': 1, '5': 14, '6': '.server.user.CouponType', '10': 'CouponType'},
    const {'1': 'DiscountAmount', '3': 4, '4': 1, '5': 13, '10': 'DiscountAmount'},
    const {'1': 'ExpireAt', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'ExpireAt'},
    const {'1': 'IsCombinationable', '3': 6, '4': 1, '5': 8, '10': 'IsCombinationable'},
    const {'1': 'Notices', '3': 7, '4': 3, '5': 9, '10': 'Notices'},
    const {'1': 'TargetStore', '3': 8, '4': 3, '5': 11, '6': '.server.shared.Store', '10': 'TargetStore'},
  ],
};

/// Descriptor for `Coupon`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List couponDescriptor = $convert.base64Decode('CgZDb3Vwb24SDgoCSUQYASABKAlSAklEEhIKBE5hbWUYAiABKAlSBE5hbWUSNwoKQ291cG9uVHlwZRgDIAEoDjIXLnNlcnZlci51c2VyLkNvdXBvblR5cGVSCkNvdXBvblR5cGUSJgoORGlzY291bnRBbW91bnQYBCABKA1SDkRpc2NvdW50QW1vdW50EjYKCEV4cGlyZUF0GAUgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIIRXhwaXJlQXQSLAoRSXNDb21iaW5hdGlvbmFibGUYBiABKAhSEUlzQ29tYmluYXRpb25hYmxlEhgKB05vdGljZXMYByADKAlSB05vdGljZXMSNgoLVGFyZ2V0U3RvcmUYCCADKAsyFC5zZXJ2ZXIuc2hhcmVkLlN0b3JlUgtUYXJnZXRTdG9yZQ==');
@$core.Deprecated('Use myCouponsResponseDescriptor instead')
const MyCouponsResponse$json = const {
  '1': 'MyCouponsResponse',
  '2': const [
    const {'1': 'Coupons', '3': 1, '4': 3, '5': 11, '6': '.server.user.Coupon', '10': 'Coupons'},
  ],
};

/// Descriptor for `MyCouponsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List myCouponsResponseDescriptor = $convert.base64Decode('ChFNeUNvdXBvbnNSZXNwb25zZRItCgdDb3Vwb25zGAEgAygLMhMuc2VydmVyLnVzZXIuQ291cG9uUgdDb3Vwb25z');
const $core.Map<$core.String, $core.dynamic> MyCouponControllerServiceBase$json = const {
  '1': 'MyCouponController',
  '2': const [
    const {'1': 'GetDetail', '2': '.server.user.CouponIDRequest', '3': '.server.user.Coupon', '4': const {}},
    const {'1': 'GetList', '2': '.google.protobuf.Empty', '3': '.server.user.MyCouponsResponse', '4': const {}},
    const {'1': 'Use', '2': '.server.user.CouponIDRequest', '3': '.google.protobuf.Empty', '4': const {}},
  ],
};

@$core.Deprecated('Use myCouponControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> MyCouponControllerServiceBase$messageJson = const {
  '.server.user.CouponIDRequest': CouponIDRequest$json,
  '.server.user.Coupon': Coupon$json,
  '.google.protobuf.Timestamp': $6.Timestamp$json,
  '.server.shared.Store': $7.Store$json,
  '.google.protobuf.Empty': $1.Empty$json,
  '.server.user.MyCouponsResponse': MyCouponsResponse$json,
};

/// Descriptor for `MyCouponController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List myCouponControllerServiceDescriptor = $convert.base64Decode('ChJNeUNvdXBvbkNvbnRyb2xsZXISQAoJR2V0RGV0YWlsEhwuc2VydmVyLnVzZXIuQ291cG9uSURSZXF1ZXN0GhMuc2VydmVyLnVzZXIuQ291cG9uIgASQwoHR2V0TGlzdBIWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRoeLnNlcnZlci51c2VyLk15Q291cG9uc1Jlc3BvbnNlIgASPQoDVXNlEhwuc2VydmVyLnVzZXIuQ291cG9uSURSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5IgA=');

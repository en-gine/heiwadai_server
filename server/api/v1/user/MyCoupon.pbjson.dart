///
//  Generated code. Do not modify.
//  source: v1/user/MyCoupon.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../shared/Coupon.pbjson.dart' as $9;
import '../../google/protobuf/timestamp.pbjson.dart' as $6;
import '../shared/Store.pbjson.dart' as $3;
import '../../google/protobuf/empty.pbjson.dart' as $1;

@$core.Deprecated('Use couponIDRequestDescriptor instead')
const CouponIDRequest$json = const {
  '1': 'CouponIDRequest',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
  ],
};

/// Descriptor for `CouponIDRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List couponIDRequestDescriptor = $convert.base64Decode('Cg9Db3Vwb25JRFJlcXVlc3QSDgoCSUQYASABKAlSAklE');
@$core.Deprecated('Use myCouponsResponseDescriptor instead')
const MyCouponsResponse$json = const {
  '1': 'MyCouponsResponse',
  '2': const [
    const {'1': 'Coupons', '3': 1, '4': 3, '5': 11, '6': '.server.shared.Coupon', '10': 'Coupons'},
  ],
};

/// Descriptor for `MyCouponsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List myCouponsResponseDescriptor = $convert.base64Decode('ChFNeUNvdXBvbnNSZXNwb25zZRIvCgdDb3Vwb25zGAEgAygLMhUuc2VydmVyLnNoYXJlZC5Db3Vwb25SB0NvdXBvbnM=');
const $core.Map<$core.String, $core.dynamic> MyCouponControllerServiceBase$json = const {
  '1': 'MyCouponController',
  '2': const [
    const {'1': 'GetDetail', '2': '.server.user.CouponIDRequest', '3': '.server.shared.Coupon', '4': const {}},
    const {'1': 'GetList', '2': '.google.protobuf.Empty', '3': '.server.user.MyCouponsResponse', '4': const {}},
    const {'1': 'Use', '2': '.server.user.CouponIDRequest', '3': '.google.protobuf.Empty', '4': const {}},
  ],
};

@$core.Deprecated('Use myCouponControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> MyCouponControllerServiceBase$messageJson = const {
  '.server.user.CouponIDRequest': CouponIDRequest$json,
  '.server.shared.Coupon': $9.Coupon$json,
  '.google.protobuf.Timestamp': $6.Timestamp$json,
  '.server.shared.Store': $3.Store$json,
  '.google.protobuf.Empty': $1.Empty$json,
  '.server.user.MyCouponsResponse': MyCouponsResponse$json,
};

/// Descriptor for `MyCouponController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List myCouponControllerServiceDescriptor = $convert.base64Decode('ChJNeUNvdXBvbkNvbnRyb2xsZXISQgoJR2V0RGV0YWlsEhwuc2VydmVyLnVzZXIuQ291cG9uSURSZXF1ZXN0GhUuc2VydmVyLnNoYXJlZC5Db3Vwb24iABJDCgdHZXRMaXN0EhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5Gh4uc2VydmVyLnVzZXIuTXlDb3Vwb25zUmVzcG9uc2UiABI9CgNVc2USHC5zZXJ2ZXIudXNlci5Db3Vwb25JRFJlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkiAA==');

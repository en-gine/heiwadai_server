///
//  Generated code. Do not modify.
//  source: v1/user/Checkin.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/empty.pbjson.dart' as $1;
import '../../google/protobuf/timestamp.pbjson.dart' as $4;
import '../shared/Coupon.pbjson.dart' as $8;
import '../shared/Store.pbjson.dart' as $5;

@$core.Deprecated('Use checkinRequestDescriptor instead')
const CheckinRequest$json = const {
  '1': 'CheckinRequest',
  '2': const [
    const {'1': 'qr_hash', '3': 1, '4': 1, '5': 9, '10': 'qrHash'},
  ],
};

/// Descriptor for `CheckinRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List checkinRequestDescriptor = $convert.base64Decode('Cg5DaGVja2luUmVxdWVzdBIXCgdxcl9oYXNoGAEgASgJUgZxckhhc2g=');
@$core.Deprecated('Use checkinStampDescriptor instead')
const CheckinStamp$json = const {
  '1': 'CheckinStamp',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
    const {'1': 'StoreID', '3': 2, '4': 1, '5': 9, '10': 'StoreID'},
    const {'1': 'StoreName', '3': 3, '4': 1, '5': 9, '10': 'StoreName'},
    const {'1': 'StoreStampImage', '3': 4, '4': 1, '5': 9, '10': 'StoreStampImage'},
    const {'1': 'CheckInAt', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'CheckInAt'},
  ],
};

/// Descriptor for `CheckinStamp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List checkinStampDescriptor = $convert.base64Decode('CgxDaGVja2luU3RhbXASDgoCSUQYASABKAlSAklEEhgKB1N0b3JlSUQYAiABKAlSB1N0b3JlSUQSHAoJU3RvcmVOYW1lGAMgASgJUglTdG9yZU5hbWUSKAoPU3RvcmVTdGFtcEltYWdlGAQgASgJUg9TdG9yZVN0YW1wSW1hZ2USOAoJQ2hlY2tJbkF0GAUgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIJQ2hlY2tJbkF0');
@$core.Deprecated('Use stampCardResponseDescriptor instead')
const StampCardResponse$json = const {
  '1': 'StampCardResponse',
  '2': const [
    const {'1': 'stamps', '3': 1, '4': 3, '5': 11, '6': '.server.user.CheckinStamp', '10': 'stamps'},
  ],
};

/// Descriptor for `StampCardResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List stampCardResponseDescriptor = $convert.base64Decode('ChFTdGFtcENhcmRSZXNwb25zZRIxCgZzdGFtcHMYASADKAsyGS5zZXJ2ZXIudXNlci5DaGVja2luU3RhbXBSBnN0YW1wcw==');
@$core.Deprecated('Use checkinResponseDescriptor instead')
const CheckinResponse$json = const {
  '1': 'CheckinResponse',
  '2': const [
    const {'1': 'MayCoupon', '3': 1, '4': 1, '5': 11, '6': '.server.shared.Coupon', '10': 'MayCoupon'},
  ],
};

/// Descriptor for `CheckinResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List checkinResponseDescriptor = $convert.base64Decode('Cg9DaGVja2luUmVzcG9uc2USMwoJTWF5Q291cG9uGAEgASgLMhUuc2VydmVyLnNoYXJlZC5Db3Vwb25SCU1heUNvdXBvbg==');
const $core.Map<$core.String, $core.dynamic> CheckinControllerServiceBase$json = const {
  '1': 'CheckinController',
  '2': const [
    const {'1': 'GetStampCard', '2': '.google.protobuf.Empty', '3': '.server.user.StampCardResponse', '4': const {}},
    const {'1': 'Checkin', '2': '.server.user.CheckinRequest', '3': '.server.user.CheckinResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use checkinControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> CheckinControllerServiceBase$messageJson = const {
  '.google.protobuf.Empty': $1.Empty$json,
  '.server.user.StampCardResponse': StampCardResponse$json,
  '.server.user.CheckinStamp': CheckinStamp$json,
  '.google.protobuf.Timestamp': $4.Timestamp$json,
  '.server.user.CheckinRequest': CheckinRequest$json,
  '.server.user.CheckinResponse': CheckinResponse$json,
  '.server.shared.Coupon': $8.Coupon$json,
  '.server.shared.Store': $5.Store$json,
};

/// Descriptor for `CheckinController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List checkinControllerServiceDescriptor = $convert.base64Decode('ChFDaGVja2luQ29udHJvbGxlchJICgxHZXRTdGFtcENhcmQSFi5nb29nbGUucHJvdG9idWYuRW1wdHkaHi5zZXJ2ZXIudXNlci5TdGFtcENhcmRSZXNwb25zZSIAEkYKB0NoZWNraW4SGy5zZXJ2ZXIudXNlci5DaGVja2luUmVxdWVzdBocLnNlcnZlci51c2VyLkNoZWNraW5SZXNwb25zZSIA');

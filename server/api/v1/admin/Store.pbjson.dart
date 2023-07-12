///
//  Generated code. Do not modify.
//  source: v1/admin/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../shared/Store.pbjson.dart' as $1;

@$core.Deprecated('Use storeRegisterRequestDescriptor instead')
const StoreRegisterRequest$json = const {
  '1': 'StoreRegisterRequest',
  '2': const [
    const {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'ZipCode', '3': 2, '4': 1, '5': 9, '10': 'ZipCode'},
    const {'1': 'Address', '3': 3, '4': 1, '5': 9, '10': 'Address'},
    const {'1': 'Tel', '3': 4, '4': 1, '5': 9, '10': 'Tel'},
    const {'1': 'Parking', '3': 5, '4': 1, '5': 9, '10': 'Parking'},
    const {'1': 'AccessInfo', '3': 6, '4': 1, '5': 9, '10': 'AccessInfo'},
    const {'1': 'IsActive', '3': 7, '4': 1, '5': 8, '10': 'IsActive'},
    const {'1': 'Stayable', '3': 8, '4': 1, '5': 8, '10': 'Stayable'},
    const {'1': 'QRCode', '3': 9, '4': 1, '5': 9, '10': 'QRCode'},
    const {'1': 'UnLimitedQRCode', '3': 10, '4': 1, '5': 9, '10': 'UnLimitedQRCode'},
  ],
};

/// Descriptor for `StoreRegisterRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storeRegisterRequestDescriptor = $convert.base64Decode('ChRTdG9yZVJlZ2lzdGVyUmVxdWVzdBISCgROYW1lGAEgASgJUgROYW1lEhgKB1ppcENvZGUYAiABKAlSB1ppcENvZGUSGAoHQWRkcmVzcxgDIAEoCVIHQWRkcmVzcxIQCgNUZWwYBCABKAlSA1RlbBIYCgdQYXJraW5nGAUgASgJUgdQYXJraW5nEh4KCkFjY2Vzc0luZm8YBiABKAlSCkFjY2Vzc0luZm8SGgoISXNBY3RpdmUYByABKAhSCElzQWN0aXZlEhoKCFN0YXlhYmxlGAggASgIUghTdGF5YWJsZRIWCgZRUkNvZGUYCSABKAlSBlFSQ29kZRIoCg9VbkxpbWl0ZWRRUkNvZGUYCiABKAlSD1VuTGltaXRlZFFSQ29kZQ==');
@$core.Deprecated('Use storeUpdateRequestDescriptor instead')
const StoreUpdateRequest$json = const {
  '1': 'StoreUpdateRequest',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
    const {'1': 'Name', '3': 2, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'ZipCode', '3': 3, '4': 1, '5': 9, '10': 'ZipCode'},
    const {'1': 'Address', '3': 4, '4': 1, '5': 9, '10': 'Address'},
    const {'1': 'Tel', '3': 5, '4': 1, '5': 9, '10': 'Tel'},
    const {'1': 'Parking', '3': 6, '4': 1, '5': 9, '10': 'Parking'},
    const {'1': 'AccessInfo', '3': 7, '4': 1, '5': 9, '10': 'AccessInfo'},
    const {'1': 'IsActive', '3': 8, '4': 1, '5': 8, '10': 'IsActive'},
    const {'1': 'Stayable', '3': 9, '4': 1, '5': 8, '10': 'Stayable'},
    const {'1': 'QRCode', '3': 10, '4': 1, '5': 9, '10': 'QRCode'},
    const {'1': 'UnLimitedQRCode', '3': 11, '4': 1, '5': 9, '10': 'UnLimitedQRCode'},
  ],
};

/// Descriptor for `StoreUpdateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storeUpdateRequestDescriptor = $convert.base64Decode('ChJTdG9yZVVwZGF0ZVJlcXVlc3QSDgoCSUQYASABKAlSAklEEhIKBE5hbWUYAiABKAlSBE5hbWUSGAoHWmlwQ29kZRgDIAEoCVIHWmlwQ29kZRIYCgdBZGRyZXNzGAQgASgJUgdBZGRyZXNzEhAKA1RlbBgFIAEoCVIDVGVsEhgKB1BhcmtpbmcYBiABKAlSB1BhcmtpbmcSHgoKQWNjZXNzSW5mbxgHIAEoCVIKQWNjZXNzSW5mbxIaCghJc0FjdGl2ZRgIIAEoCFIISXNBY3RpdmUSGgoIU3RheWFibGUYCSABKAhSCFN0YXlhYmxlEhYKBlFSQ29kZRgKIAEoCVIGUVJDb2RlEigKD1VuTGltaXRlZFFSQ29kZRgLIAEoCVIPVW5MaW1pdGVkUVJDb2Rl');
const $core.Map<$core.String, $core.dynamic> StoreControllerServiceBase$json = const {
  '1': 'StoreController',
  '2': const [
    const {'1': 'Register', '2': '.server.admin.StoreRegisterRequest', '3': '.server.shared.StoreResponse', '4': const {}},
    const {'1': 'Update', '2': '.server.admin.StoreUpdateRequest', '3': '.server.shared.StoreResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use storeControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> StoreControllerServiceBase$messageJson = const {
  '.server.admin.StoreRegisterRequest': StoreRegisterRequest$json,
  '.server.shared.StoreResponse': $1.StoreResponse$json,
  '.server.admin.StoreUpdateRequest': StoreUpdateRequest$json,
};

/// Descriptor for `StoreController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List storeControllerServiceDescriptor = $convert.base64Decode('Cg9TdG9yZUNvbnRyb2xsZXISTgoIUmVnaXN0ZXISIi5zZXJ2ZXIuYWRtaW4uU3RvcmVSZWdpc3RlclJlcXVlc3QaHC5zZXJ2ZXIuc2hhcmVkLlN0b3JlUmVzcG9uc2UiABJKCgZVcGRhdGUSIC5zZXJ2ZXIuYWRtaW4uU3RvcmVVcGRhdGVSZXF1ZXN0Ghwuc2VydmVyLnNoYXJlZC5TdG9yZVJlc3BvbnNlIgA=');

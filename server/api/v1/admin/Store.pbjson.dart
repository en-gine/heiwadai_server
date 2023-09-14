///
//  Generated code. Do not modify.
//  source: v1/admin/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/empty.pbjson.dart' as $1;

@$core.Deprecated('Use storeRegisterRequestDescriptor instead')
const StoreRegisterRequest$json = const {
  '1': 'StoreRegisterRequest',
  '2': const [
    const {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'BranchName', '3': 2, '4': 1, '5': 9, '10': 'BranchName'},
    const {'1': 'ZipCode', '3': 3, '4': 1, '5': 9, '10': 'ZipCode'},
    const {'1': 'Address', '3': 4, '4': 1, '5': 9, '10': 'Address'},
    const {'1': 'Tel', '3': 5, '4': 1, '5': 9, '10': 'Tel'},
    const {'1': 'SiteURL', '3': 6, '4': 1, '5': 9, '10': 'SiteURL'},
    const {'1': 'StampImageURL', '3': 7, '4': 1, '5': 9, '10': 'StampImageURL'},
    const {'1': 'Stayable', '3': 8, '4': 1, '5': 8, '10': 'Stayable'},
    const {'1': 'IsActive', '3': 9, '4': 1, '5': 8, '10': 'IsActive'},
    const {'1': 'QRCode', '3': 10, '4': 1, '5': 9, '10': 'QRCode'},
    const {'1': 'UnLimitedQRCode', '3': 11, '4': 1, '5': 9, '10': 'UnLimitedQRCode'},
    const {'1': 'StayableInfo', '3': 12, '4': 1, '5': 11, '6': '.server.admin.StayableInfo', '10': 'StayableInfo'},
  ],
};

/// Descriptor for `StoreRegisterRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storeRegisterRequestDescriptor = $convert.base64Decode('ChRTdG9yZVJlZ2lzdGVyUmVxdWVzdBISCgROYW1lGAEgASgJUgROYW1lEh4KCkJyYW5jaE5hbWUYAiABKAlSCkJyYW5jaE5hbWUSGAoHWmlwQ29kZRgDIAEoCVIHWmlwQ29kZRIYCgdBZGRyZXNzGAQgASgJUgdBZGRyZXNzEhAKA1RlbBgFIAEoCVIDVGVsEhgKB1NpdGVVUkwYBiABKAlSB1NpdGVVUkwSJAoNU3RhbXBJbWFnZVVSTBgHIAEoCVINU3RhbXBJbWFnZVVSTBIaCghTdGF5YWJsZRgIIAEoCFIIU3RheWFibGUSGgoISXNBY3RpdmUYCSABKAhSCElzQWN0aXZlEhYKBlFSQ29kZRgKIAEoCVIGUVJDb2RlEigKD1VuTGltaXRlZFFSQ29kZRgLIAEoCVIPVW5MaW1pdGVkUVJDb2RlEj4KDFN0YXlhYmxlSW5mbxgMIAEoCzIaLnNlcnZlci5hZG1pbi5TdGF5YWJsZUluZm9SDFN0YXlhYmxlSW5mbw==');
@$core.Deprecated('Use storeUpdateRequestDescriptor instead')
const StoreUpdateRequest$json = const {
  '1': 'StoreUpdateRequest',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
    const {'1': 'Name', '3': 2, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'BranchName', '3': 3, '4': 1, '5': 9, '10': 'BranchName'},
    const {'1': 'ZipCode', '3': 4, '4': 1, '5': 9, '10': 'ZipCode'},
    const {'1': 'Address', '3': 5, '4': 1, '5': 9, '10': 'Address'},
    const {'1': 'Tel', '3': 6, '4': 1, '5': 9, '10': 'Tel'},
    const {'1': 'SiteURL', '3': 7, '4': 1, '5': 9, '10': 'SiteURL'},
    const {'1': 'StampImageURL', '3': 8, '4': 1, '5': 9, '10': 'StampImageURL'},
    const {'1': 'Stayable', '3': 9, '4': 1, '5': 8, '10': 'Stayable'},
    const {'1': 'IsActive', '3': 10, '4': 1, '5': 8, '10': 'IsActive'},
    const {'1': 'QRCode', '3': 11, '4': 1, '5': 9, '10': 'QRCode'},
    const {'1': 'UnLimitedQRCode', '3': 12, '4': 1, '5': 9, '10': 'UnLimitedQRCode'},
    const {'1': 'StayableInfo', '3': 13, '4': 1, '5': 11, '6': '.server.admin.StayableInfo', '10': 'StayableInfo'},
  ],
};

/// Descriptor for `StoreUpdateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storeUpdateRequestDescriptor = $convert.base64Decode('ChJTdG9yZVVwZGF0ZVJlcXVlc3QSDgoCSUQYASABKAlSAklEEhIKBE5hbWUYAiABKAlSBE5hbWUSHgoKQnJhbmNoTmFtZRgDIAEoCVIKQnJhbmNoTmFtZRIYCgdaaXBDb2RlGAQgASgJUgdaaXBDb2RlEhgKB0FkZHJlc3MYBSABKAlSB0FkZHJlc3MSEAoDVGVsGAYgASgJUgNUZWwSGAoHU2l0ZVVSTBgHIAEoCVIHU2l0ZVVSTBIkCg1TdGFtcEltYWdlVVJMGAggASgJUg1TdGFtcEltYWdlVVJMEhoKCFN0YXlhYmxlGAkgASgIUghTdGF5YWJsZRIaCghJc0FjdGl2ZRgKIAEoCFIISXNBY3RpdmUSFgoGUVJDb2RlGAsgASgJUgZRUkNvZGUSKAoPVW5MaW1pdGVkUVJDb2RlGAwgASgJUg9VbkxpbWl0ZWRRUkNvZGUSPgoMU3RheWFibGVJbmZvGA0gASgLMhouc2VydmVyLmFkbWluLlN0YXlhYmxlSW5mb1IMU3RheWFibGVJbmZv');
@$core.Deprecated('Use stayableInfoDescriptor instead')
const StayableInfo$json = const {
  '1': 'StayableInfo',
  '2': const [
    const {'1': 'Parking', '3': 1, '4': 1, '5': 9, '10': 'Parking'},
    const {'1': 'Latitude', '3': 2, '4': 1, '5': 1, '10': 'Latitude'},
    const {'1': 'Longitude', '3': 3, '4': 1, '5': 1, '10': 'Longitude'},
    const {'1': 'AccessInfo', '3': 4, '4': 1, '5': 9, '10': 'AccessInfo'},
    const {'1': 'RestAPIURL', '3': 5, '4': 1, '5': 9, '10': 'RestAPIURL'},
    const {'1': 'BookingSystemID', '3': 6, '4': 1, '5': 9, '10': 'BookingSystemID'},
  ],
};

/// Descriptor for `StayableInfo`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List stayableInfoDescriptor = $convert.base64Decode('CgxTdGF5YWJsZUluZm8SGAoHUGFya2luZxgBIAEoCVIHUGFya2luZxIaCghMYXRpdHVkZRgCIAEoAVIITGF0aXR1ZGUSHAoJTG9uZ2l0dWRlGAMgASgBUglMb25naXR1ZGUSHgoKQWNjZXNzSW5mbxgEIAEoCVIKQWNjZXNzSW5mbxIeCgpSZXN0QVBJVVJMGAUgASgJUgpSZXN0QVBJVVJMEigKD0Jvb2tpbmdTeXN0ZW1JRBgGIAEoCVIPQm9va2luZ1N5c3RlbUlE');
const $core.Map<$core.String, $core.dynamic> StoreControllerServiceBase$json = const {
  '1': 'StoreController',
  '2': const [
    const {'1': 'Register', '2': '.server.admin.StoreRegisterRequest', '3': '.google.protobuf.Empty', '4': const {}},
    const {'1': 'Update', '2': '.server.admin.StoreUpdateRequest', '3': '.google.protobuf.Empty', '4': const {}},
  ],
};

@$core.Deprecated('Use storeControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> StoreControllerServiceBase$messageJson = const {
  '.server.admin.StoreRegisterRequest': StoreRegisterRequest$json,
  '.server.admin.StayableInfo': StayableInfo$json,
  '.google.protobuf.Empty': $1.Empty$json,
  '.server.admin.StoreUpdateRequest': StoreUpdateRequest$json,
};

/// Descriptor for `StoreController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List storeControllerServiceDescriptor = $convert.base64Decode('Cg9TdG9yZUNvbnRyb2xsZXISSAoIUmVnaXN0ZXISIi5zZXJ2ZXIuYWRtaW4uU3RvcmVSZWdpc3RlclJlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkiABJECgZVcGRhdGUSIC5zZXJ2ZXIuYWRtaW4uU3RvcmVVcGRhdGVSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5IgA=');

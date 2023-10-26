///
//  Generated code. Do not modify.
//  source: v1/admin/AdminData.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use adminUpdateDataRequestDescriptor instead')
const AdminUpdateDataRequest$json = const {
  '1': 'AdminUpdateDataRequest',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
    const {'1': 'Name', '3': 2, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'IsActive', '3': 3, '4': 1, '5': 9, '10': 'IsActive'},
    const {'1': 'Mail', '3': 4, '4': 1, '5': 9, '10': 'Mail'},
    const {'1': 'StoreID', '3': 5, '4': 1, '5': 9, '10': 'StoreID'},
  ],
};

/// Descriptor for `AdminUpdateDataRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List adminUpdateDataRequestDescriptor = $convert.base64Decode('ChZBZG1pblVwZGF0ZURhdGFSZXF1ZXN0Eg4KAklEGAEgASgJUgJJRBISCgROYW1lGAIgASgJUgROYW1lEhoKCElzQWN0aXZlGAMgASgJUghJc0FjdGl2ZRISCgRNYWlsGAQgASgJUgRNYWlsEhgKB1N0b3JlSUQYBSABKAlSB1N0b3JlSUQ=');
@$core.Deprecated('Use adminDataResponseDescriptor instead')
const AdminDataResponse$json = const {
  '1': 'AdminDataResponse',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
    const {'1': 'Name', '3': 2, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'IsActive', '3': 3, '4': 1, '5': 9, '10': 'IsActive'},
    const {'1': 'Mail', '3': 4, '4': 1, '5': 9, '10': 'Mail'},
    const {'1': 'StoreID', '3': 5, '4': 1, '5': 9, '10': 'StoreID'},
  ],
};

/// Descriptor for `AdminDataResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List adminDataResponseDescriptor = $convert.base64Decode('ChFBZG1pbkRhdGFSZXNwb25zZRIOCgJJRBgBIAEoCVICSUQSEgoETmFtZRgCIAEoCVIETmFtZRIaCghJc0FjdGl2ZRgDIAEoCVIISXNBY3RpdmUSEgoETWFpbBgEIAEoCVIETWFpbBIYCgdTdG9yZUlEGAUgASgJUgdTdG9yZUlE');
const $core.Map<$core.String, $core.dynamic> AdminDataControllerServiceBase$json = const {
  '1': 'AdminDataController',
  '2': const [
    const {'1': 'Update', '2': '.server.user.AdminUpdateDataRequest', '3': '.server.user.AdminDataResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use adminDataControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> AdminDataControllerServiceBase$messageJson = const {
  '.server.user.AdminUpdateDataRequest': AdminUpdateDataRequest$json,
  '.server.user.AdminDataResponse': AdminDataResponse$json,
};

/// Descriptor for `AdminDataController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List adminDataControllerServiceDescriptor = $convert.base64Decode('ChNBZG1pbkRhdGFDb250cm9sbGVyEk8KBlVwZGF0ZRIjLnNlcnZlci51c2VyLkFkbWluVXBkYXRlRGF0YVJlcXVlc3QaHi5zZXJ2ZXIudXNlci5BZG1pbkRhdGFSZXNwb25zZSIA');

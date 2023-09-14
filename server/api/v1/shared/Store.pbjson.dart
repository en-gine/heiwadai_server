///
//  Generated code. Do not modify.
//  source: v1/shared/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use storeDescriptor instead')
const Store$json = const {
  '1': 'Store',
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
  ],
};

/// Descriptor for `Store`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storeDescriptor = $convert.base64Decode('CgVTdG9yZRIOCgJJRBgBIAEoCVICSUQSEgoETmFtZRgCIAEoCVIETmFtZRIeCgpCcmFuY2hOYW1lGAMgASgJUgpCcmFuY2hOYW1lEhgKB1ppcENvZGUYBCABKAlSB1ppcENvZGUSGAoHQWRkcmVzcxgFIAEoCVIHQWRkcmVzcxIQCgNUZWwYBiABKAlSA1RlbBIYCgdTaXRlVVJMGAcgASgJUgdTaXRlVVJMEiQKDVN0YW1wSW1hZ2VVUkwYCCABKAlSDVN0YW1wSW1hZ2VVUkwSGgoIU3RheWFibGUYCSABKAhSCFN0YXlhYmxlEhoKCElzQWN0aXZlGAogASgIUghJc0FjdGl2ZRIWCgZRUkNvZGUYCyABKAlSBlFSQ29kZRIoCg9VbkxpbWl0ZWRRUkNvZGUYDCABKAlSD1VuTGltaXRlZFFSQ29kZQ==');
@$core.Deprecated('Use storesDescriptor instead')
const Stores$json = const {
  '1': 'Stores',
  '2': const [
    const {'1': 'Stores', '3': 1, '4': 3, '5': 11, '6': '.server.shared.Store', '10': 'Stores'},
  ],
};

/// Descriptor for `Stores`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storesDescriptor = $convert.base64Decode('CgZTdG9yZXMSLAoGU3RvcmVzGAEgAygLMhQuc2VydmVyLnNoYXJlZC5TdG9yZVIGU3RvcmVz');
@$core.Deprecated('Use stayableStoreInfoDescriptor instead')
const StayableStoreInfo$json = const {
  '1': 'StayableStoreInfo',
  '2': const [
    const {'1': 'Parking', '3': 1, '4': 1, '5': 9, '10': 'Parking'},
    const {'1': 'Latitude', '3': 2, '4': 1, '5': 1, '10': 'Latitude'},
    const {'1': 'Longitude', '3': 3, '4': 1, '5': 1, '10': 'Longitude'},
    const {'1': 'AccessInfo', '3': 4, '4': 1, '5': 9, '10': 'AccessInfo'},
    const {'1': 'RestAPIURL', '3': 5, '4': 1, '5': 9, '10': 'RestAPIURL'},
    const {'1': 'BookingSystemID', '3': 6, '4': 1, '5': 9, '10': 'BookingSystemID'},
  ],
};

/// Descriptor for `StayableStoreInfo`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List stayableStoreInfoDescriptor = $convert.base64Decode('ChFTdGF5YWJsZVN0b3JlSW5mbxIYCgdQYXJraW5nGAEgASgJUgdQYXJraW5nEhoKCExhdGl0dWRlGAIgASgBUghMYXRpdHVkZRIcCglMb25naXR1ZGUYAyABKAFSCUxvbmdpdHVkZRIeCgpBY2Nlc3NJbmZvGAQgASgJUgpBY2Nlc3NJbmZvEh4KClJlc3RBUElVUkwYBSABKAlSClJlc3RBUElVUkwSKAoPQm9va2luZ1N5c3RlbUlEGAYgASgJUg9Cb29raW5nU3lzdGVtSUQ=');
@$core.Deprecated('Use stayableStoreDescriptor instead')
const StayableStore$json = const {
  '1': 'StayableStore',
  '2': const [
    const {'1': 'Store', '3': 1, '4': 1, '5': 11, '6': '.server.shared.Store', '10': 'Store'},
    const {'1': 'Info', '3': 2, '4': 1, '5': 11, '6': '.server.shared.StayableStoreInfo', '10': 'Info'},
  ],
};

/// Descriptor for `StayableStore`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List stayableStoreDescriptor = $convert.base64Decode('Cg1TdGF5YWJsZVN0b3JlEioKBVN0b3JlGAEgASgLMhQuc2VydmVyLnNoYXJlZC5TdG9yZVIFU3RvcmUSNAoESW5mbxgCIAEoCzIgLnNlcnZlci5zaGFyZWQuU3RheWFibGVTdG9yZUluZm9SBEluZm8=');
@$core.Deprecated('Use stayableStoresDescriptor instead')
const StayableStores$json = const {
  '1': 'StayableStores',
  '2': const [
    const {'1': 'StayableStores', '3': 1, '4': 3, '5': 11, '6': '.server.shared.StayableStore', '10': 'StayableStores'},
  ],
};

/// Descriptor for `StayableStores`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List stayableStoresDescriptor = $convert.base64Decode('Cg5TdGF5YWJsZVN0b3JlcxJECg5TdGF5YWJsZVN0b3JlcxgBIAMoCzIcLnNlcnZlci5zaGFyZWQuU3RheWFibGVTdG9yZVIOU3RheWFibGVTdG9yZXM=');

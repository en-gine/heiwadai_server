///
//  Generated code. Do not modify.
//  source: v1/user/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../shared/Store.pbjson.dart' as $7;
import '../../google/protobuf/empty.pbjson.dart' as $1;

@$core.Deprecated('Use soreIDRequestDescriptor instead')
const SoreIDRequest$json = const {
  '1': 'SoreIDRequest',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
  ],
};

/// Descriptor for `SoreIDRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List soreIDRequestDescriptor = $convert.base64Decode('Cg1Tb3JlSURSZXF1ZXN0Eg4KAklEGAEgASgJUgJJRA==');
const $core.Map<$core.String, $core.dynamic> StoreControllerServiceBase$json = const {
  '1': 'StoreController',
  '2': const [
    const {'1': 'GetByID', '2': '.server.user.SoreIDRequest', '3': '.server.shared.Store', '4': const {}},
    const {'1': 'GetAll', '2': '.google.protobuf.Empty', '3': '.server.shared.Stores', '4': const {}},
    const {'1': 'GetStayables', '2': '.google.protobuf.Empty', '3': '.server.shared.StayableStores', '4': const {}},
    const {'1': 'GetStayableByID', '2': '.server.user.SoreIDRequest', '3': '.server.shared.StayableStore', '4': const {}},
  ],
};

@$core.Deprecated('Use storeControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> StoreControllerServiceBase$messageJson = const {
  '.server.user.SoreIDRequest': SoreIDRequest$json,
  '.server.shared.Store': $7.Store$json,
  '.google.protobuf.Empty': $1.Empty$json,
  '.server.shared.Stores': $7.Stores$json,
  '.server.shared.StayableStores': $7.StayableStores$json,
  '.server.shared.StayableStore': $7.StayableStore$json,
  '.server.shared.StayableStoreInfo': $7.StayableStoreInfo$json,
};

/// Descriptor for `StoreController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List storeControllerServiceDescriptor = $convert.base64Decode('Cg9TdG9yZUNvbnRyb2xsZXISPQoHR2V0QnlJRBIaLnNlcnZlci51c2VyLlNvcmVJRFJlcXVlc3QaFC5zZXJ2ZXIuc2hhcmVkLlN0b3JlIgASOQoGR2V0QWxsEhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5GhUuc2VydmVyLnNoYXJlZC5TdG9yZXMiABJHCgxHZXRTdGF5YWJsZXMSFi5nb29nbGUucHJvdG9idWYuRW1wdHkaHS5zZXJ2ZXIuc2hhcmVkLlN0YXlhYmxlU3RvcmVzIgASTQoPR2V0U3RheWFibGVCeUlEEhouc2VydmVyLnVzZXIuU29yZUlEUmVxdWVzdBocLnNlcnZlci5zaGFyZWQuU3RheWFibGVTdG9yZSIA');

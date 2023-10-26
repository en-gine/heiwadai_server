///
//  Generated code. Do not modify.
//  source: v1/user/UserReport.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/empty.pbjson.dart' as $1;

@$core.Deprecated('Use userReportRequestDescriptor instead')
const UserReportRequest$json = const {
  '1': 'UserReportRequest',
  '2': const [
    const {'1': 'Title', '3': 2, '4': 1, '5': 9, '10': 'Title'},
    const {'1': 'Content', '3': 3, '4': 1, '5': 9, '10': 'Content'},
  ],
};

/// Descriptor for `UserReportRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userReportRequestDescriptor = $convert.base64Decode('ChFVc2VyUmVwb3J0UmVxdWVzdBIUCgVUaXRsZRgCIAEoCVIFVGl0bGUSGAoHQ29udGVudBgDIAEoCVIHQ29udGVudA==');
const $core.Map<$core.String, $core.dynamic> UserReportControllerServiceBase$json = const {
  '1': 'UserReportController',
  '2': const [
    const {'1': 'Send', '2': '.server.user.UserReportRequest', '3': '.google.protobuf.Empty', '4': const {}},
  ],
};

@$core.Deprecated('Use userReportControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> UserReportControllerServiceBase$messageJson = const {
  '.server.user.UserReportRequest': UserReportRequest$json,
  '.google.protobuf.Empty': $1.Empty$json,
};

/// Descriptor for `UserReportController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List userReportControllerServiceDescriptor = $convert.base64Decode('ChRVc2VyUmVwb3J0Q29udHJvbGxlchJACgRTZW5kEh4uc2VydmVyLnVzZXIuVXNlclJlcG9ydFJlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkiAA==');

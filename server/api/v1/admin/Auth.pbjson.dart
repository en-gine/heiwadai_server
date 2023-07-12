///
//  Generated code. Do not modify.
//  source: v1/admin/Auth.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use adminAuthRequestDescriptor instead')
const AdminAuthRequest$json = const {
  '1': 'AdminAuthRequest',
  '2': const [
    const {'1': 'email', '3': 1, '4': 1, '5': 9, '10': 'email'},
    const {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
  ],
};

/// Descriptor for `AdminAuthRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List adminAuthRequestDescriptor = $convert.base64Decode('ChBBZG1pbkF1dGhSZXF1ZXN0EhQKBWVtYWlsGAEgASgJUgVlbWFpbBIaCghwYXNzd29yZBgCIAEoCVIIcGFzc3dvcmQ=');
@$core.Deprecated('Use adminAuthResponseDescriptor instead')
const AdminAuthResponse$json = const {
  '1': 'AdminAuthResponse',
  '2': const [
    const {'1': 'accessToken', '3': 1, '4': 1, '5': 9, '10': 'accessToken'},
    const {'1': 'expiresIn', '3': 2, '4': 1, '5': 3, '10': 'expiresIn'},
    const {'1': 'refreshToken', '3': 3, '4': 1, '5': 9, '10': 'refreshToken'},
  ],
};

/// Descriptor for `AdminAuthResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List adminAuthResponseDescriptor = $convert.base64Decode('ChFBZG1pbkF1dGhSZXNwb25zZRIgCgthY2Nlc3NUb2tlbhgBIAEoCVILYWNjZXNzVG9rZW4SHAoJZXhwaXJlc0luGAIgASgDUglleHBpcmVzSW4SIgoMcmVmcmVzaFRva2VuGAMgASgJUgxyZWZyZXNoVG9rZW4=');
const $core.Map<$core.String, $core.dynamic> AuthControllerServiceBase$json = const {
  '1': 'AuthController',
  '2': const [
    const {'1': 'Call', '2': '.server.admin.AdminAuthRequest', '3': '.server.admin.AdminAuthResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use authControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> AuthControllerServiceBase$messageJson = const {
  '.server.admin.AdminAuthRequest': AdminAuthRequest$json,
  '.server.admin.AdminAuthResponse': AdminAuthResponse$json,
};

/// Descriptor for `AuthController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List authControllerServiceDescriptor = $convert.base64Decode('Cg5BdXRoQ29udHJvbGxlchJJCgRDYWxsEh4uc2VydmVyLmFkbWluLkFkbWluQXV0aFJlcXVlc3QaHy5zZXJ2ZXIuYWRtaW4uQWRtaW5BdXRoUmVzcG9uc2UiAA==');

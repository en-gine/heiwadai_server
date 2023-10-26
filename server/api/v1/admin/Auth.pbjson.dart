///
//  Generated code. Do not modify.
//  source: v1/admin/Auth.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/empty.pbjson.dart' as $1;

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
@$core.Deprecated('Use adminRegisterRequestDescriptor instead')
const AdminRegisterRequest$json = const {
  '1': 'AdminRegisterRequest',
  '2': const [
    const {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'Mail', '3': 2, '4': 1, '5': 9, '10': 'Mail'},
    const {'1': 'BelongStoreID', '3': 4, '4': 1, '5': 9, '10': 'BelongStoreID'},
  ],
};

/// Descriptor for `AdminRegisterRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List adminRegisterRequestDescriptor = $convert.base64Decode('ChRBZG1pblJlZ2lzdGVyUmVxdWVzdBISCgROYW1lGAEgASgJUgROYW1lEhIKBE1haWwYAiABKAlSBE1haWwSJAoNQmVsb25nU3RvcmVJRBgEIAEoCVINQmVsb25nU3RvcmVJRA==');
@$core.Deprecated('Use resetPasswordRequestDescriptor instead')
const ResetPasswordRequest$json = const {
  '1': 'ResetPasswordRequest',
  '2': const [
    const {'1': 'email', '3': 1, '4': 1, '5': 9, '10': 'email'},
  ],
};

/// Descriptor for `ResetPasswordRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List resetPasswordRequestDescriptor = $convert.base64Decode('ChRSZXNldFBhc3N3b3JkUmVxdWVzdBIUCgVlbWFpbBgBIAEoCVIFZW1haWw=');
@$core.Deprecated('Use updatePasswordRequestDescriptor instead')
const UpdatePasswordRequest$json = const {
  '1': 'UpdatePasswordRequest',
  '2': const [
    const {'1': 'password', '3': 1, '4': 1, '5': 9, '10': 'password'},
    const {'1': 'token', '3': 2, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `UpdatePasswordRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updatePasswordRequestDescriptor = $convert.base64Decode('ChVVcGRhdGVQYXNzd29yZFJlcXVlc3QSGgoIcGFzc3dvcmQYASABKAlSCHBhc3N3b3JkEhQKBXRva2VuGAIgASgJUgV0b2tlbg==');
@$core.Deprecated('Use updateEmailRequestDescriptor instead')
const UpdateEmailRequest$json = const {
  '1': 'UpdateEmailRequest',
  '2': const [
    const {'1': 'email', '3': 1, '4': 1, '5': 9, '10': 'email'},
    const {'1': 'token', '3': 2, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `UpdateEmailRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateEmailRequestDescriptor = $convert.base64Decode('ChJVcGRhdGVFbWFpbFJlcXVlc3QSFAoFZW1haWwYASABKAlSBWVtYWlsEhQKBXRva2VuGAIgASgJUgV0b2tlbg==');
const $core.Map<$core.String, $core.dynamic> AuthControllerServiceBase$json = const {
  '1': 'AuthController',
  '2': const [
    const {'1': 'Register', '2': '.server.admin.AdminRegisterRequest', '3': '.google.protobuf.Empty', '4': const {}},
    const {'1': 'SignUp', '2': '.server.admin.AdminAuthRequest', '3': '.google.protobuf.Empty', '4': const {}},
    const {'1': 'SignIn', '2': '.server.admin.AdminAuthRequest', '3': '.server.admin.AdminAuthResponse', '4': const {}},
    const {'1': 'ResetPasswordMail', '2': '.server.admin.ResetPasswordRequest', '3': '.google.protobuf.Empty', '4': const {}},
    const {'1': 'UpdatePassword', '2': '.server.admin.UpdatePasswordRequest', '3': '.google.protobuf.Empty', '4': const {}},
    const {'1': 'UpdateEmail', '2': '.server.admin.UpdateEmailRequest', '3': '.google.protobuf.Empty', '4': const {}},
  ],
};

@$core.Deprecated('Use authControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> AuthControllerServiceBase$messageJson = const {
  '.server.admin.AdminRegisterRequest': AdminRegisterRequest$json,
  '.google.protobuf.Empty': $1.Empty$json,
  '.server.admin.AdminAuthRequest': AdminAuthRequest$json,
  '.server.admin.AdminAuthResponse': AdminAuthResponse$json,
  '.server.admin.ResetPasswordRequest': ResetPasswordRequest$json,
  '.server.admin.UpdatePasswordRequest': UpdatePasswordRequest$json,
  '.server.admin.UpdateEmailRequest': UpdateEmailRequest$json,
};

/// Descriptor for `AuthController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List authControllerServiceDescriptor = $convert.base64Decode('Cg5BdXRoQ29udHJvbGxlchJICghSZWdpc3RlchIiLnNlcnZlci5hZG1pbi5BZG1pblJlZ2lzdGVyUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eSIAEkIKBlNpZ25VcBIeLnNlcnZlci5hZG1pbi5BZG1pbkF1dGhSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5IgASSwoGU2lnbkluEh4uc2VydmVyLmFkbWluLkFkbWluQXV0aFJlcXVlc3QaHy5zZXJ2ZXIuYWRtaW4uQWRtaW5BdXRoUmVzcG9uc2UiABJRChFSZXNldFBhc3N3b3JkTWFpbBIiLnNlcnZlci5hZG1pbi5SZXNldFBhc3N3b3JkUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eSIAEk8KDlVwZGF0ZVBhc3N3b3JkEiMuc2VydmVyLmFkbWluLlVwZGF0ZVBhc3N3b3JkUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eSIAEkkKC1VwZGF0ZUVtYWlsEiAuc2VydmVyLmFkbWluLlVwZGF0ZUVtYWlsUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eSIA');

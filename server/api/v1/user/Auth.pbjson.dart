///
//  Generated code. Do not modify.
//  source: v1/user/Auth.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/empty.pbjson.dart' as $2;

@$core.Deprecated('Use userRegisterRequestDescriptor instead')
const UserRegisterRequest$json = const {
  '1': 'UserRegisterRequest',
  '2': const [
    const {'1': 'FirstName', '3': 1, '4': 1, '5': 9, '10': 'FirstName'},
    const {'1': 'LastName', '3': 2, '4': 1, '5': 9, '10': 'LastName'},
    const {'1': 'FirstNameKana', '3': 3, '4': 1, '5': 9, '10': 'FirstNameKana'},
    const {'1': 'LastNameKana', '3': 4, '4': 1, '5': 9, '10': 'LastNameKana'},
    const {'1': 'CompanyName', '3': 5, '4': 1, '5': 9, '10': 'CompanyName'},
    const {'1': 'BirthDate', '3': 6, '4': 1, '5': 9, '10': 'BirthDate'},
    const {'1': 'ZipCode', '3': 7, '4': 1, '5': 9, '10': 'ZipCode'},
    const {'1': 'Prefecture', '3': 8, '4': 1, '5': 9, '10': 'Prefecture'},
    const {'1': 'City', '3': 9, '4': 1, '5': 9, '10': 'City'},
    const {'1': 'Address', '3': 10, '4': 1, '5': 9, '10': 'Address'},
    const {'1': 'Tel', '3': 11, '4': 1, '5': 9, '10': 'Tel'},
    const {'1': 'Mail', '3': 12, '4': 1, '5': 9, '10': 'Mail'},
    const {'1': 'AcceptMail', '3': 13, '4': 1, '5': 8, '10': 'AcceptMail'},
    const {'1': 'AcceptTerm', '3': 14, '4': 1, '5': 8, '10': 'AcceptTerm'},
  ],
};

/// Descriptor for `UserRegisterRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userRegisterRequestDescriptor = $convert.base64Decode('ChNVc2VyUmVnaXN0ZXJSZXF1ZXN0EhwKCUZpcnN0TmFtZRgBIAEoCVIJRmlyc3ROYW1lEhoKCExhc3ROYW1lGAIgASgJUghMYXN0TmFtZRIkCg1GaXJzdE5hbWVLYW5hGAMgASgJUg1GaXJzdE5hbWVLYW5hEiIKDExhc3ROYW1lS2FuYRgEIAEoCVIMTGFzdE5hbWVLYW5hEiAKC0NvbXBhbnlOYW1lGAUgASgJUgtDb21wYW55TmFtZRIcCglCaXJ0aERhdGUYBiABKAlSCUJpcnRoRGF0ZRIYCgdaaXBDb2RlGAcgASgJUgdaaXBDb2RlEh4KClByZWZlY3R1cmUYCCABKAlSClByZWZlY3R1cmUSEgoEQ2l0eRgJIAEoCVIEQ2l0eRIYCgdBZGRyZXNzGAogASgJUgdBZGRyZXNzEhAKA1RlbBgLIAEoCVIDVGVsEhIKBE1haWwYDCABKAlSBE1haWwSHgoKQWNjZXB0TWFpbBgNIAEoCFIKQWNjZXB0TWFpbBIeCgpBY2NlcHRUZXJtGA4gASgIUgpBY2NlcHRUZXJt');
@$core.Deprecated('Use userAuthRequestDescriptor instead')
const UserAuthRequest$json = const {
  '1': 'UserAuthRequest',
  '2': const [
    const {'1': 'email', '3': 1, '4': 1, '5': 9, '10': 'email'},
    const {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
  ],
};

/// Descriptor for `UserAuthRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userAuthRequestDescriptor = $convert.base64Decode('Cg9Vc2VyQXV0aFJlcXVlc3QSFAoFZW1haWwYASABKAlSBWVtYWlsEhoKCHBhc3N3b3JkGAIgASgJUghwYXNzd29yZA==');
@$core.Deprecated('Use userAuthResponseDescriptor instead')
const UserAuthResponse$json = const {
  '1': 'UserAuthResponse',
  '2': const [
    const {'1': 'accessToken', '3': 1, '4': 1, '5': 9, '10': 'accessToken'},
    const {'1': 'expiresIn', '3': 2, '4': 1, '5': 3, '10': 'expiresIn'},
    const {'1': 'refreshToken', '3': 3, '4': 1, '5': 9, '10': 'refreshToken'},
  ],
};

/// Descriptor for `UserAuthResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userAuthResponseDescriptor = $convert.base64Decode('ChBVc2VyQXV0aFJlc3BvbnNlEiAKC2FjY2Vzc1Rva2VuGAEgASgJUgthY2Nlc3NUb2tlbhIcCglleHBpcmVzSW4YAiABKANSCWV4cGlyZXNJbhIiCgxyZWZyZXNoVG9rZW4YAyABKAlSDHJlZnJlc2hUb2tlbg==');
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
    const {'1': 'Register', '2': '.server.user.UserRegisterRequest', '3': '.google.protobuf.Empty', '4': const {}},
    const {'1': 'SignUp', '2': '.server.user.UserAuthRequest', '3': '.google.protobuf.Empty', '4': const {}},
    const {'1': 'SignIn', '2': '.server.user.UserAuthRequest', '3': '.server.user.UserAuthResponse', '4': const {}},
    const {'1': 'ResetPasswordMail', '2': '.server.user.ResetPasswordRequest', '3': '.google.protobuf.Empty', '4': const {}},
    const {'1': 'UpdatePassword', '2': '.server.user.UpdatePasswordRequest', '3': '.google.protobuf.Empty', '4': const {}},
    const {'1': 'UpdateEmail', '2': '.server.user.UpdateEmailRequest', '3': '.google.protobuf.Empty', '4': const {}},
  ],
};

@$core.Deprecated('Use authControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> AuthControllerServiceBase$messageJson = const {
  '.server.user.UserRegisterRequest': UserRegisterRequest$json,
  '.google.protobuf.Empty': $2.Empty$json,
  '.server.user.UserAuthRequest': UserAuthRequest$json,
  '.server.user.UserAuthResponse': UserAuthResponse$json,
  '.server.user.ResetPasswordRequest': ResetPasswordRequest$json,
  '.server.user.UpdatePasswordRequest': UpdatePasswordRequest$json,
  '.server.user.UpdateEmailRequest': UpdateEmailRequest$json,
};

/// Descriptor for `AuthController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List authControllerServiceDescriptor = $convert.base64Decode('Cg5BdXRoQ29udHJvbGxlchJGCghSZWdpc3RlchIgLnNlcnZlci51c2VyLlVzZXJSZWdpc3RlclJlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkiABJACgZTaWduVXASHC5zZXJ2ZXIudXNlci5Vc2VyQXV0aFJlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkiABJHCgZTaWduSW4SHC5zZXJ2ZXIudXNlci5Vc2VyQXV0aFJlcXVlc3QaHS5zZXJ2ZXIudXNlci5Vc2VyQXV0aFJlc3BvbnNlIgASUAoRUmVzZXRQYXNzd29yZE1haWwSIS5zZXJ2ZXIudXNlci5SZXNldFBhc3N3b3JkUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eSIAEk4KDlVwZGF0ZVBhc3N3b3JkEiIuc2VydmVyLnVzZXIuVXBkYXRlUGFzc3dvcmRSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5IgASSAoLVXBkYXRlRW1haWwSHy5zZXJ2ZXIudXNlci5VcGRhdGVFbWFpbFJlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkiAA==');

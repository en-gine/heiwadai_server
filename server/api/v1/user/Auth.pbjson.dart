///
//  Generated code. Do not modify.
//  source: v1/user/Auth.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
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
    const {'1': 'isUserAcceptable', '3': 4, '4': 1, '5': 8, '10': 'isUserAcceptable'},
  ],
};

/// Descriptor for `UserAuthResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userAuthResponseDescriptor = $convert.base64Decode('ChBVc2VyQXV0aFJlc3BvbnNlEiAKC2FjY2Vzc1Rva2VuGAEgASgJUgthY2Nlc3NUb2tlbhIcCglleHBpcmVzSW4YAiABKANSCWV4cGlyZXNJbhIiCgxyZWZyZXNoVG9rZW4YAyABKAlSDHJlZnJlc2hUb2tlbhIqChBpc1VzZXJBY2NlcHRhYmxlGAQgASgIUhBpc1VzZXJBY2NlcHRhYmxl');
const $core.Map<$core.String, $core.dynamic> AuthControllerServiceBase$json = const {
  '1': 'AuthController',
  '2': const [
    const {'1': 'Call', '2': '.server.user.UserAuthRequest', '3': '.server.user.UserAuthResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use authControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> AuthControllerServiceBase$messageJson = const {
  '.server.user.UserAuthRequest': UserAuthRequest$json,
  '.server.user.UserAuthResponse': UserAuthResponse$json,
};

/// Descriptor for `AuthController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List authControllerServiceDescriptor = $convert.base64Decode('Cg5BdXRoQ29udHJvbGxlchJFCgRDYWxsEhwuc2VydmVyLnVzZXIuVXNlckF1dGhSZXF1ZXN0Gh0uc2VydmVyLnVzZXIuVXNlckF1dGhSZXNwb25zZSIA');

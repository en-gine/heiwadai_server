///
//  Generated code. Do not modify.
//  source: v1/user/UserData.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
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
@$core.Deprecated('Use userUpdateDataRequestDescriptor instead')
const UserUpdateDataRequest$json = const {
  '1': 'UserUpdateDataRequest',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
    const {'1': 'FirstName', '3': 2, '4': 1, '5': 9, '10': 'FirstName'},
    const {'1': 'LastName', '3': 3, '4': 1, '5': 9, '10': 'LastName'},
    const {'1': 'FirstNameKana', '3': 4, '4': 1, '5': 9, '10': 'FirstNameKana'},
    const {'1': 'LastNameKana', '3': 5, '4': 1, '5': 9, '10': 'LastNameKana'},
    const {'1': 'CompanyName', '3': 6, '4': 1, '5': 9, '10': 'CompanyName'},
    const {'1': 'BirthDate', '3': 7, '4': 1, '5': 9, '10': 'BirthDate'},
    const {'1': 'ZipCode', '3': 8, '4': 1, '5': 9, '10': 'ZipCode'},
    const {'1': 'Prefecture', '3': 9, '4': 1, '5': 9, '10': 'Prefecture'},
    const {'1': 'City', '3': 10, '4': 1, '5': 9, '10': 'City'},
    const {'1': 'Address', '3': 11, '4': 1, '5': 9, '10': 'Address'},
    const {'1': 'Tel', '3': 12, '4': 1, '5': 9, '10': 'Tel'},
    const {'1': 'Mail', '3': 13, '4': 1, '5': 9, '10': 'Mail'},
    const {'1': 'AcceptMail', '3': 14, '4': 1, '5': 8, '10': 'AcceptMail'},
  ],
};

/// Descriptor for `UserUpdateDataRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userUpdateDataRequestDescriptor = $convert.base64Decode('ChVVc2VyVXBkYXRlRGF0YVJlcXVlc3QSDgoCSUQYASABKAlSAklEEhwKCUZpcnN0TmFtZRgCIAEoCVIJRmlyc3ROYW1lEhoKCExhc3ROYW1lGAMgASgJUghMYXN0TmFtZRIkCg1GaXJzdE5hbWVLYW5hGAQgASgJUg1GaXJzdE5hbWVLYW5hEiIKDExhc3ROYW1lS2FuYRgFIAEoCVIMTGFzdE5hbWVLYW5hEiAKC0NvbXBhbnlOYW1lGAYgASgJUgtDb21wYW55TmFtZRIcCglCaXJ0aERhdGUYByABKAlSCUJpcnRoRGF0ZRIYCgdaaXBDb2RlGAggASgJUgdaaXBDb2RlEh4KClByZWZlY3R1cmUYCSABKAlSClByZWZlY3R1cmUSEgoEQ2l0eRgKIAEoCVIEQ2l0eRIYCgdBZGRyZXNzGAsgASgJUgdBZGRyZXNzEhAKA1RlbBgMIAEoCVIDVGVsEhIKBE1haWwYDSABKAlSBE1haWwSHgoKQWNjZXB0TWFpbBgOIAEoCFIKQWNjZXB0TWFpbA==');
@$core.Deprecated('Use userDataResponseDescriptor instead')
const UserDataResponse$json = const {
  '1': 'UserDataResponse',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
    const {'1': 'FirstName', '3': 2, '4': 1, '5': 9, '10': 'FirstName'},
    const {'1': 'LastName', '3': 3, '4': 1, '5': 9, '10': 'LastName'},
    const {'1': 'FirstNameKana', '3': 4, '4': 1, '5': 9, '10': 'FirstNameKana'},
    const {'1': 'LastNameKana', '3': 5, '4': 1, '5': 9, '10': 'LastNameKana'},
    const {'1': 'CompanyName', '3': 6, '4': 1, '5': 9, '10': 'CompanyName'},
    const {'1': 'BirthDate', '3': 7, '4': 1, '5': 9, '10': 'BirthDate'},
    const {'1': 'ZipCode', '3': 8, '4': 1, '5': 9, '10': 'ZipCode'},
    const {'1': 'Prefecture', '3': 9, '4': 1, '5': 9, '10': 'Prefecture'},
    const {'1': 'City', '3': 10, '4': 1, '5': 9, '10': 'City'},
    const {'1': 'Address', '3': 11, '4': 1, '5': 9, '10': 'Address'},
    const {'1': 'Tel', '3': 12, '4': 1, '5': 9, '10': 'Tel'},
    const {'1': 'Mail', '3': 13, '4': 1, '5': 9, '10': 'Mail'},
    const {'1': 'AcceptMail', '3': 14, '4': 1, '5': 8, '10': 'AcceptMail'},
  ],
};

/// Descriptor for `UserDataResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userDataResponseDescriptor = $convert.base64Decode('ChBVc2VyRGF0YVJlc3BvbnNlEg4KAklEGAEgASgJUgJJRBIcCglGaXJzdE5hbWUYAiABKAlSCUZpcnN0TmFtZRIaCghMYXN0TmFtZRgDIAEoCVIITGFzdE5hbWUSJAoNRmlyc3ROYW1lS2FuYRgEIAEoCVINRmlyc3ROYW1lS2FuYRIiCgxMYXN0TmFtZUthbmEYBSABKAlSDExhc3ROYW1lS2FuYRIgCgtDb21wYW55TmFtZRgGIAEoCVILQ29tcGFueU5hbWUSHAoJQmlydGhEYXRlGAcgASgJUglCaXJ0aERhdGUSGAoHWmlwQ29kZRgIIAEoCVIHWmlwQ29kZRIeCgpQcmVmZWN0dXJlGAkgASgJUgpQcmVmZWN0dXJlEhIKBENpdHkYCiABKAlSBENpdHkSGAoHQWRkcmVzcxgLIAEoCVIHQWRkcmVzcxIQCgNUZWwYDCABKAlSA1RlbBISCgRNYWlsGA0gASgJUgRNYWlsEh4KCkFjY2VwdE1haWwYDiABKAhSCkFjY2VwdE1haWw=');
const $core.Map<$core.String, $core.dynamic> UserDataControllerServiceBase$json = const {
  '1': 'UserDataController',
  '2': const [
    const {'1': 'Register', '2': '.server.user.UserRegisterRequest', '3': '.server.user.UserDataResponse', '4': const {}},
    const {'1': 'Update', '2': '.server.user.UserUpdateDataRequest', '3': '.server.user.UserDataResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use userDataControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> UserDataControllerServiceBase$messageJson = const {
  '.server.user.UserRegisterRequest': UserRegisterRequest$json,
  '.server.user.UserDataResponse': UserDataResponse$json,
  '.server.user.UserUpdateDataRequest': UserUpdateDataRequest$json,
};

/// Descriptor for `UserDataController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List userDataControllerServiceDescriptor = $convert.base64Decode('ChJVc2VyRGF0YUNvbnRyb2xsZXISTQoIUmVnaXN0ZXISIC5zZXJ2ZXIudXNlci5Vc2VyUmVnaXN0ZXJSZXF1ZXN0Gh0uc2VydmVyLnVzZXIuVXNlckRhdGFSZXNwb25zZSIAEk0KBlVwZGF0ZRIiLnNlcnZlci51c2VyLlVzZXJVcGRhdGVEYXRhUmVxdWVzdBodLnNlcnZlci51c2VyLlVzZXJEYXRhUmVzcG9uc2UiAA==');

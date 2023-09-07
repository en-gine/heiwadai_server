///
//  Generated code. Do not modify.
//  source: v1/user/Banner.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/empty.pbjson.dart' as $3;

@$core.Deprecated('Use bannerDescriptor instead')
const Banner$json = const {
  '1': 'Banner',
  '2': const [
    const {'1': 'ImageURL', '3': 1, '4': 1, '5': 9, '10': 'ImageURL'},
    const {'1': 'URL', '3': 2, '4': 1, '5': 9, '10': 'URL'},
  ],
};

/// Descriptor for `Banner`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List bannerDescriptor = $convert.base64Decode('CgZCYW5uZXISGgoISW1hZ2VVUkwYASABKAlSCEltYWdlVVJMEhAKA1VSTBgCIAEoCVIDVVJM');
@$core.Deprecated('Use bannerResponseDescriptor instead')
const BannerResponse$json = const {
  '1': 'BannerResponse',
  '2': const [
    const {'1': 'banners', '3': 1, '4': 3, '5': 11, '6': '.server.user.Banner', '10': 'banners'},
  ],
};

/// Descriptor for `BannerResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List bannerResponseDescriptor = $convert.base64Decode('Cg5CYW5uZXJSZXNwb25zZRItCgdiYW5uZXJzGAEgAygLMhMuc2VydmVyLnVzZXIuQmFubmVyUgdiYW5uZXJz');
const $core.Map<$core.String, $core.dynamic> BannerControllerServiceBase$json = const {
  '1': 'BannerController',
  '2': const [
    const {'1': 'GetBanner', '2': '.google.protobuf.Empty', '3': '.server.user.BannerResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use bannerControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> BannerControllerServiceBase$messageJson = const {
  '.google.protobuf.Empty': $3.Empty$json,
  '.server.user.BannerResponse': BannerResponse$json,
  '.server.user.Banner': Banner$json,
};

/// Descriptor for `BannerController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List bannerControllerServiceDescriptor = $convert.base64Decode('ChBCYW5uZXJDb250cm9sbGVyEkIKCUdldEJhbm5lchIWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRobLnNlcnZlci51c2VyLkJhbm5lclJlc3BvbnNlIgA=');

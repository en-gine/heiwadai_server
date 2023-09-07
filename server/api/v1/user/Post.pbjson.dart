///
//  Generated code. Do not modify.
//  source: v1/user/Post.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/empty.pbjson.dart' as $3;

@$core.Deprecated('Use postRequestDescriptor instead')
const PostRequest$json = const {
  '1': 'PostRequest',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 13, '10': 'ID'},
  ],
};

/// Descriptor for `PostRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List postRequestDescriptor = $convert.base64Decode('CgtQb3N0UmVxdWVzdBIOCgJJRBgBIAEoDVICSUQ=');
@$core.Deprecated('Use postsResponseDescriptor instead')
const PostsResponse$json = const {
  '1': 'PostsResponse',
  '2': const [
    const {'1': 'posts', '3': 1, '4': 3, '5': 11, '6': '.server.user.PostResponse', '10': 'posts'},
  ],
};

/// Descriptor for `PostsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List postsResponseDescriptor = $convert.base64Decode('Cg1Qb3N0c1Jlc3BvbnNlEi8KBXBvc3RzGAEgAygLMhkuc2VydmVyLnVzZXIuUG9zdFJlc3BvbnNlUgVwb3N0cw==');
@$core.Deprecated('Use postResponseDescriptor instead')
const PostResponse$json = const {
  '1': 'PostResponse',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 13, '10': 'ID'},
    const {'1': 'Title', '3': 2, '4': 1, '5': 9, '10': 'Title'},
    const {'1': 'Content', '3': 3, '4': 1, '5': 9, '10': 'Content'},
    const {'1': 'Author', '3': 4, '4': 1, '5': 9, '10': 'Author'},
    const {'1': 'PostDate', '3': 5, '4': 1, '5': 9, '10': 'PostDate'},
  ],
};

/// Descriptor for `PostResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List postResponseDescriptor = $convert.base64Decode('CgxQb3N0UmVzcG9uc2USDgoCSUQYASABKA1SAklEEhQKBVRpdGxlGAIgASgJUgVUaXRsZRIYCgdDb250ZW50GAMgASgJUgdDb250ZW50EhYKBkF1dGhvchgEIAEoCVIGQXV0aG9yEhoKCFBvc3REYXRlGAUgASgJUghQb3N0RGF0ZQ==');
const $core.Map<$core.String, $core.dynamic> PostControllerServiceBase$json = const {
  '1': 'PostController',
  '2': const [
    const {'1': 'GetPostByID', '2': '.server.user.PostRequest', '3': '.server.user.PostResponse', '4': const {}},
    const {'1': 'GetPosts', '2': '.google.protobuf.Empty', '3': '.server.user.PostsResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use postControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> PostControllerServiceBase$messageJson = const {
  '.server.user.PostRequest': PostRequest$json,
  '.server.user.PostResponse': PostResponse$json,
  '.google.protobuf.Empty': $3.Empty$json,
  '.server.user.PostsResponse': PostsResponse$json,
};

/// Descriptor for `PostController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List postControllerServiceDescriptor = $convert.base64Decode('Cg5Qb3N0Q29udHJvbGxlchJECgtHZXRQb3N0QnlJRBIYLnNlcnZlci51c2VyLlBvc3RSZXF1ZXN0Ghkuc2VydmVyLnVzZXIuUG9zdFJlc3BvbnNlIgASQAoIR2V0UG9zdHMSFi5nb29nbGUucHJvdG9idWYuRW1wdHkaGi5zZXJ2ZXIudXNlci5Qb3N0c1Jlc3BvbnNlIgA=');

///
//  Generated code. Do not modify.
//  source: v1/user/Messages.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/timestamp.pbjson.dart' as $5;

@$core.Deprecated('Use messageRequestDescriptor instead')
const MessageRequest$json = const {
  '1': 'MessageRequest',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
  ],
};

/// Descriptor for `MessageRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageRequestDescriptor = $convert.base64Decode('Cg5NZXNzYWdlUmVxdWVzdBIOCgJJRBgBIAEoCVICSUQ=');
@$core.Deprecated('Use messagesResponseDescriptor instead')
const MessagesResponse$json = const {
  '1': 'MessagesResponse',
  '2': const [
    const {'1': 'messages', '3': 1, '4': 3, '5': 11, '6': '.server.user.MessageResponse', '10': 'messages'},
  ],
};

/// Descriptor for `MessagesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messagesResponseDescriptor = $convert.base64Decode('ChBNZXNzYWdlc1Jlc3BvbnNlEjgKCG1lc3NhZ2VzGAEgAygLMhwuc2VydmVyLnVzZXIuTWVzc2FnZVJlc3BvbnNlUghtZXNzYWdlcw==');
@$core.Deprecated('Use messageResponseDescriptor instead')
const MessageResponse$json = const {
  '1': 'MessageResponse',
  '2': const [
    const {'1': 'ID', '3': 1, '4': 1, '5': 9, '10': 'ID'},
    const {'1': 'Title', '3': 2, '4': 1, '5': 9, '10': 'Title'},
    const {'1': 'Content', '3': 3, '4': 1, '5': 9, '10': 'Content'},
    const {'1': 'AuthorID', '3': 4, '4': 1, '5': 9, '10': 'AuthorID'},
    const {'1': 'DisplayDate', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'DisplayDate'},
  ],
};

/// Descriptor for `MessageResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageResponseDescriptor = $convert.base64Decode('Cg9NZXNzYWdlUmVzcG9uc2USDgoCSUQYASABKAlSAklEEhQKBVRpdGxlGAIgASgJUgVUaXRsZRIYCgdDb250ZW50GAMgASgJUgdDb250ZW50EhoKCEF1dGhvcklEGAQgASgJUghBdXRob3JJRBI8CgtEaXNwbGF5RGF0ZRgFIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSC0Rpc3BsYXlEYXRl');
const $core.Map<$core.String, $core.dynamic> MessageControllerServiceBase$json = const {
  '1': 'MessageController',
  '2': const [
    const {'1': 'GetMessagesAfter', '2': '.server.user.MessageRequest', '3': '.server.user.MessagesResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use messageControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> MessageControllerServiceBase$messageJson = const {
  '.server.user.MessageRequest': MessageRequest$json,
  '.server.user.MessagesResponse': MessagesResponse$json,
  '.server.user.MessageResponse': MessageResponse$json,
  '.google.protobuf.Timestamp': $5.Timestamp$json,
};

/// Descriptor for `MessageController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List messageControllerServiceDescriptor = $convert.base64Decode('ChFNZXNzYWdlQ29udHJvbGxlchJQChBHZXRNZXNzYWdlc0FmdGVyEhsuc2VydmVyLnVzZXIuTWVzc2FnZVJlcXVlc3QaHS5zZXJ2ZXIudXNlci5NZXNzYWdlc1Jlc3BvbnNlIgA=');

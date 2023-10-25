///
//  Generated code. Do not modify.
//  source: v1/user/Reservation.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/empty.pbjson.dart' as $2;

@$core.Deprecated('Use reservationDescriptor instead')
const Reservation$json = const {
  '1': 'Reservation',
};

/// Descriptor for `Reservation`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List reservationDescriptor = $convert.base64Decode('CgtSZXNlcnZhdGlvbg==');
@$core.Deprecated('Use reservationResponseDescriptor instead')
const ReservationResponse$json = const {
  '1': 'ReservationResponse',
};

/// Descriptor for `ReservationResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List reservationResponseDescriptor = $convert.base64Decode('ChNSZXNlcnZhdGlvblJlc3BvbnNl');
const $core.Map<$core.String, $core.dynamic> ReservationControllerServiceBase$json = const {
  '1': 'ReservationController',
  '2': const [
    const {'1': 'GetReservation', '2': '.google.protobuf.Empty', '3': '.server.user.ReservationResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use reservationControllerServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> ReservationControllerServiceBase$messageJson = const {
  '.google.protobuf.Empty': $2.Empty$json,
  '.server.user.ReservationResponse': ReservationResponse$json,
};

/// Descriptor for `ReservationController`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List reservationControllerServiceDescriptor = $convert.base64Decode('ChVSZXNlcnZhdGlvbkNvbnRyb2xsZXISTAoOR2V0UmVzZXJ2YXRpb24SFi5nb29nbGUucHJvdG9idWYuRW1wdHkaIC5zZXJ2ZXIudXNlci5SZXNlcnZhdGlvblJlc3BvbnNlIgA=');

///
//  Generated code. Do not modify.
//  source: v1/user/Auth.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'Auth.pb.dart' as $4;
import '../../google/protobuf/empty.pb.dart' as $1;
import 'Auth.pbjson.dart';

export 'Auth.pb.dart';

abstract class AuthControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$1.Empty> register($pb.ServerContext ctx, $4.UserRegisterRequest request);
  $async.Future<$1.Empty> signUp($pb.ServerContext ctx, $4.UserAuthRequest request);
  $async.Future<$4.UserAuthResponse> signIn($pb.ServerContext ctx, $4.UserAuthRequest request);
  $async.Future<$1.Empty> resetPasswordMail($pb.ServerContext ctx, $4.ResetPasswordRequest request);
  $async.Future<$1.Empty> updatePassword($pb.ServerContext ctx, $4.UpdatePasswordRequest request);
  $async.Future<$1.Empty> updateEmail($pb.ServerContext ctx, $4.UpdateEmailRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Register': return $4.UserRegisterRequest();
      case 'SignUp': return $4.UserAuthRequest();
      case 'SignIn': return $4.UserAuthRequest();
      case 'ResetPasswordMail': return $4.ResetPasswordRequest();
      case 'UpdatePassword': return $4.UpdatePasswordRequest();
      case 'UpdateEmail': return $4.UpdateEmailRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Register': return this.register(ctx, request as $4.UserRegisterRequest);
      case 'SignUp': return this.signUp(ctx, request as $4.UserAuthRequest);
      case 'SignIn': return this.signIn(ctx, request as $4.UserAuthRequest);
      case 'ResetPasswordMail': return this.resetPasswordMail(ctx, request as $4.ResetPasswordRequest);
      case 'UpdatePassword': return this.updatePassword(ctx, request as $4.UpdatePasswordRequest);
      case 'UpdateEmail': return this.updateEmail(ctx, request as $4.UpdateEmailRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => AuthControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => AuthControllerServiceBase$messageJson;
}


///
//  Generated code. Do not modify.
//  source: v1/user/Auth.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'Auth.pb.dart' as $6;
import '../../google/protobuf/empty.pb.dart' as $2;
import 'Auth.pbjson.dart';

export 'Auth.pb.dart';

abstract class AuthControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$2.Empty> register($pb.ServerContext ctx, $6.UserRegisterRequest request);
  $async.Future<$2.Empty> signUp($pb.ServerContext ctx, $6.UserAuthRequest request);
  $async.Future<$6.UserAuthResponse> signIn($pb.ServerContext ctx, $6.UserAuthRequest request);
  $async.Future<$2.Empty> resetPasswordMail($pb.ServerContext ctx, $6.ResetPasswordRequest request);
  $async.Future<$2.Empty> updatePassword($pb.ServerContext ctx, $6.UpdatePasswordRequest request);
  $async.Future<$2.Empty> updateEmail($pb.ServerContext ctx, $6.UpdateEmailRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Register': return $6.UserRegisterRequest();
      case 'SignUp': return $6.UserAuthRequest();
      case 'SignIn': return $6.UserAuthRequest();
      case 'ResetPasswordMail': return $6.ResetPasswordRequest();
      case 'UpdatePassword': return $6.UpdatePasswordRequest();
      case 'UpdateEmail': return $6.UpdateEmailRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Register': return this.register(ctx, request as $6.UserRegisterRequest);
      case 'SignUp': return this.signUp(ctx, request as $6.UserAuthRequest);
      case 'SignIn': return this.signIn(ctx, request as $6.UserAuthRequest);
      case 'ResetPasswordMail': return this.resetPasswordMail(ctx, request as $6.ResetPasswordRequest);
      case 'UpdatePassword': return this.updatePassword(ctx, request as $6.UpdatePasswordRequest);
      case 'UpdateEmail': return this.updateEmail(ctx, request as $6.UpdateEmailRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => AuthControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => AuthControllerServiceBase$messageJson;
}


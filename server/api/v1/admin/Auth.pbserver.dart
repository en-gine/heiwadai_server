///
//  Generated code. Do not modify.
//  source: v1/admin/Auth.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'Auth.pb.dart' as $2;
import '../../google/protobuf/empty.pb.dart' as $1;
import 'Auth.pbjson.dart';

export 'Auth.pb.dart';

abstract class AuthControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$1.Empty> register($pb.ServerContext ctx, $2.AdminRegisterRequest request);
  $async.Future<$1.Empty> signUp($pb.ServerContext ctx, $2.AdminAuthRequest request);
  $async.Future<$2.AdminAuthResponse> signIn($pb.ServerContext ctx, $2.AdminAuthRequest request);
  $async.Future<$1.Empty> resetPasswordMail($pb.ServerContext ctx, $2.ResetPasswordRequest request);
  $async.Future<$1.Empty> updatePassword($pb.ServerContext ctx, $2.UpdatePasswordRequest request);
  $async.Future<$1.Empty> updateEmail($pb.ServerContext ctx, $2.UpdateEmailRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Register': return $2.AdminRegisterRequest();
      case 'SignUp': return $2.AdminAuthRequest();
      case 'SignIn': return $2.AdminAuthRequest();
      case 'ResetPasswordMail': return $2.ResetPasswordRequest();
      case 'UpdatePassword': return $2.UpdatePasswordRequest();
      case 'UpdateEmail': return $2.UpdateEmailRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Register': return this.register(ctx, request as $2.AdminRegisterRequest);
      case 'SignUp': return this.signUp(ctx, request as $2.AdminAuthRequest);
      case 'SignIn': return this.signIn(ctx, request as $2.AdminAuthRequest);
      case 'ResetPasswordMail': return this.resetPasswordMail(ctx, request as $2.ResetPasswordRequest);
      case 'UpdatePassword': return this.updatePassword(ctx, request as $2.UpdatePasswordRequest);
      case 'UpdateEmail': return this.updateEmail(ctx, request as $2.UpdateEmailRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => AuthControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => AuthControllerServiceBase$messageJson;
}


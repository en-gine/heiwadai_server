///
//  Generated code. Do not modify.
//  source: v1/user/Auth.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/empty.pb.dart' as $1;

class UserRegisterRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserRegisterRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FirstName', protoName: 'FirstName')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'LastName', protoName: 'LastName')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FirstNameKana', protoName: 'FirstNameKana')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'LastNameKana', protoName: 'LastNameKana')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'CompanyName', protoName: 'CompanyName')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'BirthDate', protoName: 'BirthDate')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ZipCode', protoName: 'ZipCode')
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Prefecture', protoName: 'Prefecture')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'City', protoName: 'City')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Address', protoName: 'Address')
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Tel', protoName: 'Tel')
    ..aOS(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Mail', protoName: 'Mail')
    ..aOB(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AcceptMail', protoName: 'AcceptMail')
    ..aOB(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AcceptTerm', protoName: 'AcceptTerm')
    ..hasRequiredFields = false
  ;

  UserRegisterRequest._() : super();
  factory UserRegisterRequest({
    $core.String? firstName,
    $core.String? lastName,
    $core.String? firstNameKana,
    $core.String? lastNameKana,
    $core.String? companyName,
    $core.String? birthDate,
    $core.String? zipCode,
    $core.String? prefecture,
    $core.String? city,
    $core.String? address,
    $core.String? tel,
    $core.String? mail,
    $core.bool? acceptMail,
    $core.bool? acceptTerm,
  }) {
    final _result = create();
    if (firstName != null) {
      _result.firstName = firstName;
    }
    if (lastName != null) {
      _result.lastName = lastName;
    }
    if (firstNameKana != null) {
      _result.firstNameKana = firstNameKana;
    }
    if (lastNameKana != null) {
      _result.lastNameKana = lastNameKana;
    }
    if (companyName != null) {
      _result.companyName = companyName;
    }
    if (birthDate != null) {
      _result.birthDate = birthDate;
    }
    if (zipCode != null) {
      _result.zipCode = zipCode;
    }
    if (prefecture != null) {
      _result.prefecture = prefecture;
    }
    if (city != null) {
      _result.city = city;
    }
    if (address != null) {
      _result.address = address;
    }
    if (tel != null) {
      _result.tel = tel;
    }
    if (mail != null) {
      _result.mail = mail;
    }
    if (acceptMail != null) {
      _result.acceptMail = acceptMail;
    }
    if (acceptTerm != null) {
      _result.acceptTerm = acceptTerm;
    }
    return _result;
  }
  factory UserRegisterRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserRegisterRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserRegisterRequest clone() => UserRegisterRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserRegisterRequest copyWith(void Function(UserRegisterRequest) updates) => super.copyWith((message) => updates(message as UserRegisterRequest)) as UserRegisterRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserRegisterRequest create() => UserRegisterRequest._();
  UserRegisterRequest createEmptyInstance() => create();
  static $pb.PbList<UserRegisterRequest> createRepeated() => $pb.PbList<UserRegisterRequest>();
  @$core.pragma('dart2js:noInline')
  static UserRegisterRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserRegisterRequest>(create);
  static UserRegisterRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get firstName => $_getSZ(0);
  @$pb.TagNumber(1)
  set firstName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFirstName() => $_has(0);
  @$pb.TagNumber(1)
  void clearFirstName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get lastName => $_getSZ(1);
  @$pb.TagNumber(2)
  set lastName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLastName() => $_has(1);
  @$pb.TagNumber(2)
  void clearLastName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get firstNameKana => $_getSZ(2);
  @$pb.TagNumber(3)
  set firstNameKana($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasFirstNameKana() => $_has(2);
  @$pb.TagNumber(3)
  void clearFirstNameKana() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get lastNameKana => $_getSZ(3);
  @$pb.TagNumber(4)
  set lastNameKana($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasLastNameKana() => $_has(3);
  @$pb.TagNumber(4)
  void clearLastNameKana() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get companyName => $_getSZ(4);
  @$pb.TagNumber(5)
  set companyName($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasCompanyName() => $_has(4);
  @$pb.TagNumber(5)
  void clearCompanyName() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get birthDate => $_getSZ(5);
  @$pb.TagNumber(6)
  set birthDate($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasBirthDate() => $_has(5);
  @$pb.TagNumber(6)
  void clearBirthDate() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get zipCode => $_getSZ(6);
  @$pb.TagNumber(7)
  set zipCode($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasZipCode() => $_has(6);
  @$pb.TagNumber(7)
  void clearZipCode() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get prefecture => $_getSZ(7);
  @$pb.TagNumber(8)
  set prefecture($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasPrefecture() => $_has(7);
  @$pb.TagNumber(8)
  void clearPrefecture() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get city => $_getSZ(8);
  @$pb.TagNumber(9)
  set city($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasCity() => $_has(8);
  @$pb.TagNumber(9)
  void clearCity() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get address => $_getSZ(9);
  @$pb.TagNumber(10)
  set address($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasAddress() => $_has(9);
  @$pb.TagNumber(10)
  void clearAddress() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get tel => $_getSZ(10);
  @$pb.TagNumber(11)
  set tel($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasTel() => $_has(10);
  @$pb.TagNumber(11)
  void clearTel() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get mail => $_getSZ(11);
  @$pb.TagNumber(12)
  set mail($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasMail() => $_has(11);
  @$pb.TagNumber(12)
  void clearMail() => clearField(12);

  @$pb.TagNumber(13)
  $core.bool get acceptMail => $_getBF(12);
  @$pb.TagNumber(13)
  set acceptMail($core.bool v) { $_setBool(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasAcceptMail() => $_has(12);
  @$pb.TagNumber(13)
  void clearAcceptMail() => clearField(13);

  @$pb.TagNumber(14)
  $core.bool get acceptTerm => $_getBF(13);
  @$pb.TagNumber(14)
  set acceptTerm($core.bool v) { $_setBool(13, v); }
  @$pb.TagNumber(14)
  $core.bool hasAcceptTerm() => $_has(13);
  @$pb.TagNumber(14)
  void clearAcceptTerm() => clearField(14);
}

class UserAuthRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserAuthRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'email')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'password')
    ..hasRequiredFields = false
  ;

  UserAuthRequest._() : super();
  factory UserAuthRequest({
    $core.String? email,
    $core.String? password,
  }) {
    final _result = create();
    if (email != null) {
      _result.email = email;
    }
    if (password != null) {
      _result.password = password;
    }
    return _result;
  }
  factory UserAuthRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserAuthRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserAuthRequest clone() => UserAuthRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserAuthRequest copyWith(void Function(UserAuthRequest) updates) => super.copyWith((message) => updates(message as UserAuthRequest)) as UserAuthRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserAuthRequest create() => UserAuthRequest._();
  UserAuthRequest createEmptyInstance() => create();
  static $pb.PbList<UserAuthRequest> createRepeated() => $pb.PbList<UserAuthRequest>();
  @$core.pragma('dart2js:noInline')
  static UserAuthRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserAuthRequest>(create);
  static UserAuthRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get email => $_getSZ(0);
  @$pb.TagNumber(1)
  set email($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasEmail() => $_has(0);
  @$pb.TagNumber(1)
  void clearEmail() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get password => $_getSZ(1);
  @$pb.TagNumber(2)
  set password($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPassword() => $_has(1);
  @$pb.TagNumber(2)
  void clearPassword() => clearField(2);
}

class UserAuthResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserAuthResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'accessToken', protoName: 'accessToken')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'expiresIn', protoName: 'expiresIn')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'refreshToken', protoName: 'refreshToken')
    ..hasRequiredFields = false
  ;

  UserAuthResponse._() : super();
  factory UserAuthResponse({
    $core.String? accessToken,
    $fixnum.Int64? expiresIn,
    $core.String? refreshToken,
  }) {
    final _result = create();
    if (accessToken != null) {
      _result.accessToken = accessToken;
    }
    if (expiresIn != null) {
      _result.expiresIn = expiresIn;
    }
    if (refreshToken != null) {
      _result.refreshToken = refreshToken;
    }
    return _result;
  }
  factory UserAuthResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserAuthResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserAuthResponse clone() => UserAuthResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserAuthResponse copyWith(void Function(UserAuthResponse) updates) => super.copyWith((message) => updates(message as UserAuthResponse)) as UserAuthResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserAuthResponse create() => UserAuthResponse._();
  UserAuthResponse createEmptyInstance() => create();
  static $pb.PbList<UserAuthResponse> createRepeated() => $pb.PbList<UserAuthResponse>();
  @$core.pragma('dart2js:noInline')
  static UserAuthResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserAuthResponse>(create);
  static UserAuthResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get accessToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set accessToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccessToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccessToken() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get expiresIn => $_getI64(1);
  @$pb.TagNumber(2)
  set expiresIn($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasExpiresIn() => $_has(1);
  @$pb.TagNumber(2)
  void clearExpiresIn() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get refreshToken => $_getSZ(2);
  @$pb.TagNumber(3)
  set refreshToken($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasRefreshToken() => $_has(2);
  @$pb.TagNumber(3)
  void clearRefreshToken() => clearField(3);
}

class ResetPasswordRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ResetPasswordRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'email')
    ..hasRequiredFields = false
  ;

  ResetPasswordRequest._() : super();
  factory ResetPasswordRequest({
    $core.String? email,
  }) {
    final _result = create();
    if (email != null) {
      _result.email = email;
    }
    return _result;
  }
  factory ResetPasswordRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ResetPasswordRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ResetPasswordRequest clone() => ResetPasswordRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ResetPasswordRequest copyWith(void Function(ResetPasswordRequest) updates) => super.copyWith((message) => updates(message as ResetPasswordRequest)) as ResetPasswordRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ResetPasswordRequest create() => ResetPasswordRequest._();
  ResetPasswordRequest createEmptyInstance() => create();
  static $pb.PbList<ResetPasswordRequest> createRepeated() => $pb.PbList<ResetPasswordRequest>();
  @$core.pragma('dart2js:noInline')
  static ResetPasswordRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ResetPasswordRequest>(create);
  static ResetPasswordRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get email => $_getSZ(0);
  @$pb.TagNumber(1)
  set email($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasEmail() => $_has(0);
  @$pb.TagNumber(1)
  void clearEmail() => clearField(1);
}

class UpdatePasswordRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdatePasswordRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'password')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..hasRequiredFields = false
  ;

  UpdatePasswordRequest._() : super();
  factory UpdatePasswordRequest({
    $core.String? password,
    $core.String? token,
  }) {
    final _result = create();
    if (password != null) {
      _result.password = password;
    }
    if (token != null) {
      _result.token = token;
    }
    return _result;
  }
  factory UpdatePasswordRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdatePasswordRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdatePasswordRequest clone() => UpdatePasswordRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdatePasswordRequest copyWith(void Function(UpdatePasswordRequest) updates) => super.copyWith((message) => updates(message as UpdatePasswordRequest)) as UpdatePasswordRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdatePasswordRequest create() => UpdatePasswordRequest._();
  UpdatePasswordRequest createEmptyInstance() => create();
  static $pb.PbList<UpdatePasswordRequest> createRepeated() => $pb.PbList<UpdatePasswordRequest>();
  @$core.pragma('dart2js:noInline')
  static UpdatePasswordRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdatePasswordRequest>(create);
  static UpdatePasswordRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get password => $_getSZ(0);
  @$pb.TagNumber(1)
  set password($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasPassword() => $_has(0);
  @$pb.TagNumber(1)
  void clearPassword() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get token => $_getSZ(1);
  @$pb.TagNumber(2)
  set token($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearToken() => clearField(2);
}

class UpdateEmailRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateEmailRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'email')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..hasRequiredFields = false
  ;

  UpdateEmailRequest._() : super();
  factory UpdateEmailRequest({
    $core.String? email,
    $core.String? token,
  }) {
    final _result = create();
    if (email != null) {
      _result.email = email;
    }
    if (token != null) {
      _result.token = token;
    }
    return _result;
  }
  factory UpdateEmailRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateEmailRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateEmailRequest clone() => UpdateEmailRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateEmailRequest copyWith(void Function(UpdateEmailRequest) updates) => super.copyWith((message) => updates(message as UpdateEmailRequest)) as UpdateEmailRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateEmailRequest create() => UpdateEmailRequest._();
  UpdateEmailRequest createEmptyInstance() => create();
  static $pb.PbList<UpdateEmailRequest> createRepeated() => $pb.PbList<UpdateEmailRequest>();
  @$core.pragma('dart2js:noInline')
  static UpdateEmailRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateEmailRequest>(create);
  static UpdateEmailRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get email => $_getSZ(0);
  @$pb.TagNumber(1)
  set email($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasEmail() => $_has(0);
  @$pb.TagNumber(1)
  void clearEmail() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get token => $_getSZ(1);
  @$pb.TagNumber(2)
  set token($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearToken() => clearField(2);
}

class AuthControllerApi {
  $pb.RpcClient _client;
  AuthControllerApi(this._client);

  $async.Future<$1.Empty> register($pb.ClientContext? ctx, UserRegisterRequest request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'AuthController', 'Register', request, emptyResponse);
  }
  $async.Future<$1.Empty> signUp($pb.ClientContext? ctx, UserAuthRequest request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'AuthController', 'SignUp', request, emptyResponse);
  }
  $async.Future<UserAuthResponse> signIn($pb.ClientContext? ctx, UserAuthRequest request) {
    var emptyResponse = UserAuthResponse();
    return _client.invoke<UserAuthResponse>(ctx, 'AuthController', 'SignIn', request, emptyResponse);
  }
  $async.Future<$1.Empty> resetPasswordMail($pb.ClientContext? ctx, ResetPasswordRequest request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'AuthController', 'ResetPasswordMail', request, emptyResponse);
  }
  $async.Future<$1.Empty> updatePassword($pb.ClientContext? ctx, UpdatePasswordRequest request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'AuthController', 'UpdatePassword', request, emptyResponse);
  }
  $async.Future<$1.Empty> updateEmail($pb.ClientContext? ctx, UpdateEmailRequest request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'AuthController', 'UpdateEmail', request, emptyResponse);
  }
}


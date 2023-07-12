///
//  Generated code. Do not modify.
//  source: v1/user/UserData.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

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

class UserUpdateDataRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserUpdateDataRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FirstName', protoName: 'FirstName')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'LastName', protoName: 'LastName')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FirstNameKana', protoName: 'FirstNameKana')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'LastNameKana', protoName: 'LastNameKana')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'CompanyName', protoName: 'CompanyName')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'BirthDate', protoName: 'BirthDate')
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ZipCode', protoName: 'ZipCode')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Prefecture', protoName: 'Prefecture')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'City', protoName: 'City')
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Address', protoName: 'Address')
    ..aOS(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Tel', protoName: 'Tel')
    ..aOS(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Mail', protoName: 'Mail')
    ..aOB(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AcceptMail', protoName: 'AcceptMail')
    ..hasRequiredFields = false
  ;

  UserUpdateDataRequest._() : super();
  factory UserUpdateDataRequest({
    $core.String? iD,
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
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
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
    return _result;
  }
  factory UserUpdateDataRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserUpdateDataRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserUpdateDataRequest clone() => UserUpdateDataRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserUpdateDataRequest copyWith(void Function(UserUpdateDataRequest) updates) => super.copyWith((message) => updates(message as UserUpdateDataRequest)) as UserUpdateDataRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserUpdateDataRequest create() => UserUpdateDataRequest._();
  UserUpdateDataRequest createEmptyInstance() => create();
  static $pb.PbList<UserUpdateDataRequest> createRepeated() => $pb.PbList<UserUpdateDataRequest>();
  @$core.pragma('dart2js:noInline')
  static UserUpdateDataRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserUpdateDataRequest>(create);
  static UserUpdateDataRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get iD => $_getSZ(0);
  @$pb.TagNumber(1)
  set iD($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get firstName => $_getSZ(1);
  @$pb.TagNumber(2)
  set firstName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFirstName() => $_has(1);
  @$pb.TagNumber(2)
  void clearFirstName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get lastName => $_getSZ(2);
  @$pb.TagNumber(3)
  set lastName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLastName() => $_has(2);
  @$pb.TagNumber(3)
  void clearLastName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get firstNameKana => $_getSZ(3);
  @$pb.TagNumber(4)
  set firstNameKana($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasFirstNameKana() => $_has(3);
  @$pb.TagNumber(4)
  void clearFirstNameKana() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get lastNameKana => $_getSZ(4);
  @$pb.TagNumber(5)
  set lastNameKana($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasLastNameKana() => $_has(4);
  @$pb.TagNumber(5)
  void clearLastNameKana() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get companyName => $_getSZ(5);
  @$pb.TagNumber(6)
  set companyName($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasCompanyName() => $_has(5);
  @$pb.TagNumber(6)
  void clearCompanyName() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get birthDate => $_getSZ(6);
  @$pb.TagNumber(7)
  set birthDate($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasBirthDate() => $_has(6);
  @$pb.TagNumber(7)
  void clearBirthDate() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get zipCode => $_getSZ(7);
  @$pb.TagNumber(8)
  set zipCode($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasZipCode() => $_has(7);
  @$pb.TagNumber(8)
  void clearZipCode() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get prefecture => $_getSZ(8);
  @$pb.TagNumber(9)
  set prefecture($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasPrefecture() => $_has(8);
  @$pb.TagNumber(9)
  void clearPrefecture() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get city => $_getSZ(9);
  @$pb.TagNumber(10)
  set city($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasCity() => $_has(9);
  @$pb.TagNumber(10)
  void clearCity() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get address => $_getSZ(10);
  @$pb.TagNumber(11)
  set address($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasAddress() => $_has(10);
  @$pb.TagNumber(11)
  void clearAddress() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get tel => $_getSZ(11);
  @$pb.TagNumber(12)
  set tel($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasTel() => $_has(11);
  @$pb.TagNumber(12)
  void clearTel() => clearField(12);

  @$pb.TagNumber(13)
  $core.String get mail => $_getSZ(12);
  @$pb.TagNumber(13)
  set mail($core.String v) { $_setString(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasMail() => $_has(12);
  @$pb.TagNumber(13)
  void clearMail() => clearField(13);

  @$pb.TagNumber(14)
  $core.bool get acceptMail => $_getBF(13);
  @$pb.TagNumber(14)
  set acceptMail($core.bool v) { $_setBool(13, v); }
  @$pb.TagNumber(14)
  $core.bool hasAcceptMail() => $_has(13);
  @$pb.TagNumber(14)
  void clearAcceptMail() => clearField(14);
}

class UserDataResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserDataResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FirstName', protoName: 'FirstName')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'LastName', protoName: 'LastName')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FirstNameKana', protoName: 'FirstNameKana')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'LastNameKana', protoName: 'LastNameKana')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'CompanyName', protoName: 'CompanyName')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'BirthDate', protoName: 'BirthDate')
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ZipCode', protoName: 'ZipCode')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Prefecture', protoName: 'Prefecture')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'City', protoName: 'City')
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Address', protoName: 'Address')
    ..aOS(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Tel', protoName: 'Tel')
    ..aOS(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Mail', protoName: 'Mail')
    ..aOB(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AcceptMail', protoName: 'AcceptMail')
    ..hasRequiredFields = false
  ;

  UserDataResponse._() : super();
  factory UserDataResponse({
    $core.String? iD,
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
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
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
    return _result;
  }
  factory UserDataResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserDataResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserDataResponse clone() => UserDataResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserDataResponse copyWith(void Function(UserDataResponse) updates) => super.copyWith((message) => updates(message as UserDataResponse)) as UserDataResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserDataResponse create() => UserDataResponse._();
  UserDataResponse createEmptyInstance() => create();
  static $pb.PbList<UserDataResponse> createRepeated() => $pb.PbList<UserDataResponse>();
  @$core.pragma('dart2js:noInline')
  static UserDataResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserDataResponse>(create);
  static UserDataResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get iD => $_getSZ(0);
  @$pb.TagNumber(1)
  set iD($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get firstName => $_getSZ(1);
  @$pb.TagNumber(2)
  set firstName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFirstName() => $_has(1);
  @$pb.TagNumber(2)
  void clearFirstName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get lastName => $_getSZ(2);
  @$pb.TagNumber(3)
  set lastName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLastName() => $_has(2);
  @$pb.TagNumber(3)
  void clearLastName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get firstNameKana => $_getSZ(3);
  @$pb.TagNumber(4)
  set firstNameKana($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasFirstNameKana() => $_has(3);
  @$pb.TagNumber(4)
  void clearFirstNameKana() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get lastNameKana => $_getSZ(4);
  @$pb.TagNumber(5)
  set lastNameKana($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasLastNameKana() => $_has(4);
  @$pb.TagNumber(5)
  void clearLastNameKana() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get companyName => $_getSZ(5);
  @$pb.TagNumber(6)
  set companyName($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasCompanyName() => $_has(5);
  @$pb.TagNumber(6)
  void clearCompanyName() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get birthDate => $_getSZ(6);
  @$pb.TagNumber(7)
  set birthDate($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasBirthDate() => $_has(6);
  @$pb.TagNumber(7)
  void clearBirthDate() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get zipCode => $_getSZ(7);
  @$pb.TagNumber(8)
  set zipCode($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasZipCode() => $_has(7);
  @$pb.TagNumber(8)
  void clearZipCode() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get prefecture => $_getSZ(8);
  @$pb.TagNumber(9)
  set prefecture($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasPrefecture() => $_has(8);
  @$pb.TagNumber(9)
  void clearPrefecture() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get city => $_getSZ(9);
  @$pb.TagNumber(10)
  set city($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasCity() => $_has(9);
  @$pb.TagNumber(10)
  void clearCity() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get address => $_getSZ(10);
  @$pb.TagNumber(11)
  set address($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasAddress() => $_has(10);
  @$pb.TagNumber(11)
  void clearAddress() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get tel => $_getSZ(11);
  @$pb.TagNumber(12)
  set tel($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasTel() => $_has(11);
  @$pb.TagNumber(12)
  void clearTel() => clearField(12);

  @$pb.TagNumber(13)
  $core.String get mail => $_getSZ(12);
  @$pb.TagNumber(13)
  set mail($core.String v) { $_setString(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasMail() => $_has(12);
  @$pb.TagNumber(13)
  void clearMail() => clearField(13);

  @$pb.TagNumber(14)
  $core.bool get acceptMail => $_getBF(13);
  @$pb.TagNumber(14)
  set acceptMail($core.bool v) { $_setBool(13, v); }
  @$pb.TagNumber(14)
  $core.bool hasAcceptMail() => $_has(13);
  @$pb.TagNumber(14)
  void clearAcceptMail() => clearField(14);
}

class UserDataControllerApi {
  $pb.RpcClient _client;
  UserDataControllerApi(this._client);

  $async.Future<UserDataResponse> register($pb.ClientContext? ctx, UserRegisterRequest request) {
    var emptyResponse = UserDataResponse();
    return _client.invoke<UserDataResponse>(ctx, 'UserDataController', 'Register', request, emptyResponse);
  }
  $async.Future<UserDataResponse> update($pb.ClientContext? ctx, UserUpdateDataRequest request) {
    var emptyResponse = UserDataResponse();
    return _client.invoke<UserDataResponse>(ctx, 'UserDataController', 'Update', request, emptyResponse);
  }
}


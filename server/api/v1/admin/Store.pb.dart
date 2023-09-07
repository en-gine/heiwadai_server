///
//  Generated code. Do not modify.
//  source: v1/admin/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../shared/Store.pb.dart' as $1;

class StoreRegisterRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StoreRegisterRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.admin'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ZipCode', protoName: 'ZipCode')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Address', protoName: 'Address')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Tel', protoName: 'Tel')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Parking', protoName: 'Parking')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AccessInfo', protoName: 'AccessInfo')
    ..aOB(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsActive', protoName: 'IsActive')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'QRCode', protoName: 'QRCode')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'UnLimitedQRCode', protoName: 'UnLimitedQRCode')
    ..hasRequiredFields = false
  ;

  StoreRegisterRequest._() : super();
  factory StoreRegisterRequest({
    $core.String? name,
    $core.String? zipCode,
    $core.String? address,
    $core.String? tel,
    $core.String? parking,
    $core.String? accessInfo,
    $core.bool? isActive,
    $core.String? qRCode,
    $core.String? unLimitedQRCode,
  }) {
    final _result = create();
    if (name != null) {
      _result.name = name;
    }
    if (zipCode != null) {
      _result.zipCode = zipCode;
    }
    if (address != null) {
      _result.address = address;
    }
    if (tel != null) {
      _result.tel = tel;
    }
    if (parking != null) {
      _result.parking = parking;
    }
    if (accessInfo != null) {
      _result.accessInfo = accessInfo;
    }
    if (isActive != null) {
      _result.isActive = isActive;
    }
    if (qRCode != null) {
      _result.qRCode = qRCode;
    }
    if (unLimitedQRCode != null) {
      _result.unLimitedQRCode = unLimitedQRCode;
    }
    return _result;
  }
  factory StoreRegisterRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StoreRegisterRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StoreRegisterRequest clone() => StoreRegisterRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StoreRegisterRequest copyWith(void Function(StoreRegisterRequest) updates) => super.copyWith((message) => updates(message as StoreRegisterRequest)) as StoreRegisterRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StoreRegisterRequest create() => StoreRegisterRequest._();
  StoreRegisterRequest createEmptyInstance() => create();
  static $pb.PbList<StoreRegisterRequest> createRepeated() => $pb.PbList<StoreRegisterRequest>();
  @$core.pragma('dart2js:noInline')
  static StoreRegisterRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StoreRegisterRequest>(create);
  static StoreRegisterRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get zipCode => $_getSZ(1);
  @$pb.TagNumber(2)
  set zipCode($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasZipCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearZipCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get address => $_getSZ(2);
  @$pb.TagNumber(3)
  set address($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasAddress() => $_has(2);
  @$pb.TagNumber(3)
  void clearAddress() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get tel => $_getSZ(3);
  @$pb.TagNumber(4)
  set tel($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTel() => $_has(3);
  @$pb.TagNumber(4)
  void clearTel() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get parking => $_getSZ(4);
  @$pb.TagNumber(5)
  set parking($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasParking() => $_has(4);
  @$pb.TagNumber(5)
  void clearParking() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get accessInfo => $_getSZ(5);
  @$pb.TagNumber(6)
  set accessInfo($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasAccessInfo() => $_has(5);
  @$pb.TagNumber(6)
  void clearAccessInfo() => clearField(6);

  @$pb.TagNumber(7)
  $core.bool get isActive => $_getBF(6);
  @$pb.TagNumber(7)
  set isActive($core.bool v) { $_setBool(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasIsActive() => $_has(6);
  @$pb.TagNumber(7)
  void clearIsActive() => clearField(7);

  @$pb.TagNumber(9)
  $core.String get qRCode => $_getSZ(7);
  @$pb.TagNumber(9)
  set qRCode($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(9)
  $core.bool hasQRCode() => $_has(7);
  @$pb.TagNumber(9)
  void clearQRCode() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get unLimitedQRCode => $_getSZ(8);
  @$pb.TagNumber(10)
  set unLimitedQRCode($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(10)
  $core.bool hasUnLimitedQRCode() => $_has(8);
  @$pb.TagNumber(10)
  void clearUnLimitedQRCode() => clearField(10);
}

class StoreUpdateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StoreUpdateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.admin'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ZipCode', protoName: 'ZipCode')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Address', protoName: 'Address')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Tel', protoName: 'Tel')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Parking', protoName: 'Parking')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AccessInfo', protoName: 'AccessInfo')
    ..aOB(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsActive', protoName: 'IsActive')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'QRCode', protoName: 'QRCode')
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'UnLimitedQRCode', protoName: 'UnLimitedQRCode')
    ..hasRequiredFields = false
  ;

  StoreUpdateRequest._() : super();
  factory StoreUpdateRequest({
    $core.String? iD,
    $core.String? name,
    $core.String? zipCode,
    $core.String? address,
    $core.String? tel,
    $core.String? parking,
    $core.String? accessInfo,
    $core.bool? isActive,
    $core.String? qRCode,
    $core.String? unLimitedQRCode,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    if (name != null) {
      _result.name = name;
    }
    if (zipCode != null) {
      _result.zipCode = zipCode;
    }
    if (address != null) {
      _result.address = address;
    }
    if (tel != null) {
      _result.tel = tel;
    }
    if (parking != null) {
      _result.parking = parking;
    }
    if (accessInfo != null) {
      _result.accessInfo = accessInfo;
    }
    if (isActive != null) {
      _result.isActive = isActive;
    }
    if (qRCode != null) {
      _result.qRCode = qRCode;
    }
    if (unLimitedQRCode != null) {
      _result.unLimitedQRCode = unLimitedQRCode;
    }
    return _result;
  }
  factory StoreUpdateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StoreUpdateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StoreUpdateRequest clone() => StoreUpdateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StoreUpdateRequest copyWith(void Function(StoreUpdateRequest) updates) => super.copyWith((message) => updates(message as StoreUpdateRequest)) as StoreUpdateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StoreUpdateRequest create() => StoreUpdateRequest._();
  StoreUpdateRequest createEmptyInstance() => create();
  static $pb.PbList<StoreUpdateRequest> createRepeated() => $pb.PbList<StoreUpdateRequest>();
  @$core.pragma('dart2js:noInline')
  static StoreUpdateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StoreUpdateRequest>(create);
  static StoreUpdateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get iD => $_getSZ(0);
  @$pb.TagNumber(1)
  set iD($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get zipCode => $_getSZ(2);
  @$pb.TagNumber(3)
  set zipCode($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasZipCode() => $_has(2);
  @$pb.TagNumber(3)
  void clearZipCode() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get address => $_getSZ(3);
  @$pb.TagNumber(4)
  set address($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAddress() => $_has(3);
  @$pb.TagNumber(4)
  void clearAddress() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get tel => $_getSZ(4);
  @$pb.TagNumber(5)
  set tel($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasTel() => $_has(4);
  @$pb.TagNumber(5)
  void clearTel() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get parking => $_getSZ(5);
  @$pb.TagNumber(6)
  set parking($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasParking() => $_has(5);
  @$pb.TagNumber(6)
  void clearParking() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get accessInfo => $_getSZ(6);
  @$pb.TagNumber(7)
  set accessInfo($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasAccessInfo() => $_has(6);
  @$pb.TagNumber(7)
  void clearAccessInfo() => clearField(7);

  @$pb.TagNumber(8)
  $core.bool get isActive => $_getBF(7);
  @$pb.TagNumber(8)
  set isActive($core.bool v) { $_setBool(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasIsActive() => $_has(7);
  @$pb.TagNumber(8)
  void clearIsActive() => clearField(8);

  @$pb.TagNumber(10)
  $core.String get qRCode => $_getSZ(8);
  @$pb.TagNumber(10)
  set qRCode($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(10)
  $core.bool hasQRCode() => $_has(8);
  @$pb.TagNumber(10)
  void clearQRCode() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get unLimitedQRCode => $_getSZ(9);
  @$pb.TagNumber(11)
  set unLimitedQRCode($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(11)
  $core.bool hasUnLimitedQRCode() => $_has(9);
  @$pb.TagNumber(11)
  void clearUnLimitedQRCode() => clearField(11);
}

class StoreControllerApi {
  $pb.RpcClient _client;
  StoreControllerApi(this._client);

  $async.Future<$1.StoreResponse> register($pb.ClientContext? ctx, StoreRegisterRequest request) {
    var emptyResponse = $1.StoreResponse();
    return _client.invoke<$1.StoreResponse>(ctx, 'StoreController', 'Register', request, emptyResponse);
  }
  $async.Future<$1.StoreResponse> update($pb.ClientContext? ctx, StoreUpdateRequest request) {
    var emptyResponse = $1.StoreResponse();
    return _client.invoke<$1.StoreResponse>(ctx, 'StoreController', 'Update', request, emptyResponse);
  }
}


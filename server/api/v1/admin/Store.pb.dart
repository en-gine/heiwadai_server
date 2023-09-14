///
//  Generated code. Do not modify.
//  source: v1/admin/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/empty.pb.dart' as $1;

class StoreRegisterRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StoreRegisterRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.admin'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'BranchName', protoName: 'BranchName')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ZipCode', protoName: 'ZipCode')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Address', protoName: 'Address')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Tel', protoName: 'Tel')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SiteURL', protoName: 'SiteURL')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'StampImageURL', protoName: 'StampImageURL')
    ..aOB(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Stayable', protoName: 'Stayable')
    ..aOB(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsActive', protoName: 'IsActive')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'QRCode', protoName: 'QRCode')
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'UnLimitedQRCode', protoName: 'UnLimitedQRCode')
    ..aOM<StayableInfo>(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'StayableInfo', protoName: 'StayableInfo', subBuilder: StayableInfo.create)
    ..hasRequiredFields = false
  ;

  StoreRegisterRequest._() : super();
  factory StoreRegisterRequest({
    $core.String? name,
    $core.String? branchName,
    $core.String? zipCode,
    $core.String? address,
    $core.String? tel,
    $core.String? siteURL,
    $core.String? stampImageURL,
    $core.bool? stayable,
    $core.bool? isActive,
    $core.String? qRCode,
    $core.String? unLimitedQRCode,
    StayableInfo? stayableInfo,
  }) {
    final _result = create();
    if (name != null) {
      _result.name = name;
    }
    if (branchName != null) {
      _result.branchName = branchName;
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
    if (siteURL != null) {
      _result.siteURL = siteURL;
    }
    if (stampImageURL != null) {
      _result.stampImageURL = stampImageURL;
    }
    if (stayable != null) {
      _result.stayable = stayable;
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
    if (stayableInfo != null) {
      _result.stayableInfo = stayableInfo;
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
  $core.String get branchName => $_getSZ(1);
  @$pb.TagNumber(2)
  set branchName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasBranchName() => $_has(1);
  @$pb.TagNumber(2)
  void clearBranchName() => clearField(2);

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
  $core.String get siteURL => $_getSZ(5);
  @$pb.TagNumber(6)
  set siteURL($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasSiteURL() => $_has(5);
  @$pb.TagNumber(6)
  void clearSiteURL() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get stampImageURL => $_getSZ(6);
  @$pb.TagNumber(7)
  set stampImageURL($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasStampImageURL() => $_has(6);
  @$pb.TagNumber(7)
  void clearStampImageURL() => clearField(7);

  @$pb.TagNumber(8)
  $core.bool get stayable => $_getBF(7);
  @$pb.TagNumber(8)
  set stayable($core.bool v) { $_setBool(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasStayable() => $_has(7);
  @$pb.TagNumber(8)
  void clearStayable() => clearField(8);

  @$pb.TagNumber(9)
  $core.bool get isActive => $_getBF(8);
  @$pb.TagNumber(9)
  set isActive($core.bool v) { $_setBool(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasIsActive() => $_has(8);
  @$pb.TagNumber(9)
  void clearIsActive() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get qRCode => $_getSZ(9);
  @$pb.TagNumber(10)
  set qRCode($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasQRCode() => $_has(9);
  @$pb.TagNumber(10)
  void clearQRCode() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get unLimitedQRCode => $_getSZ(10);
  @$pb.TagNumber(11)
  set unLimitedQRCode($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasUnLimitedQRCode() => $_has(10);
  @$pb.TagNumber(11)
  void clearUnLimitedQRCode() => clearField(11);

  @$pb.TagNumber(12)
  StayableInfo get stayableInfo => $_getN(11);
  @$pb.TagNumber(12)
  set stayableInfo(StayableInfo v) { setField(12, v); }
  @$pb.TagNumber(12)
  $core.bool hasStayableInfo() => $_has(11);
  @$pb.TagNumber(12)
  void clearStayableInfo() => clearField(12);
  @$pb.TagNumber(12)
  StayableInfo ensureStayableInfo() => $_ensure(11);
}

class StoreUpdateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StoreUpdateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.admin'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'BranchName', protoName: 'BranchName')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ZipCode', protoName: 'ZipCode')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Address', protoName: 'Address')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Tel', protoName: 'Tel')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SiteURL', protoName: 'SiteURL')
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'StampImageURL', protoName: 'StampImageURL')
    ..aOB(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Stayable', protoName: 'Stayable')
    ..aOB(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsActive', protoName: 'IsActive')
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'QRCode', protoName: 'QRCode')
    ..aOS(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'UnLimitedQRCode', protoName: 'UnLimitedQRCode')
    ..aOM<StayableInfo>(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'StayableInfo', protoName: 'StayableInfo', subBuilder: StayableInfo.create)
    ..hasRequiredFields = false
  ;

  StoreUpdateRequest._() : super();
  factory StoreUpdateRequest({
    $core.String? iD,
    $core.String? name,
    $core.String? branchName,
    $core.String? zipCode,
    $core.String? address,
    $core.String? tel,
    $core.String? siteURL,
    $core.String? stampImageURL,
    $core.bool? stayable,
    $core.bool? isActive,
    $core.String? qRCode,
    $core.String? unLimitedQRCode,
    StayableInfo? stayableInfo,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    if (name != null) {
      _result.name = name;
    }
    if (branchName != null) {
      _result.branchName = branchName;
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
    if (siteURL != null) {
      _result.siteURL = siteURL;
    }
    if (stampImageURL != null) {
      _result.stampImageURL = stampImageURL;
    }
    if (stayable != null) {
      _result.stayable = stayable;
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
    if (stayableInfo != null) {
      _result.stayableInfo = stayableInfo;
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
  $core.String get branchName => $_getSZ(2);
  @$pb.TagNumber(3)
  set branchName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasBranchName() => $_has(2);
  @$pb.TagNumber(3)
  void clearBranchName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get zipCode => $_getSZ(3);
  @$pb.TagNumber(4)
  set zipCode($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasZipCode() => $_has(3);
  @$pb.TagNumber(4)
  void clearZipCode() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get address => $_getSZ(4);
  @$pb.TagNumber(5)
  set address($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasAddress() => $_has(4);
  @$pb.TagNumber(5)
  void clearAddress() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get tel => $_getSZ(5);
  @$pb.TagNumber(6)
  set tel($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasTel() => $_has(5);
  @$pb.TagNumber(6)
  void clearTel() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get siteURL => $_getSZ(6);
  @$pb.TagNumber(7)
  set siteURL($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasSiteURL() => $_has(6);
  @$pb.TagNumber(7)
  void clearSiteURL() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get stampImageURL => $_getSZ(7);
  @$pb.TagNumber(8)
  set stampImageURL($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasStampImageURL() => $_has(7);
  @$pb.TagNumber(8)
  void clearStampImageURL() => clearField(8);

  @$pb.TagNumber(9)
  $core.bool get stayable => $_getBF(8);
  @$pb.TagNumber(9)
  set stayable($core.bool v) { $_setBool(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasStayable() => $_has(8);
  @$pb.TagNumber(9)
  void clearStayable() => clearField(9);

  @$pb.TagNumber(10)
  $core.bool get isActive => $_getBF(9);
  @$pb.TagNumber(10)
  set isActive($core.bool v) { $_setBool(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasIsActive() => $_has(9);
  @$pb.TagNumber(10)
  void clearIsActive() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get qRCode => $_getSZ(10);
  @$pb.TagNumber(11)
  set qRCode($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasQRCode() => $_has(10);
  @$pb.TagNumber(11)
  void clearQRCode() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get unLimitedQRCode => $_getSZ(11);
  @$pb.TagNumber(12)
  set unLimitedQRCode($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasUnLimitedQRCode() => $_has(11);
  @$pb.TagNumber(12)
  void clearUnLimitedQRCode() => clearField(12);

  @$pb.TagNumber(13)
  StayableInfo get stayableInfo => $_getN(12);
  @$pb.TagNumber(13)
  set stayableInfo(StayableInfo v) { setField(13, v); }
  @$pb.TagNumber(13)
  $core.bool hasStayableInfo() => $_has(12);
  @$pb.TagNumber(13)
  void clearStayableInfo() => clearField(13);
  @$pb.TagNumber(13)
  StayableInfo ensureStayableInfo() => $_ensure(12);
}

class StayableInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StayableInfo', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.admin'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Parking', protoName: 'Parking')
    ..a<$core.double>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Latitude', $pb.PbFieldType.OD, protoName: 'Latitude')
    ..a<$core.double>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Longitude', $pb.PbFieldType.OD, protoName: 'Longitude')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AccessInfo', protoName: 'AccessInfo')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'RestAPIURL', protoName: 'RestAPIURL')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'BookingSystemID', protoName: 'BookingSystemID')
    ..hasRequiredFields = false
  ;

  StayableInfo._() : super();
  factory StayableInfo({
    $core.String? parking,
    $core.double? latitude,
    $core.double? longitude,
    $core.String? accessInfo,
    $core.String? restAPIURL,
    $core.String? bookingSystemID,
  }) {
    final _result = create();
    if (parking != null) {
      _result.parking = parking;
    }
    if (latitude != null) {
      _result.latitude = latitude;
    }
    if (longitude != null) {
      _result.longitude = longitude;
    }
    if (accessInfo != null) {
      _result.accessInfo = accessInfo;
    }
    if (restAPIURL != null) {
      _result.restAPIURL = restAPIURL;
    }
    if (bookingSystemID != null) {
      _result.bookingSystemID = bookingSystemID;
    }
    return _result;
  }
  factory StayableInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StayableInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StayableInfo clone() => StayableInfo()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StayableInfo copyWith(void Function(StayableInfo) updates) => super.copyWith((message) => updates(message as StayableInfo)) as StayableInfo; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StayableInfo create() => StayableInfo._();
  StayableInfo createEmptyInstance() => create();
  static $pb.PbList<StayableInfo> createRepeated() => $pb.PbList<StayableInfo>();
  @$core.pragma('dart2js:noInline')
  static StayableInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StayableInfo>(create);
  static StayableInfo? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get parking => $_getSZ(0);
  @$pb.TagNumber(1)
  set parking($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasParking() => $_has(0);
  @$pb.TagNumber(1)
  void clearParking() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get latitude => $_getN(1);
  @$pb.TagNumber(2)
  set latitude($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLatitude() => $_has(1);
  @$pb.TagNumber(2)
  void clearLatitude() => clearField(2);

  @$pb.TagNumber(3)
  $core.double get longitude => $_getN(2);
  @$pb.TagNumber(3)
  set longitude($core.double v) { $_setDouble(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLongitude() => $_has(2);
  @$pb.TagNumber(3)
  void clearLongitude() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get accessInfo => $_getSZ(3);
  @$pb.TagNumber(4)
  set accessInfo($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAccessInfo() => $_has(3);
  @$pb.TagNumber(4)
  void clearAccessInfo() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get restAPIURL => $_getSZ(4);
  @$pb.TagNumber(5)
  set restAPIURL($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasRestAPIURL() => $_has(4);
  @$pb.TagNumber(5)
  void clearRestAPIURL() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get bookingSystemID => $_getSZ(5);
  @$pb.TagNumber(6)
  set bookingSystemID($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasBookingSystemID() => $_has(5);
  @$pb.TagNumber(6)
  void clearBookingSystemID() => clearField(6);
}

class StoreControllerApi {
  $pb.RpcClient _client;
  StoreControllerApi(this._client);

  $async.Future<$1.Empty> register($pb.ClientContext? ctx, StoreRegisterRequest request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'StoreController', 'Register', request, emptyResponse);
  }
  $async.Future<$1.Empty> update($pb.ClientContext? ctx, StoreUpdateRequest request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'StoreController', 'Update', request, emptyResponse);
  }
}


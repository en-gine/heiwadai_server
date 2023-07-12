///
//  Generated code. Do not modify.
//  source: v1/shared/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class StoreResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StoreResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.shared'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ZipCode', protoName: 'ZipCode')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Address', protoName: 'Address')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Tel', protoName: 'Tel')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Parking', protoName: 'Parking')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AccessInfo', protoName: 'AccessInfo')
    ..aOB(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsActive', protoName: 'IsActive')
    ..aOB(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Stayable', protoName: 'Stayable')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'QRCode', protoName: 'QRCode')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'UnLimitedQRCode', protoName: 'UnLimitedQRCode')
    ..hasRequiredFields = false
  ;

  StoreResponse._() : super();
  factory StoreResponse({
    $core.String? name,
    $core.String? zipCode,
    $core.String? address,
    $core.String? tel,
    $core.String? parking,
    $core.String? accessInfo,
    $core.bool? isActive,
    $core.bool? stayable,
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
    if (stayable != null) {
      _result.stayable = stayable;
    }
    if (qRCode != null) {
      _result.qRCode = qRCode;
    }
    if (unLimitedQRCode != null) {
      _result.unLimitedQRCode = unLimitedQRCode;
    }
    return _result;
  }
  factory StoreResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StoreResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StoreResponse clone() => StoreResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StoreResponse copyWith(void Function(StoreResponse) updates) => super.copyWith((message) => updates(message as StoreResponse)) as StoreResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StoreResponse create() => StoreResponse._();
  StoreResponse createEmptyInstance() => create();
  static $pb.PbList<StoreResponse> createRepeated() => $pb.PbList<StoreResponse>();
  @$core.pragma('dart2js:noInline')
  static StoreResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StoreResponse>(create);
  static StoreResponse? _defaultInstance;

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

  @$pb.TagNumber(8)
  $core.bool get stayable => $_getBF(7);
  @$pb.TagNumber(8)
  set stayable($core.bool v) { $_setBool(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasStayable() => $_has(7);
  @$pb.TagNumber(8)
  void clearStayable() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get qRCode => $_getSZ(8);
  @$pb.TagNumber(9)
  set qRCode($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasQRCode() => $_has(8);
  @$pb.TagNumber(9)
  void clearQRCode() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get unLimitedQRCode => $_getSZ(9);
  @$pb.TagNumber(10)
  set unLimitedQRCode($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasUnLimitedQRCode() => $_has(9);
  @$pb.TagNumber(10)
  void clearUnLimitedQRCode() => clearField(10);
}


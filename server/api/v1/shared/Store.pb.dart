///
//  Generated code. Do not modify.
//  source: v1/shared/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Store extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Store', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.shared'), createEmptyInstance: create)
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
    ..hasRequiredFields = false
  ;

  Store._() : super();
  factory Store({
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
    return _result;
  }
  factory Store.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Store.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Store clone() => Store()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Store copyWith(void Function(Store) updates) => super.copyWith((message) => updates(message as Store)) as Store; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Store create() => Store._();
  Store createEmptyInstance() => create();
  static $pb.PbList<Store> createRepeated() => $pb.PbList<Store>();
  @$core.pragma('dart2js:noInline')
  static Store getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Store>(create);
  static Store? _defaultInstance;

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
}

class Stores extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Stores', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.shared'), createEmptyInstance: create)
    ..pc<Store>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Stores', $pb.PbFieldType.PM, protoName: 'Stores', subBuilder: Store.create)
    ..hasRequiredFields = false
  ;

  Stores._() : super();
  factory Stores({
    $core.Iterable<Store>? stores,
  }) {
    final _result = create();
    if (stores != null) {
      _result.stores.addAll(stores);
    }
    return _result;
  }
  factory Stores.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Stores.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Stores clone() => Stores()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Stores copyWith(void Function(Stores) updates) => super.copyWith((message) => updates(message as Stores)) as Stores; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Stores create() => Stores._();
  Stores createEmptyInstance() => create();
  static $pb.PbList<Stores> createRepeated() => $pb.PbList<Stores>();
  @$core.pragma('dart2js:noInline')
  static Stores getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Stores>(create);
  static Stores? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Store> get stores => $_getList(0);
}

class StayableStoreInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StayableStoreInfo', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.shared'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Parking', protoName: 'Parking')
    ..a<$core.double>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Latitude', $pb.PbFieldType.OD, protoName: 'Latitude')
    ..a<$core.double>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Longitude', $pb.PbFieldType.OD, protoName: 'Longitude')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AccessInfo', protoName: 'AccessInfo')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'RestAPIURL', protoName: 'RestAPIURL')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'BookingSystemID', protoName: 'BookingSystemID')
    ..hasRequiredFields = false
  ;

  StayableStoreInfo._() : super();
  factory StayableStoreInfo({
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
  factory StayableStoreInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StayableStoreInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StayableStoreInfo clone() => StayableStoreInfo()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StayableStoreInfo copyWith(void Function(StayableStoreInfo) updates) => super.copyWith((message) => updates(message as StayableStoreInfo)) as StayableStoreInfo; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StayableStoreInfo create() => StayableStoreInfo._();
  StayableStoreInfo createEmptyInstance() => create();
  static $pb.PbList<StayableStoreInfo> createRepeated() => $pb.PbList<StayableStoreInfo>();
  @$core.pragma('dart2js:noInline')
  static StayableStoreInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StayableStoreInfo>(create);
  static StayableStoreInfo? _defaultInstance;

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

class StayableStore extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StayableStore', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.shared'), createEmptyInstance: create)
    ..aOM<Store>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Store', protoName: 'Store', subBuilder: Store.create)
    ..aOM<StayableStoreInfo>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Info', protoName: 'Info', subBuilder: StayableStoreInfo.create)
    ..hasRequiredFields = false
  ;

  StayableStore._() : super();
  factory StayableStore({
    Store? store,
    StayableStoreInfo? info,
  }) {
    final _result = create();
    if (store != null) {
      _result.store = store;
    }
    if (info != null) {
      _result.info = info;
    }
    return _result;
  }
  factory StayableStore.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StayableStore.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StayableStore clone() => StayableStore()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StayableStore copyWith(void Function(StayableStore) updates) => super.copyWith((message) => updates(message as StayableStore)) as StayableStore; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StayableStore create() => StayableStore._();
  StayableStore createEmptyInstance() => create();
  static $pb.PbList<StayableStore> createRepeated() => $pb.PbList<StayableStore>();
  @$core.pragma('dart2js:noInline')
  static StayableStore getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StayableStore>(create);
  static StayableStore? _defaultInstance;

  @$pb.TagNumber(1)
  Store get store => $_getN(0);
  @$pb.TagNumber(1)
  set store(Store v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStore() => $_has(0);
  @$pb.TagNumber(1)
  void clearStore() => clearField(1);
  @$pb.TagNumber(1)
  Store ensureStore() => $_ensure(0);

  @$pb.TagNumber(2)
  StayableStoreInfo get info => $_getN(1);
  @$pb.TagNumber(2)
  set info(StayableStoreInfo v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasInfo() => $_has(1);
  @$pb.TagNumber(2)
  void clearInfo() => clearField(2);
  @$pb.TagNumber(2)
  StayableStoreInfo ensureInfo() => $_ensure(1);
}

class StayableStores extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StayableStores', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.shared'), createEmptyInstance: create)
    ..pc<StayableStore>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'StayableStores', $pb.PbFieldType.PM, protoName: 'StayableStores', subBuilder: StayableStore.create)
    ..hasRequiredFields = false
  ;

  StayableStores._() : super();
  factory StayableStores({
    $core.Iterable<StayableStore>? stayableStores,
  }) {
    final _result = create();
    if (stayableStores != null) {
      _result.stayableStores.addAll(stayableStores);
    }
    return _result;
  }
  factory StayableStores.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StayableStores.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StayableStores clone() => StayableStores()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StayableStores copyWith(void Function(StayableStores) updates) => super.copyWith((message) => updates(message as StayableStores)) as StayableStores; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StayableStores create() => StayableStores._();
  StayableStores createEmptyInstance() => create();
  static $pb.PbList<StayableStores> createRepeated() => $pb.PbList<StayableStores>();
  @$core.pragma('dart2js:noInline')
  static StayableStores getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StayableStores>(create);
  static StayableStores? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<StayableStore> get stayableStores => $_getList(0);
}


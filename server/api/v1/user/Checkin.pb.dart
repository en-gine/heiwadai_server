///
//  Generated code. Do not modify.
//  source: v1/user/Checkin.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/timestamp.pb.dart' as $5;
import '../shared/Coupon.pb.dart' as $8;
import '../../google/protobuf/empty.pb.dart' as $2;

class CheckinRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CheckinRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'qrHash', protoName: 'qrHash')
    ..hasRequiredFields = false
  ;

  CheckinRequest._() : super();
  factory CheckinRequest({
    $core.String? qrHash,
  }) {
    final _result = create();
    if (qrHash != null) {
      _result.qrHash = qrHash;
    }
    return _result;
  }
  factory CheckinRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CheckinRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CheckinRequest clone() => CheckinRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CheckinRequest copyWith(void Function(CheckinRequest) updates) => super.copyWith((message) => updates(message as CheckinRequest)) as CheckinRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CheckinRequest create() => CheckinRequest._();
  CheckinRequest createEmptyInstance() => create();
  static $pb.PbList<CheckinRequest> createRepeated() => $pb.PbList<CheckinRequest>();
  @$core.pragma('dart2js:noInline')
  static CheckinRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CheckinRequest>(create);
  static CheckinRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get qrHash => $_getSZ(0);
  @$pb.TagNumber(1)
  set qrHash($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasQrHash() => $_has(0);
  @$pb.TagNumber(1)
  void clearQrHash() => clearField(1);
}

class CheckinStamp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CheckinStamp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'StoreID', protoName: 'StoreID')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'StoreName', protoName: 'StoreName')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'StoreStampImage', protoName: 'StoreStampImage')
    ..aOM<$5.Timestamp>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'CheckInAt', protoName: 'CheckInAt', subBuilder: $5.Timestamp.create)
    ..hasRequiredFields = false
  ;

  CheckinStamp._() : super();
  factory CheckinStamp({
    $core.String? iD,
    $core.String? storeID,
    $core.String? storeName,
    $core.String? storeStampImage,
    $5.Timestamp? checkInAt,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    if (storeID != null) {
      _result.storeID = storeID;
    }
    if (storeName != null) {
      _result.storeName = storeName;
    }
    if (storeStampImage != null) {
      _result.storeStampImage = storeStampImage;
    }
    if (checkInAt != null) {
      _result.checkInAt = checkInAt;
    }
    return _result;
  }
  factory CheckinStamp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CheckinStamp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CheckinStamp clone() => CheckinStamp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CheckinStamp copyWith(void Function(CheckinStamp) updates) => super.copyWith((message) => updates(message as CheckinStamp)) as CheckinStamp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CheckinStamp create() => CheckinStamp._();
  CheckinStamp createEmptyInstance() => create();
  static $pb.PbList<CheckinStamp> createRepeated() => $pb.PbList<CheckinStamp>();
  @$core.pragma('dart2js:noInline')
  static CheckinStamp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CheckinStamp>(create);
  static CheckinStamp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get iD => $_getSZ(0);
  @$pb.TagNumber(1)
  set iD($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get storeID => $_getSZ(1);
  @$pb.TagNumber(2)
  set storeID($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasStoreID() => $_has(1);
  @$pb.TagNumber(2)
  void clearStoreID() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get storeName => $_getSZ(2);
  @$pb.TagNumber(3)
  set storeName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasStoreName() => $_has(2);
  @$pb.TagNumber(3)
  void clearStoreName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get storeStampImage => $_getSZ(3);
  @$pb.TagNumber(4)
  set storeStampImage($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasStoreStampImage() => $_has(3);
  @$pb.TagNumber(4)
  void clearStoreStampImage() => clearField(4);

  @$pb.TagNumber(5)
  $5.Timestamp get checkInAt => $_getN(4);
  @$pb.TagNumber(5)
  set checkInAt($5.Timestamp v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasCheckInAt() => $_has(4);
  @$pb.TagNumber(5)
  void clearCheckInAt() => clearField(5);
  @$pb.TagNumber(5)
  $5.Timestamp ensureCheckInAt() => $_ensure(4);
}

class StampCardResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StampCardResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..pc<CheckinStamp>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'stamps', $pb.PbFieldType.PM, subBuilder: CheckinStamp.create)
    ..hasRequiredFields = false
  ;

  StampCardResponse._() : super();
  factory StampCardResponse({
    $core.Iterable<CheckinStamp>? stamps,
  }) {
    final _result = create();
    if (stamps != null) {
      _result.stamps.addAll(stamps);
    }
    return _result;
  }
  factory StampCardResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StampCardResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StampCardResponse clone() => StampCardResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StampCardResponse copyWith(void Function(StampCardResponse) updates) => super.copyWith((message) => updates(message as StampCardResponse)) as StampCardResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StampCardResponse create() => StampCardResponse._();
  StampCardResponse createEmptyInstance() => create();
  static $pb.PbList<StampCardResponse> createRepeated() => $pb.PbList<StampCardResponse>();
  @$core.pragma('dart2js:noInline')
  static StampCardResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StampCardResponse>(create);
  static StampCardResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<CheckinStamp> get stamps => $_getList(0);
}

class CheckinResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CheckinResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOM<$8.Coupon>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'MayCoupon', protoName: 'MayCoupon', subBuilder: $8.Coupon.create)
    ..hasRequiredFields = false
  ;

  CheckinResponse._() : super();
  factory CheckinResponse({
    $8.Coupon? mayCoupon,
  }) {
    final _result = create();
    if (mayCoupon != null) {
      _result.mayCoupon = mayCoupon;
    }
    return _result;
  }
  factory CheckinResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CheckinResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CheckinResponse clone() => CheckinResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CheckinResponse copyWith(void Function(CheckinResponse) updates) => super.copyWith((message) => updates(message as CheckinResponse)) as CheckinResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CheckinResponse create() => CheckinResponse._();
  CheckinResponse createEmptyInstance() => create();
  static $pb.PbList<CheckinResponse> createRepeated() => $pb.PbList<CheckinResponse>();
  @$core.pragma('dart2js:noInline')
  static CheckinResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CheckinResponse>(create);
  static CheckinResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $8.Coupon get mayCoupon => $_getN(0);
  @$pb.TagNumber(1)
  set mayCoupon($8.Coupon v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasMayCoupon() => $_has(0);
  @$pb.TagNumber(1)
  void clearMayCoupon() => clearField(1);
  @$pb.TagNumber(1)
  $8.Coupon ensureMayCoupon() => $_ensure(0);
}

class CheckinControllerApi {
  $pb.RpcClient _client;
  CheckinControllerApi(this._client);

  $async.Future<StampCardResponse> getStampCard($pb.ClientContext? ctx, $2.Empty request) {
    var emptyResponse = StampCardResponse();
    return _client.invoke<StampCardResponse>(ctx, 'CheckinController', 'GetStampCard', request, emptyResponse);
  }
  $async.Future<CheckinResponse> checkin($pb.ClientContext? ctx, CheckinRequest request) {
    var emptyResponse = CheckinResponse();
    return _client.invoke<CheckinResponse>(ctx, 'CheckinController', 'Checkin', request, emptyResponse);
  }
}


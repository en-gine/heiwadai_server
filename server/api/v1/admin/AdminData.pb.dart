///
//  Generated code. Do not modify.
//  source: v1/admin/AdminData.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class AdminUpdateDataRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AdminUpdateDataRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsActive', protoName: 'IsActive')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Mail', protoName: 'Mail')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'StoreID', protoName: 'StoreID')
    ..hasRequiredFields = false
  ;

  AdminUpdateDataRequest._() : super();
  factory AdminUpdateDataRequest({
    $core.String? iD,
    $core.String? name,
    $core.String? isActive,
    $core.String? mail,
    $core.String? storeID,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    if (name != null) {
      _result.name = name;
    }
    if (isActive != null) {
      _result.isActive = isActive;
    }
    if (mail != null) {
      _result.mail = mail;
    }
    if (storeID != null) {
      _result.storeID = storeID;
    }
    return _result;
  }
  factory AdminUpdateDataRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdminUpdateDataRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AdminUpdateDataRequest clone() => AdminUpdateDataRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AdminUpdateDataRequest copyWith(void Function(AdminUpdateDataRequest) updates) => super.copyWith((message) => updates(message as AdminUpdateDataRequest)) as AdminUpdateDataRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AdminUpdateDataRequest create() => AdminUpdateDataRequest._();
  AdminUpdateDataRequest createEmptyInstance() => create();
  static $pb.PbList<AdminUpdateDataRequest> createRepeated() => $pb.PbList<AdminUpdateDataRequest>();
  @$core.pragma('dart2js:noInline')
  static AdminUpdateDataRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdminUpdateDataRequest>(create);
  static AdminUpdateDataRequest? _defaultInstance;

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
  $core.String get isActive => $_getSZ(2);
  @$pb.TagNumber(3)
  set isActive($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasIsActive() => $_has(2);
  @$pb.TagNumber(3)
  void clearIsActive() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get mail => $_getSZ(3);
  @$pb.TagNumber(4)
  set mail($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasMail() => $_has(3);
  @$pb.TagNumber(4)
  void clearMail() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get storeID => $_getSZ(4);
  @$pb.TagNumber(5)
  set storeID($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasStoreID() => $_has(4);
  @$pb.TagNumber(5)
  void clearStoreID() => clearField(5);
}

class AdminDataResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AdminDataResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Name', protoName: 'Name')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsActive', protoName: 'IsActive')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Mail', protoName: 'Mail')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'StoreID', protoName: 'StoreID')
    ..hasRequiredFields = false
  ;

  AdminDataResponse._() : super();
  factory AdminDataResponse({
    $core.String? iD,
    $core.String? name,
    $core.String? isActive,
    $core.String? mail,
    $core.String? storeID,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    if (name != null) {
      _result.name = name;
    }
    if (isActive != null) {
      _result.isActive = isActive;
    }
    if (mail != null) {
      _result.mail = mail;
    }
    if (storeID != null) {
      _result.storeID = storeID;
    }
    return _result;
  }
  factory AdminDataResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdminDataResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AdminDataResponse clone() => AdminDataResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AdminDataResponse copyWith(void Function(AdminDataResponse) updates) => super.copyWith((message) => updates(message as AdminDataResponse)) as AdminDataResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AdminDataResponse create() => AdminDataResponse._();
  AdminDataResponse createEmptyInstance() => create();
  static $pb.PbList<AdminDataResponse> createRepeated() => $pb.PbList<AdminDataResponse>();
  @$core.pragma('dart2js:noInline')
  static AdminDataResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdminDataResponse>(create);
  static AdminDataResponse? _defaultInstance;

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
  $core.String get isActive => $_getSZ(2);
  @$pb.TagNumber(3)
  set isActive($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasIsActive() => $_has(2);
  @$pb.TagNumber(3)
  void clearIsActive() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get mail => $_getSZ(3);
  @$pb.TagNumber(4)
  set mail($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasMail() => $_has(3);
  @$pb.TagNumber(4)
  void clearMail() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get storeID => $_getSZ(4);
  @$pb.TagNumber(5)
  set storeID($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasStoreID() => $_has(4);
  @$pb.TagNumber(5)
  void clearStoreID() => clearField(5);
}

class AdminDataControllerApi {
  $pb.RpcClient _client;
  AdminDataControllerApi(this._client);

  $async.Future<AdminDataResponse> update($pb.ClientContext? ctx, AdminUpdateDataRequest request) {
    var emptyResponse = AdminDataResponse();
    return _client.invoke<AdminDataResponse>(ctx, 'AdminDataController', 'Update', request, emptyResponse);
  }
}


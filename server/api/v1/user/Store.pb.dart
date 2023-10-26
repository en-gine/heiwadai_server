///
//  Generated code. Do not modify.
//  source: v1/user/Store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../shared/Store.pb.dart' as $3;
import '../../google/protobuf/empty.pb.dart' as $1;

class SoreIDRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SoreIDRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..hasRequiredFields = false
  ;

  SoreIDRequest._() : super();
  factory SoreIDRequest({
    $core.String? iD,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    return _result;
  }
  factory SoreIDRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SoreIDRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SoreIDRequest clone() => SoreIDRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SoreIDRequest copyWith(void Function(SoreIDRequest) updates) => super.copyWith((message) => updates(message as SoreIDRequest)) as SoreIDRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SoreIDRequest create() => SoreIDRequest._();
  SoreIDRequest createEmptyInstance() => create();
  static $pb.PbList<SoreIDRequest> createRepeated() => $pb.PbList<SoreIDRequest>();
  @$core.pragma('dart2js:noInline')
  static SoreIDRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SoreIDRequest>(create);
  static SoreIDRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get iD => $_getSZ(0);
  @$pb.TagNumber(1)
  set iD($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);
}

class StoreControllerApi {
  $pb.RpcClient _client;
  StoreControllerApi(this._client);

  $async.Future<$3.Store> getByID($pb.ClientContext? ctx, SoreIDRequest request) {
    var emptyResponse = $3.Store();
    return _client.invoke<$3.Store>(ctx, 'StoreController', 'GetByID', request, emptyResponse);
  }
  $async.Future<$3.Stores> getAll($pb.ClientContext? ctx, $1.Empty request) {
    var emptyResponse = $3.Stores();
    return _client.invoke<$3.Stores>(ctx, 'StoreController', 'GetAll', request, emptyResponse);
  }
  $async.Future<$3.StayableStores> getStayables($pb.ClientContext? ctx, $1.Empty request) {
    var emptyResponse = $3.StayableStores();
    return _client.invoke<$3.StayableStores>(ctx, 'StoreController', 'GetStayables', request, emptyResponse);
  }
  $async.Future<$3.StayableStore> getStayableByID($pb.ClientContext? ctx, SoreIDRequest request) {
    var emptyResponse = $3.StayableStore();
    return _client.invoke<$3.StayableStore>(ctx, 'StoreController', 'GetStayableByID', request, emptyResponse);
  }
}


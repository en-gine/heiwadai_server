///
//  Generated code. Do not modify.
//  source: v1/user/Banner.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/empty.pb.dart' as $1;

class Banner extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Banner', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ImageURL', protoName: 'ImageURL')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'URL', protoName: 'URL')
    ..hasRequiredFields = false
  ;

  Banner._() : super();
  factory Banner({
    $core.String? imageURL,
    $core.String? uRL,
  }) {
    final _result = create();
    if (imageURL != null) {
      _result.imageURL = imageURL;
    }
    if (uRL != null) {
      _result.uRL = uRL;
    }
    return _result;
  }
  factory Banner.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Banner.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Banner clone() => Banner()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Banner copyWith(void Function(Banner) updates) => super.copyWith((message) => updates(message as Banner)) as Banner; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Banner create() => Banner._();
  Banner createEmptyInstance() => create();
  static $pb.PbList<Banner> createRepeated() => $pb.PbList<Banner>();
  @$core.pragma('dart2js:noInline')
  static Banner getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Banner>(create);
  static Banner? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get imageURL => $_getSZ(0);
  @$pb.TagNumber(1)
  set imageURL($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasImageURL() => $_has(0);
  @$pb.TagNumber(1)
  void clearImageURL() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get uRL => $_getSZ(1);
  @$pb.TagNumber(2)
  set uRL($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasURL() => $_has(1);
  @$pb.TagNumber(2)
  void clearURL() => clearField(2);
}

class BannerResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'BannerResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..pc<Banner>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'banners', $pb.PbFieldType.PM, subBuilder: Banner.create)
    ..hasRequiredFields = false
  ;

  BannerResponse._() : super();
  factory BannerResponse({
    $core.Iterable<Banner>? banners,
  }) {
    final _result = create();
    if (banners != null) {
      _result.banners.addAll(banners);
    }
    return _result;
  }
  factory BannerResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BannerResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  BannerResponse clone() => BannerResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  BannerResponse copyWith(void Function(BannerResponse) updates) => super.copyWith((message) => updates(message as BannerResponse)) as BannerResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BannerResponse create() => BannerResponse._();
  BannerResponse createEmptyInstance() => create();
  static $pb.PbList<BannerResponse> createRepeated() => $pb.PbList<BannerResponse>();
  @$core.pragma('dart2js:noInline')
  static BannerResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BannerResponse>(create);
  static BannerResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Banner> get banners => $_getList(0);
}

class BannerControllerApi {
  $pb.RpcClient _client;
  BannerControllerApi(this._client);

  $async.Future<BannerResponse> getBanner($pb.ClientContext? ctx, $1.Empty request) {
    var emptyResponse = BannerResponse();
    return _client.invoke<BannerResponse>(ctx, 'BannerController', 'GetBanner', request, emptyResponse);
  }
}


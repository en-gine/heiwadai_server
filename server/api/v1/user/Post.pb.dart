///
//  Generated code. Do not modify.
//  source: v1/user/Post.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/empty.pb.dart' as $1;

class PostRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PostRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', $pb.PbFieldType.OU3, protoName: 'ID')
    ..hasRequiredFields = false
  ;

  PostRequest._() : super();
  factory PostRequest({
    $core.int? iD,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    return _result;
  }
  factory PostRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PostRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PostRequest clone() => PostRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PostRequest copyWith(void Function(PostRequest) updates) => super.copyWith((message) => updates(message as PostRequest)) as PostRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PostRequest create() => PostRequest._();
  PostRequest createEmptyInstance() => create();
  static $pb.PbList<PostRequest> createRepeated() => $pb.PbList<PostRequest>();
  @$core.pragma('dart2js:noInline')
  static PostRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PostRequest>(create);
  static PostRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get iD => $_getIZ(0);
  @$pb.TagNumber(1)
  set iD($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);
}

class PostsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PostsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..pc<PostResponse>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'posts', $pb.PbFieldType.PM, subBuilder: PostResponse.create)
    ..hasRequiredFields = false
  ;

  PostsResponse._() : super();
  factory PostsResponse({
    $core.Iterable<PostResponse>? posts,
  }) {
    final _result = create();
    if (posts != null) {
      _result.posts.addAll(posts);
    }
    return _result;
  }
  factory PostsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PostsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PostsResponse clone() => PostsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PostsResponse copyWith(void Function(PostsResponse) updates) => super.copyWith((message) => updates(message as PostsResponse)) as PostsResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PostsResponse create() => PostsResponse._();
  PostsResponse createEmptyInstance() => create();
  static $pb.PbList<PostsResponse> createRepeated() => $pb.PbList<PostsResponse>();
  @$core.pragma('dart2js:noInline')
  static PostsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PostsResponse>(create);
  static PostsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<PostResponse> get posts => $_getList(0);
}

class PostResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PostResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.user'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', $pb.PbFieldType.OU3, protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Title', protoName: 'Title')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Content', protoName: 'Content')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Author', protoName: 'Author')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'PostDate', protoName: 'PostDate')
    ..hasRequiredFields = false
  ;

  PostResponse._() : super();
  factory PostResponse({
    $core.int? iD,
    $core.String? title,
    $core.String? content,
    $core.String? author,
    $core.String? postDate,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    if (title != null) {
      _result.title = title;
    }
    if (content != null) {
      _result.content = content;
    }
    if (author != null) {
      _result.author = author;
    }
    if (postDate != null) {
      _result.postDate = postDate;
    }
    return _result;
  }
  factory PostResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PostResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PostResponse clone() => PostResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PostResponse copyWith(void Function(PostResponse) updates) => super.copyWith((message) => updates(message as PostResponse)) as PostResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PostResponse create() => PostResponse._();
  PostResponse createEmptyInstance() => create();
  static $pb.PbList<PostResponse> createRepeated() => $pb.PbList<PostResponse>();
  @$core.pragma('dart2js:noInline')
  static PostResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PostResponse>(create);
  static PostResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get iD => $_getIZ(0);
  @$pb.TagNumber(1)
  set iD($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get content => $_getSZ(2);
  @$pb.TagNumber(3)
  set content($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasContent() => $_has(2);
  @$pb.TagNumber(3)
  void clearContent() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get author => $_getSZ(3);
  @$pb.TagNumber(4)
  set author($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAuthor() => $_has(3);
  @$pb.TagNumber(4)
  void clearAuthor() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get postDate => $_getSZ(4);
  @$pb.TagNumber(5)
  set postDate($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasPostDate() => $_has(4);
  @$pb.TagNumber(5)
  void clearPostDate() => clearField(5);
}

class PostControllerApi {
  $pb.RpcClient _client;
  PostControllerApi(this._client);

  $async.Future<PostResponse> getPostByID($pb.ClientContext? ctx, PostRequest request) {
    var emptyResponse = PostResponse();
    return _client.invoke<PostResponse>(ctx, 'PostController', 'GetPostByID', request, emptyResponse);
  }
  $async.Future<PostsResponse> getPosts($pb.ClientContext? ctx, $1.Empty request) {
    var emptyResponse = PostsResponse();
    return _client.invoke<PostsResponse>(ctx, 'PostController', 'GetPosts', request, emptyResponse);
  }
}


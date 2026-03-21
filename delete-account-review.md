# ユーザー削除機能（物理削除）実装サマリ & レビュー観点

## 概要

ユーザー（宿泊客）が自分のアカウントを削除できる API `DeleteAccount` を追加。
全データを物理削除し、削除後は同じメールアドレスで再登録可能にする。
**未来の予約（チェックアウト前）がある場合は削除不可。**

---

## 変更ファイル一覧

| # | ファイル | 変更内容 |
|---|---------|---------|
| 1 | `server/v1/user/Auth.proto` | `DeleteAccount` RPC 定義追加 |
| 2 | `server/api/v1/user/userconnect/Auth.connect.go` | `make buf` による自動生成 |
| 3 | `server/core/infra/repository/IUserRepository.go` | `HasFutureBooking` メソッド追加 |
| 4 | `server/infrastructure/repository/UserRepository.go` | `Delete()` に `user_book`, `user_report` の事前削除追加 + `HasFutureBooking()` 実装 |
| 5 | `server/core/usecase/user/AuthUsecase.go` | `DeleteAccount()` メソッド追加（未来予約チェック含む） |
| 6 | `server/controller/user/AuthController.go` | `DeleteAccount` ハンドラ追加 |

マイグレーション変更なし。新規ファイル作成なし。

---

## 変更の詳細

### 1. Proto 定義 (`Auth.proto`)

```proto
// アカウント削除
rpc DeleteAccount(google.protobuf.Empty) returns (google.protobuf.Empty) {}
```

- リクエストボディなし（`google.protobuf.Empty`）
- ユーザー ID は認証コンテキストから取得（`SignOut` と同パターン）

### 2. Repository インターフェース (`IUserRepository.go`)

```go
HasFutureBooking(id uuid.UUID) (bool, error)
```

- `IUserRepository` にメソッド追加

### 3. Repository 層 (`UserRepository.go`)

#### `HasFutureBooking()` — 新規メソッド

```go
func (ur *UserRepository) HasFutureBooking(userID uuid.UUID) (bool, error)
```

- `user_book` テーブルで `book_user_id = userID AND stay_to >= now()` の件数をカウント
- `StayTo`（チェックアウト日）が現在日時以降の予約が 1 件でもあれば `true` を返す

#### `Delete()` — 既存メソッド修正

先頭に以下を追加:

```go
// CASCADE なしのテーブルを先に削除（user_book, user_report）
_, err = models.UserBooks(models.UserBookWhere.BookUserID.EQ(userID.String())).DeleteAll(ctx, tran.Tran())
_, err = models.UserReports(models.UserReportWhere.UserID.EQ(userID.String())).DeleteAll(ctx, tran.Tran())
```

**背景:** DB の CASCADE 構造:

```
auth.users DELETE
  └─ トリガー: user_manager DELETE
       └─ CASCADE: user_data DELETE
            ├─ CASCADE: user_option
            ├─ CASCADE: checkin
            ├─ CASCADE: coupon_attached_user
            ├─ CASCADE: mail_magazine_log
            └─ ※ user_book, user_report は CASCADE なし → FK エラー！
```

→ `user_book` と `user_report` を先に削除すれば、残りは CASCADE で処理される。

### 4. Usecase 層 (`AuthUsecase.go`)

```go
func (u *AuthUsecase) DeleteAccount(userID uuid.UUID, token string) *errors.DomainError
```

処理順序:
1. ユーザー存在確認（`GetByID`）
2. **未来の予約チェック**（`HasFutureBooking`）→ あれば `UnPemitedOperation` エラー
3. 物理削除（`Delete`）
4. セッション無効化（`SignOut`）

### 5. Controller 層 (`AuthController.go`)

```go
func (ac *AuthController) DeleteAccount(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error)
```

- `ctx.Value(router.UserIDKey)` からユーザーID取得
- `ctx.Value(router.TokenKey)` からトークン取得
- `requireAuth` ミドルウェア付きで登録済みのため、ルーター変更不要

---

## 処理フロー

```
Client → DeleteAccount RPC（認証必須）
  │
  ├─ Controller: ctx から userID + token 取得
  │
  ├─ Usecase:
  │   ├─ GetByID(userID) → ユーザー存在確認
  │   ├─ HasFutureBooking(userID) → 未来の予約チェック
  │   │   └─ あり → エラー「チェックアウト前の予約があるため、アカウントを削除できません」
  │   │
  │   ├─ Repository.Delete(userID):
  │   │   ├─ DELETE user_book WHERE book_user_id = userID（過去の予約のみ残っている状態）
  │   │   ├─ DELETE user_report WHERE user_id = userID
  │   │   ├─ DELETE user_manager WHERE id = userID
  │   │   │     └─ CASCADE で user_data, user_option, checkin, coupon, mail_log 削除
  │   │   ├─ DELETE user_data（CASCADE で既に削除済み、ErrNoRows でスキップ）
  │   │   ├─ DELETE user_option（同上）
  │   │   └─ DELETE FROM auth.users WHERE id = userID
  │   │
  │   └─ SignOut(token) → セッション無効化
  │
  └─ メールアドレスが解放 → 再登録可能
```

---

## レビュー観点

### A. 機能面

- [ ] **FK 制約の網羅性**: `user_book` と `user_report` 以外に CASCADE のない FK が存在しないか？
  - DB スキーマ（マイグレーションファイル）を確認し、`user_data` を参照する全テーブルの `ON DELETE` 制約を網羅的にチェック
- [ ] **トランザクション整合性**: `user_book` / `user_report` 削除 → `user_manager` 削除 → `auth.users` 削除が全て同一トランザクション内で実行されるか
- [ ] **存在しないデータの削除**: `user_book` や `user_report` が 0 件の場合に `DeleteAll` がエラーにならないか（SQLBoiler の `DeleteAll` は 0 件でもエラーにならない仕様だが確認）
- [ ] **Admin 側への影響**: `UserRepository.Delete()` の変更は Admin の既存 Delete RPC にも適用される。Admin 側で予約履歴のあるユーザーを削除する際の FK エラーも同時に修正されるが、意図した動作か
- [ ] **未来予約の判定基準**: `StayTo >= now()` で判定。チェックアウト日当日は「未来の予約あり」として削除不可。この基準が妥当か
- [ ] **HasFutureBooking と Delete の間のレースコンディション**: チェックと削除の間に新しい予約が入る可能性。実運用上は問題になりにくいが、トランザクション分離レベルを確認

### B. セキュリティ面

- [ ] **認証チェック**: `DeleteAccount` が認証ミドルウェア（`requireAuth`）の対象であること。ルーター登録を確認
- [ ] **認可チェック**: 自分自身のアカウントのみ削除可能であること（他ユーザーの削除不可）。`ctx.Value(router.UserIDKey)` は認証トークンから取得した値なので問題ないはずだが確認
- [ ] **レート制限**: 削除 API に対する過度なリクエストへの対策が必要か
- [ ] **確認フロー**: クライアント側でパスワード再入力や確認ダイアログなどの UX 保護が別途必要か（API 側はリクエストボディなしだが、クライアント側の実装時に検討）

### C. データ整合性

- [ ] **キャッシュ無効化**: Redis にユーザー関連のキャッシュが残る可能性はないか。`SignOut` だけで十分か
- [ ] **外部サービス連携**: Supabase Storage にユーザーのファイルが残る可能性はないか
- [ ] **メルマガ配信停止**: `mail_magazine_log` は CASCADE で削除されるが、外部メール配信サービス側にユーザー情報が残る場合の対応は必要か

### D. コード品質

- [ ] **エラーメッセージ**: 日本語メッセージがクライアント仕様と統一されているか
- [ ] **SQL インジェクション**: `Delete()` 内の `fmt.Sprintf("DELETE FROM auth.users WHERE id = '%s'", userID.String())` は既存コードだが、`uuid.UUID.String()` は安全な値を返すため問題ないことを確認
- [ ] **未使用 import**: `Auth.proto` で `UserData.proto` と `timestamp.proto` を import しているが、`DeleteAccount` では使っていない（既存の他 RPC で使用しているため問題なし）

### E. テスト

- [ ] **統合テスト**: テストユーザー（予約なし）で `DeleteAccount` → DB 上の全関連テーブルが空であることを確認
- [ ] **未来予約ありテスト**: 未来の予約があるユーザーで `DeleteAccount` → エラーメッセージ確認
- [ ] **過去予約のみテスト**: 過去の予約のみのユーザーで `DeleteAccount` → 成功し、`user_book` も削除されること
- [ ] **再登録テスト**: 削除後に同一メールアドレスで `Register` → 成功すること
- [ ] **削除済みトークン**: 削除後のトークンでリクエスト → 認証エラーになること
- [ ] **既存テスト回帰**: `make test-all` で既存テストが全て通ること（特に `bookTest`）

---

## 検証手順

1. `make buf` で API 再生成（済）
2. `go build ./...` でコンパイルエラーなし確認（済）
3. `make lint` で静的解析（要 golangci-lint インストール）
4. `make test-all` で既存テスト回帰確認
5. 未来予約ありのテストユーザーで `DeleteAccount` → エラー確認
6. 予約なし or 過去予約のみのテストユーザーで `DeleteAccount` → 成功確認
7. DB で `user_data`, `user_manager`, `user_book`, `auth.users` が全て削除されていること確認
8. 同メールアドレスで新規 `Register` → 再登録成功確認

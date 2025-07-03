# Birthday Coupon Lambda Function

毎月1日に誕生日クーポンを自動発行するAWS Lambda関数です。

## 概要

- **実行スケジュール**: 毎月1日午前0時（UTC）/ 午前9時（JST）
- **機能**: HeiwadaiサーバーのCronエンドポイントを呼び出して誕生日クーポンを発行
- **ランタイム**: Node.js 18.x
- **タイムアウト**: 60秒
- **メモリ**: 128MB

## デプロイ方法

### 1. 環境変数の設定

```bash
export CRON_ACCESS_ENDPOINT="https://your-server.com/server.cron.CronCouponController/BulkIssueBirthdayCoupon"
export CRON_ACCESS_SECRET="your-authorization-token"
export CRON_ACCESS_KEY="your-cron-key"
```

### 2. 初回デプロイ

```bash
cd server
make deploy-birthday-coupon
```

このコマンドで以下が自動実行されます：
- IAMロールの作成
- Lambda関数のデプロイ
- EventBridgeルールの作成
- 権限設定

### 3. 更新デプロイ

```bash
cd server
make deploy-lambda
```

## 利用可能なコマンド

```bash
# ヘルプを表示
make help-lambda

# Lambda関数をテスト実行
make lambda-test

# Lambda関数のみデプロイ
make deploy-lambda

# EventBridgeのみデプロイ
make deploy-eventbridge

# 全リソースを削除
make delete-birthday-coupon
```

## 設定変更

### スケジュール変更

Makefileの`eventbridge-create-rule`内のcron式を変更：

```bash
# 毎月1日午前0時UTC（現在の設定）
--schedule-expression "cron(0 0 1 * ? *)"

# 毎月15日午前2時UTC
--schedule-expression "cron(0 2 15 * ? *)"
```

### AWS地域変更

```bash
export AWS_REGION="us-east-1"
make deploy-birthday-coupon
```

## モニタリング

### CloudWatchログ

Lambda実行ログは以下で確認できます：
```
/aws/lambda/heiwadai-birthday-coupon-issuer
```

### 手動実行

```bash
cd server
make lambda-test
```

## トラブルシューティング

### デプロイエラー

1. **AWS認証情報の確認**
   ```bash
   aws sts get-caller-identity
   ```

2. **必要な権限**
   - Lambda作成・更新権限
   - IAMロール作成権限
   - EventBridge作成権限

3. **環境変数の確認**
   ```bash
   echo $CRON_ACCESS_ENDPOINT
   echo $CRON_ACCESS_SECRET
   echo $CRON_ACCESS_KEY
   ```

### 実行エラー

CloudWatchログで詳細なエラー情報を確認してください：

```bash
aws logs describe-log-groups --log-group-name-prefix "/aws/lambda/heiwadai"
```

## コスト

- **Lambda実行**: 月1回 ≈ $0.000001
- **EventBridge**: 月1イベント ≈ $0.000001
- **CloudWatchログ**: 数KB ≈ $0.000001

**合計**: 月額 約$0.01以下
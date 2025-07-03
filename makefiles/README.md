# Makefile構造

Heiwadaiプロジェクトでは、Makefileを機能別に分割して管理しています。

## ファイル構成

```
Makefile                 # メインMakefile（ヘルプとinclude）
makefiles/
├── README.md           # このファイル
├── docker.mk           # Docker関連コマンド
├── server.mk           # サーバー開発関連コマンド
├── deployment.mk       # デプロイメント関連コマンド
└── lambda.mk           # AWS Lambda関連コマンド
```

## 各ファイルの説明

### `docker.mk`
- Docker Composeの基本操作
- コンテナの起動、停止、再起動
- ログ確認、シェルアクセス

### `server.mk`
- サーバー開発に必要なコマンド
- コード生成（buf、sqlboiler）
- データベース操作（migration、seeder）
- テスト実行

### `deployment.mk`
- ECRへのイメージpush
- プロト更新
- 本番デプロイ関連

### `lambda.mk`
- AWS Lambda関数のデプロイ
- EventBridge設定
- 誕生日クーポン自動発行システム

## 利用方法

### 基本コマンド
```bash
# ヘルプを表示
make help

# 開発環境起動
make up

# 開発サーバー起動
make dev
```

### Lambda関連
```bash
# Lambda詳細ヘルプ
make help-lambda

# Lambda デプロイ
make deploy-birthday-coupon
```

## 新しい機能を追加する場合

1. 適切なカテゴリの`.mk`ファイルを編集
2. 新しいカテゴリが必要な場合は新しい`.mk`ファイルを作成
3. メインMakefileの`include`に追加
4. `help`ターゲットを更新

## .PHONYターゲット

各`.mk`ファイルで適切に`.PHONY`ターゲットを設定して、ファイル名との競合を防いでいます。

## 環境変数

Lambda関連コマンドで必要な環境変数：
- `CRON_ACCESS_ENDPOINT`
- `CRON_ACCESS_SECRET` 
- `CRON_ACCESS_KEY`
- `AWS_REGION` (オプション)

デプロイ関連で必要な環境変数：
- `AWS_ID` (ECR関連)
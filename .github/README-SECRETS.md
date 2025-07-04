# GitHub Secrets Setup for ECR Build & Push

このファイルは、GitHub ActionsでECRへの自動ビルド・プッシュを行うために必要なSecretsの設定方法を説明します。

## 必須のGitHub Secrets

GitHub リポジトリの Settings > Secrets and variables > Actions で以下のSecretsを設定してください：

### AWS認証情報（最小限の設定）
```
AWS_ACCESS_KEY_ID=<AWS Access Key ID>
AWS_SECRET_ACCESS_KEY=<AWS Secret Access Key>
```

## 仕組み

### 🔄 自動化の流れ
1. **GitHub Actions**: masterブランチにプッシュ → ECRにDockerイメージをビルド・プッシュ
2. **App Runner**: ECRの新しいイメージを自動検知 → 自動デプロイ

### 🔒 セキュリティの利点
- 環境変数はApp Runnerサービス内に安全に保存済み
- GitHubには最小限のAWS認証情報のみ
- 本番環境の設定がGitリポジトリに露出しない

## セットアップ手順

### 1. AWS IAMユーザーの設定
ECRアクセス専用のIAMユーザーを作成：

**必要な権限（最小限）**:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ecr:GetAuthorizationToken",
                "ecr:BatchCheckLayerAvailability",
                "ecr:GetDownloadUrlForLayer",
                "ecr:BatchGetImage",
                "ecr:BatchImportLayerPart",
                "ecr:InitiateLayerUpload",
                "ecr:UploadLayerPart",
                "ecr:CompleteLayerUpload",
                "ecr:PutImage"
            ],
            "Resource": "*"
        }
    ]
}
```

### 2. GitHubリポジトリでSecretsを設定
1. **GitHubリポジトリにアクセス**
   - https://github.com/your-username/heiwadai にアクセス

2. **Secretsページを開く**
   - Settings > Secrets and variables > Actions をクリック

3. **AWS認証情報を追加**
   - `AWS_ACCESS_KEY_ID`: IAMユーザーのAccess Key ID
   - `AWS_SECRET_ACCESS_KEY`: IAMユーザーのSecret Access Key

## 動作確認

### 自動デプロイのテスト
1. masterブランチにコードをプッシュ
   ```bash
   git checkout master
   git merge develop
   git push origin master
   ```

2. **進行状況の確認**
   - GitHub ActionsのWorkflowsタブでビルド状況を確認
   - [App Runner Console](https://console.aws.amazon.com/apprunner/home?region=ap-northeast-1#/services)でデプロイ状況を確認

### 手動実行
緊急時やテスト時には、GitHub ActionsのWorkflowsページから手動でワークフローを実行可能（workflow_dispatch設定済み）

## App Runner自動デプロイの確認

App Runnerサービスで自動デプロイが有効になっているか確認：

```bash
# makefileで確認
make check-apprunner-status

# 直接AWS CLIで確認
aws apprunner describe-service \
  --service-arn <service-arn> \
  --query 'Service.SourceConfiguration.AutoDeploymentsEnabled'
```

`true`が返されれば自動デプロイが有効です。

## トラブルシューティング

### よくある問題

1. **ECRプッシュ権限エラー**
   - IAMユーザーの権限を確認
   - ECRリポジトリのアクセス権限を確認

2. **App Runnerが自動デプロイしない**
   - AutoDeploymentsEnabledがtrueになっているか確認
   - ECRリポジトリのイメージタグが`latest`になっているか確認

3. **ビルドエラー**
   - Dockerfileのパスが正しいか確認
   - ビルドコンテキストに必要なファイルが含まれているか確認

### 確認コマンド

```bash
# App Runnerサービスの状態確認
make check-apprunner-status

# ECRリポジトリのイメージ確認
aws ecr describe-images --repository-name heiwadai-server --region ap-northeast-1
```

## 環境変数の管理

環境変数はApp Runnerサービス内で管理されているため、変更が必要な場合は：

1. makefileを使用: `make update-apprunner-env`
2. AWS Console: App Runnerサービスの設定で直接変更
3. AWS CLI: `aws apprunner update-service`コマンドで変更

これにより、環境変数をGitに含めることなく安全に管理できます。
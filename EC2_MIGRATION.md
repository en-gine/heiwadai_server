# App Runner から EC2 への移行手順

## 概要
App Runner + NAT Gateway（月額$60-70）からEC2 + Elastic IP（月額$15-30）へ移行し、約75%のコスト削減を実現します。

## コスト比較

| 項目 | App Runner + NAT Gateway | EC2 + Elastic IP | 削減額 |
|------|-------------------------|------------------|--------|
| App Runner | $25-35/月 | $0 | $25-35 |
| NAT Gateway | $45/月 | $0 | $45 |
| EC2 インスタンス | $0 | $15-30/月 | - |
| Elastic IP | $0 | $0（1つまで無料） | $0 |
| **合計** | **$70-80/月** | **$15-30/月** | **$40-55/月** |

## 前提条件
- AWS CLI がセットアップ済み
- 適切なIAM権限を持つAWSプロファイル
- 既存のECRリポジトリとイメージ
- `.env.prod` ファイルが適切に設定済み

## 移行手順

### ステップ1: EC2環境のデプロイ

```bash
# 1. EC2インスタンス、Elastic IP、セキュリティグループを作成
make deploy-ec2

# 結果例:
# EC2 deployment completed!
# Fixed IP Address: 203.0.113.123
# 
# To deploy your application:
# 1. Wait a few minutes for the instance to fully initialize
# 2. Run: make deploy-to-ec2
```

### ステップ2: アプリケーションのデプロイ

```bash
# 3-5分待ってからアプリケーションをデプロイ
make deploy-to-ec2

# 結果例:
# ✅ Container is running successfully!
# ✅ Health check passed!
# Server is running on port 8080
```

### ステップ3: 動作確認

```bash
# EC2インスタンスの状態確認
make check-ec2-status

# アプリケーションの動作確認
curl http://203.0.113.123:8080/health

# SSH接続（必要に応じて）
make ssh-ec2
```

### ステップ4: DNS設定の更新

```bash
# 固定IPアドレスを確認
grep ELASTIC_IP_ADDRESS /tmp/ec2-resources.env

# DNS（Route 53など）でドメインのAレコードを新しいIPアドレスに変更
# 例: api.heiwadai-hotel.app → 203.0.113.123
```

### ステップ5: App RunnerとNAT Gatewayの削除

⚠️ **重要**: EC2での動作確認完了後に実行してください

```bash
# Option 1: App RunnerとNAT Gatewayのみ削除（Lambda用VPCは保持）
make cleanup-apprunner-only

# Option 2: 全て削除（Lambda用VPCも削除）※Lambdaも削除される
make cleanup-apprunner-and-vpc
```

## EC2管理コマンド

### 日常運用
```bash
# アプリケーションの再デプロイ
make deploy-to-ec2

# EC2インスタンスの状態確認
make check-ec2-status

# SSH接続
make ssh-ec2

# ログ確認（SSH接続後）
docker logs -f heiwadai-server
```

### トラブルシューティング
```bash
# コンテナの再起動
ssh -i ~/.ssh/heiwadai-server-key.pem ec2-user@<Elastic-IP>
docker restart heiwadai-server

# コンテナのシェルに接続
docker exec -it heiwadai-server /bin/sh

# システムリソース確認
htop
df -h
```

## セキュリティ設定

作成されるセキュリティグループ:
- SSH (22): 0.0.0.0/0
- HTTP (80): 0.0.0.0/0  
- HTTPS (443): 0.0.0.0/0
- App Port (8080): 0.0.0.0/0

必要に応じてSSHアクセスを特定IPに制限:
```bash
# 特定IPからのSSHのみ許可
aws ec2 authorize-security-group-ingress \
  --group-name heiwadai-server-sg \
  --protocol tcp \
  --port 22 \
  --cidr YOUR_IP/32
```

## バックアップとメンテナンス

### アプリケーションデータ
- PostgreSQL: Supabase側でバックアップ
- Redis: 必要に応じてRDBスナップショット
- ログ: CloudWatch Logsへ転送（オプション）

### EC2インスタンス
```bash
# AMIスナップショット作成
aws ec2 create-image \
  --instance-id $(grep INSTANCE_ID /tmp/ec2-resources.env | cut -d'=' -f2) \
  --name "heiwadai-server-backup-$(date +%Y%m%d)" \
  --description "Heiwadai server backup"
```

## スケーリング

### 縦スケーリング（インスタンスタイプ変更）
```bash
# インスタンス停止
aws ec2 stop-instances --instance-ids <instance-id>

# インスタンスタイプ変更
aws ec2 modify-instance-attribute \
  --instance-id <instance-id> \
  --instance-type Value=t3.medium

# インスタンス開始
aws ec2 start-instances --instance-ids <instance-id>
```

### 横スケーリング（ロードバランサー追加）
将来的に負荷が増えた場合:
1. Application Load Balancer作成
2. 複数EC2インスタンス作成
3. Auto Scaling Group設定

## 削除手順

```bash
# 全EC2リソースの削除
make delete-ec2
```

## ヘルプ

```bash
# EC2関連のコマンド一覧
make help-ec2

# App Runner関連のコマンド一覧
make help-apprunner
```

## トラブルシューティング

### EC2インスタンスが起動しない
1. セキュリティグループ設定確認
2. AMI IDが正しいか確認
3. インスタンスタイプがリージョンで利用可能か確認

### アプリケーションが起動しない
1. 環境変数設定確認
2. ECRログイン確認
3. ポート設定確認

### SSH接続できない
1. キーペアファイルの権限確認: `chmod 400 ~/.ssh/heiwadai-server-key.pem`
2. セキュリティグループでSSH許可確認
3. Elastic IP割り当て確認

## 参考情報

- EC2インスタンスタイプ: https://aws.amazon.com/ec2/instance-types/
- Elastic IP料金: https://aws.amazon.com/ec2/pricing/on-demand/
- セキュリティベストプラクティス: https://docs.aws.amazon.com/ec2/latest/userguide/ec2-security.html
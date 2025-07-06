#!/bin/bash

# Docker コンテナデプロイスクリプト
# EC2インスタンス上で実行される（Ubuntu用）

set -e

echo "Starting Docker deployment on Ubuntu..."

# Docker のインストール（Ubuntu用）
if ! command -v docker &> /dev/null; then
    echo "Installing Docker..."
    sudo apt-get update
    sudo apt-get install -y ca-certificates curl gnupg lsb-release
    sudo mkdir -p /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    sudo apt-get update
    sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin
    sudo usermod -aG docker ubuntu
    sudo systemctl enable docker
    sudo systemctl start docker
    echo "Docker installed successfully"
else
    echo "Docker is already installed"
fi

# AWS CLI のインストール（Ubuntu用）
if ! command -v aws &> /dev/null; then
    echo "Installing AWS CLI..."
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    sudo apt-get install -y unzip
    unzip awscliv2.zip
    sudo ./aws/install
    rm -rf aws awscliv2.zip
    echo "AWS CLI installed successfully"
else
    echo "AWS CLI is already installed"
fi

# 必要な環境変数のチェック
required_vars=(
    "AWS_REGION"
    "AWS_ACCOUNT_ID" 
    "ECR_REPOSITORY_NAME"
    "PORT"
    "PSQL_HOST"
    "PSQL_PORT"
    "PSQL_USER"
    "PSQL_PASS"
    "PSQL_DBNAME"
    "SUPABASE_URL"
    "SUPABASE_KEY"
    "REDISHOST"
    "REDISPORT"
)

for var in "${required_vars[@]}"; do
    if [ -z "${!var}" ]; then
        echo "Error: Environment variable $var is not set"
        exit 1
    fi
done

# ECRログイン
echo "Logging in to ECR..."
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com

# ECR イメージURL
ECR_IMAGE_URI="$AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPOSITORY_NAME:latest"

# 既存のコンテナを停止・削除
echo "Stopping existing containers..."
docker stop heiwadai-server 2>/dev/null || true
docker rm heiwadai-server 2>/dev/null || true

# 古いイメージを削除（ディスク容量節約）
echo "Cleaning up old images..."
docker image prune -f

# 最新イメージをプル
echo "Pulling latest image from ECR..."
docker pull $ECR_IMAGE_URI

# 環境変数ファイルの作成
cat > /home/ubuntu/.env << EOF
ENV_MODE=prod
PORT=$PORT
PSQL_HOST=$PSQL_HOST
PSQL_PORT=$PSQL_PORT
PSQL_USER=$PSQL_USER
PSQL_PASS=$PSQL_PASS
PSQL_DBNAME=$PSQL_DBNAME
SUPABASE_URL=$SUPABASE_URL
SUPABASE_KEY=$SUPABASE_KEY
SUPABASE_PROJECT_ID=${SUPABASE_PROJECT_ID:-}
SUPABASE_BUCKET=${SUPABASE_BUCKET:-}
ADMIN_AUTH_REDIRECT_URL=${ADMIN_AUTH_REDIRECT_URL:-}
USER_AUTH_REDIRECT_URL=${USER_AUTH_REDIRECT_URL:-}
ENCRYPT_KEY=${ENCRYPT_KEY:-}
REDISHOST=$REDISHOST
REDISPORT=$REDISPORT
REDISUSER=${REDISUSER:-}
REDISPASS=${REDISPASS:-}
MAIL_HOST=${MAIL_HOST:-}
MAIL_PORT=${MAIL_PORT:-}
MAIL_FROM=${MAIL_FROM:-}
MAIL_PASS=${MAIL_PASS:-}
MAIL_USER=${MAIL_USER:-}
TLBOOKING_IS_TEST=${TLBOOKING_IS_TEST:-}
TLBOOKING_AVAIL_API_URL=${TLBOOKING_AVAIL_API_URL:-}
TLBOOKING_BOOKING_API_URL=${TLBOOKING_BOOKING_API_URL:-}
TLBOOKING_CANCEL_API_URL=${TLBOOKING_CANCEL_API_URL:-}
TLBOOKING_USERNAME=${TLBOOKING_USERNAME:-}
TLBOOKING_PASSWORD=${TLBOOKING_PASSWORD:-}
TEST_USER_MAIL=${TEST_USER_MAIL:-}
TEST_USER_PASS=${TEST_USER_PASS:-}
TEST_ADMIN_MAIL=${TEST_ADMIN_MAIL:-}
CRON_ACCESS_KEY=${CRON_ACCESS_KEY:-}
CRON_ACCESS_SECRET=${CRON_ACCESS_SECRET:-}
S3_REGION=${S3_REGION:-}
S3_BUCKET=${S3_BUCKET:-}
S3_ACCESS_KEY=${S3_ACCESS_KEY:-}
S3_SECRET_KEY=${S3_SECRET_KEY:-}
EOF

# コンテナを起動
echo "Starting new container..."
docker run -d \
    --name heiwadai-server \
    --restart unless-stopped \
    -p $PORT:$PORT \
    --env-file /home/ubuntu/.env \
    --log-driver json-file \
    --log-opt max-size=10m \
    --log-opt max-file=3 \
    $ECR_IMAGE_URI

# ヘルスチェック
echo "Waiting for container to be healthy..."
sleep 10

# コンテナの状態確認
if docker ps | grep -q heiwadai-server; then
    echo "✅ Container is running successfully!"
    docker ps | grep heiwadai-server
    
    # ログの表示
    echo ""
    echo "Recent logs:"
    docker logs --tail 20 heiwadai-server
    
    # 簡単なヘルスチェック
    if curl -f http://localhost:$PORT/health 2>/dev/null; then
        echo "✅ Health check passed!"
    else
        echo "⚠️ Health check failed, but container is running"
    fi
    
else
    echo "❌ Container failed to start!"
    echo "Container logs:"
    docker logs heiwadai-server 2>/dev/null || echo "No logs available"
    exit 1
fi

echo "Deployment completed successfully!"
echo "Server is running on port $PORT"

# 便利なコマンドを表示
echo ""
echo "Useful commands:"
echo "  docker logs -f heiwadai-server  # View real-time logs"
echo "  docker restart heiwadai-server  # Restart container"
echo "  docker exec -it heiwadai-server /bin/sh  # Connect to container"
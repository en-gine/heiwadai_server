#!/bin/bash

# EC2インスタンス初期化スクリプト
# Amazon Linux 2023用

set -e

echo "Starting EC2 initialization..."

# システムアップデート
dnf update -y

# Docker のインストール
dnf install -y docker
systemctl start docker
systemctl enable docker

# ec2-user を docker グループに追加
usermod -a -G docker ec2-user

# AWS CLI のインストール（通常はプリインストール済み）
if ! command -v aws &> /dev/null; then
    echo "Installing AWS CLI..."
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    unzip awscliv2.zip
    ./aws/install
    rm -rf aws awscliv2.zip
fi

# Git のインストール
dnf install -y git

# Node.js のインストール（メンテナンス用）
dnf install -y nodejs npm

# 基本的なツールのインストール
dnf install -y htop curl wget unzip

# ログディレクトリの作成
mkdir -p /var/log/heiwadai
chown ec2-user:ec2-user /var/log/heiwadai

# Docker Compose のインストール
DOCKER_COMPOSE_VERSION="2.21.0"
curl -L "https://github.com/docker/compose/releases/download/v${DOCKER_COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# アプリケーション用ディレクトリの作成
mkdir -p /home/ec2-user/heiwadai
chown ec2-user:ec2-user /home/ec2-user/heiwadai

# 環境変数ファイルの作成
cat > /home/ec2-user/.bashrc << 'EOF'
# User specific aliases and functions
alias ll='ls -alF'
alias la='ls -A'
alias l='ls -CF'

# Docker関連のエイリアス
alias docker-ps='docker ps --format "table {{.ID}}\t{{.Image}}\t{{.Status}}\t{{.Ports}}\t{{.Names}}"'
alias docker-logs='docker logs -f'

# Heiwadai関連
alias heiwadai-logs='docker logs -f heiwadai-server'
alias heiwadai-restart='docker restart heiwadai-server'
alias heiwadai-status='docker ps | grep heiwadai'

export PATH=$PATH:/usr/local/bin
EOF

# CloudWatch Logs エージェントの設定（オプション）
dnf install -y amazon-cloudwatch-agent

# 初期化完了のマーク
echo "EC2 initialization completed at $(date)" > /home/ec2-user/init-completed.txt
chown ec2-user:ec2-user /home/ec2-user/init-completed.txt

echo "EC2 initialization script completed successfully!"
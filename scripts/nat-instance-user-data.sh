#!/bin/bash

# NAT Instance初期化スクリプト
# Amazon Linux 2023用

set -e

echo "Starting NAT Instance initialization..."

# システムアップデート
dnf update -y

# IPフォワーディングを有効化
echo 'net.ipv4.ip_forward = 1' >> /etc/sysctl.conf
echo 'net.ipv4.conf.all.send_redirects = 0' >> /etc/sysctl.conf
sysctl -p

# iptablesのインストールと設定
dnf install -y iptables-services

# NAT用のiptablesルールを設定
iptables -t nat -A POSTROUTING -o eth0 -s 10.0.0.0/16 -j MASQUERADE
iptables -A FORWARD -i eth0 -o eth0 -m state --state RELATED,ESTABLISHED -j ACCEPT
iptables -A FORWARD -i eth0 -o eth0 -s 10.0.0.0/16 -j ACCEPT

# iptablesの設定を保存
service iptables save

# iptablesサービスを有効化
systemctl enable iptables
systemctl start iptables

# CloudWatch Logs エージェントのインストール（監視用）
dnf install -y amazon-cloudwatch-agent

# 基本ツールのインストール
dnf install -y htop curl wget

# ログディレクトリの作成
mkdir -p /var/log/nat-instance
chmod 755 /var/log/nat-instance

# NAT Instance監視スクリプトの作成
cat > /home/ec2-user/check-nat.sh << 'EOF'
#!/bin/bash
echo "=== NAT Instance Status ==="
echo "Date: $(date)"
echo "IP Forwarding: $(cat /proc/sys/net/ipv4/ip_forward)"
echo "Network interfaces:"
ip addr show
echo ""
echo "iptables NAT rules:"
iptables -t nat -L -v
echo ""
echo "Active connections:"
netstat -an | grep :443 | head -5
EOF

chmod +x /home/ec2-user/check-nat.sh
chown ec2-user:ec2-user /home/ec2-user/check-nat.sh

# 起動時にNAT機能を有効化するサービスの作成
cat > /etc/systemd/system/nat-setup.service << 'EOF'
[Unit]
Description=NAT Instance Setup
After=network.target

[Service]
Type=oneshot
ExecStart=/bin/bash -c 'echo 1 > /proc/sys/net/ipv4/ip_forward'
ExecStart=/usr/sbin/iptables -t nat -A POSTROUTING -o eth0 -s 10.0.0.0/16 -j MASQUERADE
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target
EOF

systemctl enable nat-setup.service

# 初期化完了のマーク
echo "NAT Instance initialization completed at $(date)" > /home/ec2-user/nat-init-completed.txt
chown ec2-user:ec2-user /home/ec2-user/nat-init-completed.txt

echo "NAT Instance initialization script completed successfully!"

# 設定確認
echo "Current IP forwarding status: $(cat /proc/sys/net/ipv4/ip_forward)"
echo "Current iptables NAT rules:"
iptables -t nat -L
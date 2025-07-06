# AWS EC2関連のコマンド
AWS_REGION ?= ap-northeast-1
AWS_PROFILE ?= heiwadai
AWS_CMD = aws --profile $(AWS_PROFILE)

# EC2設定
EC2_INSTANCE_NAME ?= heiwadai-server
EC2_INSTANCE_TYPE ?= t3.small
EC2_KEY_NAME ?= heiwadai-server-key
EC2_SECURITY_GROUP_NAME ?= heiwadai-server-sg
EC2_AMI_ID ?= ami-0d52744d6551d851e  # Amazon Linux 2023 (ap-northeast-1)

# ECR設定（既存のものを利用）
ECR_REPOSITORY_NAME ?= heiwadai-server
ECR_IMAGE_URI ?= $(AWS_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$(ECR_REPOSITORY_NAME):latest

# Elastic IP設定
ELASTIC_IP_NAME ?= heiwadai-server-eip

include .env.prod

# EC2キーペアの作成
create-ec2-keypair:
	@echo "Creating EC2 Key Pair..."
	@if ! $(AWS_CMD) ec2 describe-key-pairs --key-names $(EC2_KEY_NAME) --region $(AWS_REGION) >/dev/null 2>&1; then \
		$(AWS_CMD) ec2 create-key-pair \
			--key-name $(EC2_KEY_NAME) \
			--query 'KeyMaterial' \
			--output text \
			--region $(AWS_REGION) > ~/.ssh/$(EC2_KEY_NAME).pem; \
		chmod 400 ~/.ssh/$(EC2_KEY_NAME).pem; \
		echo "Key pair created and saved to ~/.ssh/$(EC2_KEY_NAME).pem"; \
	else \
		echo "Key pair $(EC2_KEY_NAME) already exists"; \
	fi

# セキュリティグループの作成
create-ec2-security-group:
	@echo "Creating Security Group..."
	@if ! $(AWS_CMD) ec2 describe-security-groups --group-names $(EC2_SECURITY_GROUP_NAME) --region $(AWS_REGION) >/dev/null 2>&1; then \
		SG_ID=$$($(AWS_CMD) ec2 create-security-group \
			--group-name $(EC2_SECURITY_GROUP_NAME) \
			--description "Security group for Heiwadai server" \
			--query 'GroupId' \
			--output text \
			--region $(AWS_REGION)); \
		echo "Security Group created: $$SG_ID"; \
		echo "Adding inbound rules..."; \
		$(AWS_CMD) ec2 authorize-security-group-ingress \
			--group-id $$SG_ID \
			--protocol tcp \
			--port 22 \
			--cidr 0.0.0.0/0 \
			--region $(AWS_REGION); \
		$(AWS_CMD) ec2 authorize-security-group-ingress \
			--group-id $$SG_ID \
			--protocol tcp \
			--port 80 \
			--cidr 0.0.0.0/0 \
			--region $(AWS_REGION); \
		$(AWS_CMD) ec2 authorize-security-group-ingress \
			--group-id $$SG_ID \
			--protocol tcp \
			--port 443 \
			--cidr 0.0.0.0/0 \
			--region $(AWS_REGION); \
		$(AWS_CMD) ec2 authorize-security-group-ingress \
			--group-id $$SG_ID \
			--protocol tcp \
			--port $(PORT) \
			--cidr 0.0.0.0/0 \
			--region $(AWS_REGION); \
		echo "Security group rules added"; \
	else \
		echo "Security group $(EC2_SECURITY_GROUP_NAME) already exists"; \
	fi

# Elastic IPの作成（EC2用）
create-ec2-elastic-ip:
	@echo "Creating Elastic IP..."
	@if ! $(AWS_CMD) ec2 describe-addresses --filters "Name=tag:Name,Values=$(ELASTIC_IP_NAME)" --region $(AWS_REGION) --query 'Addresses[0]' --output text | grep -v None >/dev/null 2>&1; then \
		ALLOC_ID=$$($(AWS_CMD) ec2 allocate-address \
			--domain vpc \
			--query 'AllocationId' \
			--output text \
			--region $(AWS_REGION)); \
		$(AWS_CMD) ec2 create-tags \
			--resources $$ALLOC_ID \
			--tags Key=Name,Value=$(ELASTIC_IP_NAME) \
			--region $(AWS_REGION); \
		PUBLIC_IP=$$($(AWS_CMD) ec2 describe-addresses \
			--allocation-ids $$ALLOC_ID \
			--query 'Addresses[0].PublicIp' \
			--output text \
			--region $(AWS_REGION)); \
		echo "Elastic IP created: $$PUBLIC_IP ($$ALLOC_ID)"; \
		echo "ELASTIC_IP_ALLOCATION_ID=$$ALLOC_ID" > /tmp/ec2-resources.env; \
		echo "ELASTIC_IP_ADDRESS=$$PUBLIC_IP" >> /tmp/ec2-resources.env; \
	else \
		echo "Elastic IP $(ELASTIC_IP_NAME) already exists"; \
		ALLOC_ID=$$($(AWS_CMD) ec2 describe-addresses \
			--filters "Name=tag:Name,Values=$(ELASTIC_IP_NAME)" \
			--query 'Addresses[0].AllocationId' \
			--output text \
			--region $(AWS_REGION)); \
		PUBLIC_IP=$$($(AWS_CMD) ec2 describe-addresses \
			--allocation-ids $$ALLOC_ID \
			--query 'Addresses[0].PublicIp' \
			--output text \
			--region $(AWS_REGION)); \
		echo "Using existing Elastic IP: $$PUBLIC_IP ($$ALLOC_ID)"; \
		echo "ELASTIC_IP_ALLOCATION_ID=$$ALLOC_ID" > /tmp/ec2-resources.env; \
		echo "ELASTIC_IP_ADDRESS=$$PUBLIC_IP" >> /tmp/ec2-resources.env; \
	fi

# EC2インスタンスの作成
create-ec2-instance:
	@echo "Creating EC2 instance..."
	@SG_ID=$$($(AWS_CMD) ec2 describe-security-groups \
		--group-names $(EC2_SECURITY_GROUP_NAME) \
		--query 'SecurityGroups[0].GroupId' \
		--output text \
		--region $(AWS_REGION)); \
	echo "Using Security Group: $$SG_ID"; \
	INSTANCE_ID=$$($(AWS_CMD) ec2 run-instances \
		--image-id $(EC2_AMI_ID) \
		--count 1 \
		--instance-type $(EC2_INSTANCE_TYPE) \
		--key-name $(EC2_KEY_NAME) \
		--security-group-ids $$SG_ID \
		--tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=$(EC2_INSTANCE_NAME)}]' \
		--user-data file://scripts/ec2-user-data.sh \
		--query 'Instances[0].InstanceId' \
		--output text \
		--region $(AWS_REGION)); \
	echo "EC2 instance created: $$INSTANCE_ID"; \
	echo "INSTANCE_ID=$$INSTANCE_ID" >> /tmp/ec2-resources.env; \
	echo "Waiting for instance to be running..."; \
	$(AWS_CMD) ec2 wait instance-running \
		--instance-ids $$INSTANCE_ID \
		--region $(AWS_REGION); \
	echo "Instance is now running"

# Elastic IPをEC2インスタンスに関連付け
associate-ec2-elastic-ip:
	@echo "Associating Elastic IP to EC2 instance..."
	@. /tmp/ec2-resources.env; \
	$(AWS_CMD) ec2 associate-address \
		--instance-id $$INSTANCE_ID \
		--allocation-id $$ELASTIC_IP_ALLOCATION_ID \
		--region $(AWS_REGION); \
	echo "Elastic IP $$ELASTIC_IP_ADDRESS associated with instance $$INSTANCE_ID"

# EC2インスタンスへのSSH接続
ssh-ec2:
	@echo "Connecting to EC2 instance..."
	@. /tmp/ec2-resources.env; \
	ssh -i ~/.ssh/$(EC2_KEY_NAME).pem -o StrictHostKeyChecking=no ubuntu@$$ELASTIC_IP_ADDRESS

# EC2インスタンスの状態確認
check-ec2-status:
	@echo "Checking EC2 instance status..."
	@if [ -f /tmp/ec2-resources.env ]; then \
		. /tmp/ec2-resources.env; \
		$(AWS_CMD) ec2 describe-instances \
			--instance-ids $$INSTANCE_ID \
			--query 'Reservations[0].Instances[0].{State:State.Name,PublicIP:PublicIpAddress,PrivateIP:PrivateIpAddress,Type:InstanceType}' \
			--output table \
			--region $(AWS_REGION); \
	else \
		echo "No EC2 resources found. Run 'make deploy-ec2' first."; \
	fi

# Dockerコンテナのデプロイ
deploy-to-ec2:
	@echo "Deploying Docker container to EC2..."
	@. /tmp/ec2-resources.env; \
	echo "Copying deployment script to EC2..."; \
	scp -i ~/.ssh/$(EC2_KEY_NAME).pem -o StrictHostKeyChecking=no scripts/deploy-docker.sh ubuntu@$$ELASTIC_IP_ADDRESS:/home/ubuntu/; \
	echo "Copying AWS credentials to EC2..."; \
	ssh -i ~/.ssh/$(EC2_KEY_NAME).pem -o StrictHostKeyChecking=no ubuntu@$$ELASTIC_IP_ADDRESS "mkdir -p /home/ubuntu/.aws"; \
	scp -i ~/.ssh/$(EC2_KEY_NAME).pem -o StrictHostKeyChecking=no ~/.aws/credentials ubuntu@$$ELASTIC_IP_ADDRESS:/home/ubuntu/.aws/; \
	scp -i ~/.ssh/$(EC2_KEY_NAME).pem -o StrictHostKeyChecking=no ~/.aws/config ubuntu@$$ELASTIC_IP_ADDRESS:/home/ubuntu/.aws/; \
	echo "Running deployment on EC2..."; \
	ssh -i ~/.ssh/$(EC2_KEY_NAME).pem -o StrictHostKeyChecking=no ubuntu@$$ELASTIC_IP_ADDRESS \
		"chmod +x /home/ubuntu/deploy-docker.sh && \
		 AWS_REGION=$(AWS_REGION) \
		 AWS_ACCOUNT_ID=$(AWS_ID) \
		 ECR_REPOSITORY_NAME=$(ECR_REPOSITORY_NAME) \
		 PORT=$(PORT) \
		 PSQL_HOST=$(PSQL_HOST) \
		 PSQL_PORT=$(PSQL_PORT) \
		 PSQL_USER=$(PSQL_USER) \
		 PSQL_PASS='$(PSQL_PASS)' \
		 PSQL_DBNAME=$(PSQL_DBNAME) \
		 SUPABASE_URL=$(SUPABASE_URL) \
		 SUPABASE_KEY='$(SUPABASE_KEY)' \
		 SUPABASE_PROJECT_ID=$(SUPABASE_PROJECT_ID) \
		 SUPABASE_BUCKET=$(SUPABASE_BUCKET) \
		 ADMIN_AUTH_REDIRECT_URL='$(ADMIN_AUTH_REDIRECT_URL)' \
		 USER_AUTH_REDIRECT_URL='$(USER_AUTH_REDIRECT_URL)' \
		 ENCRYPT_KEY='$(ENCRYPT_KEY)' \
		 REDISHOST=$(REDISHOST) \
		 REDISPORT=$(REDISPORT) \
		 REDISUSER=$(REDISUSER) \
		 REDISPASS='$(REDISPASS)' \
		 MAIL_HOST=$(MAIL_HOST) \
		 MAIL_PORT=$(MAIL_PORT) \
		 MAIL_FROM='$(MAIL_FROM)' \
		 MAIL_USER='$(MAIL_USER)' \
		 MAIL_PASS='$(MAIL_PASS)' \
		 TLBOOKING_IS_TEST=$(TLBOOKING_IS_TEST) \
		 TLBOOKING_AVAIL_API_URL='$(TLBOOKING_AVAIL_API_URL)' \
		 TLBOOKING_BOOKING_API_URL='$(TLBOOKING_BOOKING_API_URL)' \
		 TLBOOKING_CANCEL_API_URL='$(TLBOOKING_CANCEL_API_URL)' \
		 TLBOOKING_USERNAME='$(TLBOOKING_USERNAME)' \
		 TLBOOKING_PASSWORD='$(TLBOOKING_PASSWORD)' \
		 TEST_USER_MAIL='$(TEST_USER_MAIL)' \
		 TEST_USER_PASS='$(TEST_USER_PASS)' \
		 TEST_ADMIN_MAIL='$(TEST_ADMIN_MAIL)' \
		 CRON_ACCESS_KEY='$(CRON_ACCESS_KEY)' \
		 CRON_ACCESS_SECRET='$(CRON_ACCESS_SECRET)' \
		 /home/ubuntu/deploy-docker.sh"

# EC2完全デプロイ
deploy-ec2: create-ec2-keypair create-ec2-security-group create-ec2-elastic-ip create-ec2-instance associate-ec2-elastic-ip
	@echo "EC2 deployment completed!"
	@echo "Fixed IP Address:"
	@. /tmp/ec2-resources.env && echo "$$ELASTIC_IP_ADDRESS"
	@echo ""
	@echo "To deploy your application:"
	@echo "1. Wait a few minutes for the instance to fully initialize"
	@echo "2. Run: make deploy-to-ec2"
	@echo "3. Your server will be available at: http://$$ELASTIC_IP_ADDRESS:$(PORT)"

# EC2インスタンスの削除
delete-ec2-instance:
	@echo "Deleting EC2 instance..."
	@if [ -f /tmp/ec2-resources.env ]; then \
		. /tmp/ec2-resources.env; \
		$(AWS_CMD) ec2 terminate-instances \
			--instance-ids $$INSTANCE_ID \
			--region $(AWS_REGION); \
		echo "Waiting for instance termination..."; \
		$(AWS_CMD) ec2 wait instance-terminated \
			--instance-ids $$INSTANCE_ID \
			--region $(AWS_REGION); \
		echo "Instance terminated"; \
	else \
		echo "No EC2 instance found"; \
	fi

# Elastic IPの削除
delete-ec2-elastic-ip:
	@echo "Releasing Elastic IP..."
	@if [ -f /tmp/ec2-resources.env ]; then \
		. /tmp/ec2-resources.env; \
		$(AWS_CMD) ec2 release-address \
			--allocation-id $$ELASTIC_IP_ALLOCATION_ID \
			--region $(AWS_REGION); \
		echo "Elastic IP released"; \
	else \
		echo "No Elastic IP found"; \
	fi

# セキュリティグループの削除
delete-ec2-security-group:
	@echo "Deleting Security Group..."
	@if $(AWS_CMD) ec2 describe-security-groups --group-names $(EC2_SECURITY_GROUP_NAME) --region $(AWS_REGION) >/dev/null 2>&1; then \
		SG_ID=$$($(AWS_CMD) ec2 describe-security-groups \
			--group-names $(EC2_SECURITY_GROUP_NAME) \
			--query 'SecurityGroups[0].GroupId' \
			--output text \
			--region $(AWS_REGION)); \
		$(AWS_CMD) ec2 delete-security-group \
			--group-id $$SG_ID \
			--region $(AWS_REGION); \
		echo "Security Group deleted"; \
	else \
		echo "Security group not found"; \
	fi

# キーペアの削除
delete-ec2-keypair:
	@echo "Deleting Key Pair..."
	@$(AWS_CMD) ec2 delete-key-pair \
		--key-name $(EC2_KEY_NAME) \
		--region $(AWS_REGION) || true
	@rm -f ~/.ssh/$(EC2_KEY_NAME).pem
	@echo "Key pair deleted"

# EC2リソースの完全削除
delete-ec2: delete-ec2-instance delete-ec2-elastic-ip delete-ec2-security-group delete-ec2-keypair
	@echo "All EC2 resources deleted"
	@rm -f /tmp/ec2-resources.env

# ヘルプ
help-ec2:
	@echo "EC2 Deployment Commands:"
	@echo "  deploy-ec2             - EC2インスタンスの完全デプロイ"
	@echo "  deploy-to-ec2          - Dockerコンテナのデプロイ"
	@echo "  check-ec2-status       - EC2インスタンスの状態確認"
	@echo "  ssh-ec2                - EC2インスタンスにSSH接続"
	@echo "  delete-ec2             - 全EC2リソースの削除"
	@echo ""
	@echo "Individual Commands:"
	@echo "  create-ec2-keypair     - キーペア作成"
	@echo "  create-ec2-security-group - セキュリティグループ作成"
	@echo "  create-ec2-elastic-ip  - Elastic IP作成"
	@echo "  create-ec2-instance    - EC2インスタンス作成"
	@echo "  associate-ec2-elastic-ip - Elastic IP関連付け"
	@echo ""
	@echo "Required environment variables:"
	@echo "  AWS_ID                 - AWS Account ID"
	@echo "  All database and app settings from .env.prod"
	@echo ""
	@echo "Optional variables:"
	@echo "  EC2_INSTANCE_TYPE      - Instance type (default: t3.small)"
	@echo "  EC2_INSTANCE_NAME      - Instance name (default: heiwadai-server)"

.PHONY: create-ec2-keypair create-ec2-security-group create-ec2-elastic-ip create-ec2-instance \
        associate-ec2-elastic-ip ssh-ec2 check-ec2-status deploy-to-ec2 deploy-ec2 \
        delete-ec2-instance delete-ec2-elastic-ip delete-ec2-security-group delete-ec2-keypair \
        delete-ec2 help-ec2
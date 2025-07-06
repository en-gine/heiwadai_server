# NAT Instance 関連のコマンド（Lambda固定IP用）
AWS_REGION ?= ap-northeast-1
AWS_PROFILE ?= default
AWS_CMD = aws --profile $(AWS_PROFILE)

# NAT Instance設定
NAT_INSTANCE_NAME ?= heiwadai-nat-instance
NAT_INSTANCE_TYPE ?= t3.nano
NAT_AMI_ID ?= ami-00d101850e971728d  # Amazon Linux 2023 NAT AMI
NAT_KEY_NAME ?= heiwadai-nat-key
NAT_SECURITY_GROUP_NAME ?= heiwadai-nat-sg

# NAT Instance用キーペアの作成
create-nat-keypair:
	@echo "Creating NAT Instance Key Pair..."
	@if ! $(AWS_CMD) ec2 describe-key-pairs --key-names $(NAT_KEY_NAME) --region $(AWS_REGION) >/dev/null 2>&1; then \
		$(AWS_CMD) ec2 create-key-pair \
			--key-name $(NAT_KEY_NAME) \
			--query 'KeyMaterial' \
			--output text \
			--region $(AWS_REGION) > ~/.ssh/$(NAT_KEY_NAME).pem; \
		chmod 400 ~/.ssh/$(NAT_KEY_NAME).pem; \
		echo "NAT Instance key pair created"; \
	else \
		echo "Key pair $(NAT_KEY_NAME) already exists"; \
	fi

# NAT Instance用セキュリティグループの作成
create-nat-security-group:
	@echo "Creating NAT Instance Security Group..."
	@if [ -f /tmp/vpc-resources.env ]; then \
		. /tmp/vpc-resources.env; \
		if ! $(AWS_CMD) ec2 describe-security-groups --group-names $(NAT_SECURITY_GROUP_NAME) --region $(AWS_REGION) >/dev/null 2>&1; then \
			SG_ID=$$($(AWS_CMD) ec2 create-security-group \
				--group-name $(NAT_SECURITY_GROUP_NAME) \
				--description "Security group for NAT Instance" \
				--vpc-id $$VPC_ID \
				--query 'GroupId' \
				--output text \
				--region $(AWS_REGION)); \
			echo "NAT Security Group created: $$SG_ID"; \
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
				--source-group $$LAMBDA_SECURITY_GROUP_ID \
				--region $(AWS_REGION); \
			$(AWS_CMD) ec2 authorize-security-group-ingress \
				--group-id $$SG_ID \
				--protocol tcp \
				--port 443 \
				--source-group $$LAMBDA_SECURITY_GROUP_ID \
				--region $(AWS_REGION); \
			$(AWS_CMD) ec2 authorize-security-group-ingress \
				--group-id $$SG_ID \
				--protocol tcp \
				--port 587 \
				--source-group $$LAMBDA_SECURITY_GROUP_ID \
				--region $(AWS_REGION); \
			echo "NAT_INSTANCE_SECURITY_GROUP_ID=$$SG_ID" >> /tmp/vpc-resources.env; \
		else \
			echo "Security group $(NAT_SECURITY_GROUP_NAME) already exists"; \
		fi; \
	else \
		echo "❌ VPC resources not found. Please ensure VPC exists."; \
		exit 1; \
	fi

# NAT Instance専用Elastic IPの作成
create-nat-elastic-ip:
	@echo "Creating NAT Instance Elastic IP..."
	@NAT_EIP_NAME="$(NAT_INSTANCE_NAME)-eip"; \
	if ! $(AWS_CMD) ec2 describe-addresses --filters "Name=tag:Name,Values=$$NAT_EIP_NAME" --region $(AWS_REGION) --query 'Addresses[0]' --output text | grep -v None >/dev/null 2>&1; then \
		ALLOC_ID=$$($(AWS_CMD) ec2 allocate-address \
			--domain vpc \
			--query 'AllocationId' \
			--output text \
			--region $(AWS_REGION)); \
		$(AWS_CMD) ec2 create-tags \
			--resources $$ALLOC_ID \
			--tags Key=Name,Value=$$NAT_EIP_NAME \
			--region $(AWS_REGION); \
		PUBLIC_IP=$$($(AWS_CMD) ec2 describe-addresses \
			--allocation-ids $$ALLOC_ID \
			--query 'Addresses[0].PublicIp' \
			--output text \
			--region $(AWS_REGION)); \
		echo "NAT Instance Elastic IP created: $$PUBLIC_IP ($$ALLOC_ID)"; \
		echo "NAT_ELASTIC_IP_ALLOCATION_ID=$$ALLOC_ID" >> /tmp/vpc-resources.env; \
		echo "NAT_ELASTIC_IP_ADDRESS=$$PUBLIC_IP" >> /tmp/vpc-resources.env; \
	else \
		echo "NAT Instance Elastic IP already exists"; \
		ALLOC_ID=$$($(AWS_CMD) ec2 describe-addresses \
			--filters "Name=tag:Name,Values=$$NAT_EIP_NAME" \
			--query 'Addresses[0].AllocationId' \
			--output text \
			--region $(AWS_REGION)); \
		PUBLIC_IP=$$($(AWS_CMD) ec2 describe-addresses \
			--allocation-ids $$ALLOC_ID \
			--query 'Addresses[0].PublicIp' \
			--output text \
			--region $(AWS_REGION)); \
		echo "Using existing NAT Instance Elastic IP: $$PUBLIC_IP ($$ALLOC_ID)"; \
		echo "NAT_ELASTIC_IP_ALLOCATION_ID=$$ALLOC_ID" >> /tmp/vpc-resources.env; \
		echo "NAT_ELASTIC_IP_ADDRESS=$$PUBLIC_IP" >> /tmp/vpc-resources.env; \
	fi

# NAT Instanceの作成
create-nat-instance:
	@echo "Creating NAT Instance..."
	@if [ -f /tmp/vpc-resources.env ]; then \
		. /tmp/vpc-resources.env; \
		INSTANCE_ID=$$($(AWS_CMD) ec2 run-instances \
			--image-id $(NAT_AMI_ID) \
			--count 1 \
			--instance-type $(NAT_INSTANCE_TYPE) \
			--key-name $(NAT_KEY_NAME) \
			--security-group-ids $$NAT_INSTANCE_SECURITY_GROUP_ID \
			--subnet-id $$PUBLIC_SUBNET_ID \
			--associate-public-ip-address \
			--tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=$(NAT_INSTANCE_NAME)}]' \
			--user-data file://scripts/nat-instance-user-data.sh \
			--query 'Instances[0].InstanceId' \
			--output text \
			--region $(AWS_REGION)); \
		echo "NAT Instance created: $$INSTANCE_ID"; \
		echo "NAT_INSTANCE_ID=$$INSTANCE_ID" >> /tmp/vpc-resources.env; \
		echo "Waiting for instance to be running..."; \
		$(AWS_CMD) ec2 wait instance-running \
			--instance-ids $$INSTANCE_ID \
			--region $(AWS_REGION); \
		echo "Disabling source/destination check..."; \
		$(AWS_CMD) ec2 modify-instance-attribute \
			--instance-id $$INSTANCE_ID \
			--no-source-dest-check \
			--region $(AWS_REGION); \
		echo "NAT Instance is ready"; \
	else \
		echo "❌ VPC resources not found"; \
		exit 1; \
	fi

# NAT InstanceにElastic IPを関連付け
associate-nat-elastic-ip:
	@echo "Associating Elastic IP to NAT Instance..."
	@. /tmp/vpc-resources.env; \
	$(AWS_CMD) ec2 associate-address \
		--instance-id $$NAT_INSTANCE_ID \
		--allocation-id $$NAT_ELASTIC_IP_ALLOCATION_ID \
		--region $(AWS_REGION); \
	echo "NAT Instance Elastic IP $$NAT_ELASTIC_IP_ADDRESS associated"

# プライベートルートテーブルをNAT Instanceに変更
update-route-to-nat-instance:
	@echo "Updating route table to use NAT Instance..."
	@if [ -f /tmp/vpc-resources.env ]; then \
		. /tmp/vpc-resources.env; \
		ROUTE_TABLE_ID=$$($(AWS_CMD) ec2 describe-route-tables \
			--filters "Name=vpc-id,Values=$$VPC_ID" "Name=tag:Name,Values=heiwadai-vpc-private-rt" \
			--query 'RouteTables[0].RouteTableId' \
			--output text \
			--region $(AWS_REGION)); \
		if [ "$$ROUTE_TABLE_ID" != "None" ]; then \
			$(AWS_CMD) ec2 replace-route \
				--route-table-id $$ROUTE_TABLE_ID \
				--destination-cidr-block 0.0.0.0/0 \
				--instance-id $$NAT_INSTANCE_ID \
				--region $(AWS_REGION); \
			echo "Route table updated to use NAT Instance"; \
		else \
			echo "❌ Private route table not found"; \
			exit 1; \
		fi; \
	else \
		echo "❌ VPC resources not found"; \
		exit 1; \
	fi

# NAT Instanceの完全デプロイ
deploy-nat-instance: create-nat-keypair create-nat-security-group create-nat-elastic-ip create-nat-instance associate-nat-elastic-ip update-route-to-nat-instance
	@echo "NAT Instance deployment completed!"
	@echo "Fixed IP for Lambda traffic:"
	@. /tmp/vpc-resources.env && echo "$$NAT_ELASTIC_IP_ADDRESS"
	@echo ""
	@echo "Monthly cost: ~$5 (t3.nano)"
	@echo "SendGrid IP whitelist: $$NAT_ELASTIC_IP_ADDRESS"

# NAT Instanceの削除
delete-nat-instance:
	@echo "Deleting NAT Instance..."
	@if [ -f /tmp/vpc-resources.env ]; then \
		. /tmp/vpc-resources.env; \
		if [ -n "$$NAT_INSTANCE_ID" ]; then \
			$(AWS_CMD) ec2 terminate-instances \
				--instance-ids $$NAT_INSTANCE_ID \
				--region $(AWS_REGION); \
			echo "Waiting for NAT Instance termination..."; \
			$(AWS_CMD) ec2 wait instance-terminated \
				--instance-ids $$NAT_INSTANCE_ID \
				--region $(AWS_REGION); \
		fi; \
		if [ -n "$$NAT_ELASTIC_IP_ALLOCATION_ID" ]; then \
			$(AWS_CMD) ec2 release-address \
				--allocation-id $$NAT_ELASTIC_IP_ALLOCATION_ID \
				--region $(AWS_REGION); \
		fi; \
		if [ -n "$$NAT_INSTANCE_SECURITY_GROUP_ID" ]; then \
			$(AWS_CMD) ec2 delete-security-group \
				--group-id $$NAT_INSTANCE_SECURITY_GROUP_ID \
				--region $(AWS_REGION); \
		fi; \
		$(AWS_CMD) ec2 delete-key-pair \
			--key-name $(NAT_KEY_NAME) \
			--region $(AWS_REGION) || true; \
		rm -f ~/.ssh/$(NAT_KEY_NAME).pem; \
		echo "NAT Instance deleted"; \
	else \
		echo "No NAT Instance found"; \
	fi

# ヘルプ
help-nat-instance:
	@echo "NAT Instance Commands (for Lambda fixed IP):"
	@echo "  deploy-nat-instance    - NAT Instance完全デプロイ"
	@echo "  delete-nat-instance    - NAT Instance削除"
	@echo ""
	@echo "Cost: ~$5/month (t3.nano)"
	@echo "Purpose: Lambda → NAT Instance → Fixed IP → SendGrid"
	@echo ""
	@echo "After deployment:"
	@echo "1. Update SendGrid IP whitelist with the fixed IP"
	@echo "2. Test Lambda function: make lambda-test"

.PHONY: create-nat-keypair create-nat-security-group create-nat-elastic-ip \
        create-nat-instance associate-nat-elastic-ip update-route-to-nat-instance \
        deploy-nat-instance delete-nat-instance help-nat-instance
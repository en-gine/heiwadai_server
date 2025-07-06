# AWS App Runner関連のコマンド
AWS_REGION ?= ap-northeast-1
AWS_PROFILE ?= heiwadai
AWS_CMD = aws --profile $(AWS_PROFILE)

# App Runner設定
APP_RUNNER_SERVICE_NAME ?= heiwadai-server
ECR_IMAGE_URI ?= $(AWS_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/heiwadai-server:latest
APP_RUNNER_PORT ?= $(PORT)
# App RunnerのCPU/メモリサイズ
# CPU: 256, 512, 1024, 2048, 4096 または 0.25 vCPU, 0.5 vCPU, 1 vCPU, 2 vCPU, 4 vCPU
# Memory: 512, 1024, 2048, 4096
APP_RUNNER_CPU ?= 256
APP_RUNNER_MEMORY ?= 512

# App Runner用のIAMロール名
APP_RUNNER_ECR_ROLE ?= heiwadai-apprunner-ecr-access-role
APP_RUNNER_INSTANCE_ROLE ?= heiwadai-apprunner-instance-role

# VPC設定
VPC_NAME ?= heiwadai-vpc
VPC_CIDR ?= 10.0.0.0/16
PUBLIC_SUBNET_CIDR ?= 10.0.1.0/24
PRIVATE_SUBNET_CIDR ?= 10.0.2.0/24
NAT_GATEWAY_NAME ?= heiwadai-nat-gateway
VPC_CONNECTOR_NAME ?= heiwadai-vpc-connector

include .env.prod

# App Runner ECRアクセスロールの作成
create-apprunner-ecr-role:
	@echo "Creating App Runner ECR access role..."
	$(AWS_CMD) iam create-role \
		--role-name $(APP_RUNNER_ECR_ROLE) \
		--assume-role-policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"Service":"build.apprunner.amazonaws.com"},"Action":"sts:AssumeRole"}]}' \
		--region $(AWS_REGION) || true
	@echo "Attaching ECR access policy..."
	$(AWS_CMD) iam attach-role-policy \
		--role-name $(APP_RUNNER_ECR_ROLE) \
		--policy-arn arn:aws:iam::aws:policy/service-role/AWSAppRunnerServicePolicyForECRAccess \
		--region $(AWS_REGION) || true
	@echo "Creating custom ECR policy for specific repository..."
	@echo '{ \
		"Version": "2012-10-17", \
		"Statement": [ \
			{ \
				"Effect": "Allow", \
				"Action": [ \
					"ecr:BatchCheckLayerAvailability", \
					"ecr:GetDownloadUrlForLayer", \
					"ecr:BatchGetImage", \
					"ecr:DescribeRepositories", \
					"ecr:DescribeImages" \
				], \
				"Resource": "arn:aws:ecr:$(AWS_REGION):$(AWS_ID):repository/heiwadai-server" \
			}, \
			{ \
				"Effect": "Allow", \
				"Action": [ \
					"ecr:GetAuthorizationToken" \
				], \
				"Resource": "*" \
			} \
		] \
	}' > /tmp/ecr-policy.json
	$(AWS_CMD) iam put-role-policy \
		--role-name $(APP_RUNNER_ECR_ROLE) \
		--policy-name HeiwadaiECRAccess \
		--policy-document file:///tmp/ecr-policy.json \
		--region $(AWS_REGION) || true
	@rm -f /tmp/ecr-policy.json

# App Runnerインスタンスロールの作成
create-apprunner-instance-role:
	@echo "Creating App Runner instance role..."
	$(AWS_CMD) iam create-role \
		--role-name $(APP_RUNNER_INSTANCE_ROLE) \
		--assume-role-policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"Service":"tasks.apprunner.amazonaws.com"},"Action":"sts:AssumeRole"}]}' \
		--region $(AWS_REGION) || true
	@echo "Note: Add any required policies for your app (e.g., S3, RDS access)"

# App Runner設定ファイルの作成
create-apprunner-config:
	@echo "Creating App Runner configuration..."
	@cat > /tmp/apprunner-config.yaml << EOF
	version: 1.0
	runtime: docker
	build:
	  commands:
	    build:
	      - echo "No build commands for pre-built image"
	run:
	  runtime-version: latest
	  command: /server
	  network:
	    port: $(APP_RUNNER_PORT)
	    env: PORT
	  env:
	    - name: "PSQL_HOST"
	      value: "$(PSQL_HOST)"
	    - name: "PSQL_PORT"
	      value: "$(PSQL_PORT)"
	    - name: "PSQL_USER"
	      value: "$(PSQL_USER)"
	    - name: "PSQL_PASS"
	      value: "$(PSQL_PASS)"
	    - name: "PSQL_DBNAME"
	      value: "$(PSQL_DBNAME)"
	    - name: "REDIS_HOST"
	      value: "$(REDIS_HOST)"
	    - name: "REDIS_PORT"
	      value: "$(REDIS_PORT)"
	    - name: "SUPABASE_URL"
	      value: "$(SUPABASE_URL)"
	    - name: "SUPABASE_KEY"
	      value: "$(SUPABASE_KEY)"
	EOF
	@echo "Configuration saved to /tmp/apprunner-config.yaml"

# App Runnerサービスの作成
create-apprunner-service:
ifndef AWS_ID
	$(error AWS_ID is not set. Please set your AWS account ID)
endif
ifndef PSQL_HOST
	$(error PSQL_HOST is not set. Please set database configuration)
endif
	@echo "Creating App Runner service..."
	$(eval ECR_ROLE_ARN := $(shell $(AWS_CMD) iam get-role --role-name $(APP_RUNNER_ECR_ROLE) --query Role.Arn --output text))
	$(eval INSTANCE_ROLE_ARN := $(shell $(AWS_CMD) iam get-role --role-name $(APP_RUNNER_INSTANCE_ROLE) --query Role.Arn --output text))
	@echo "ECR Role ARN: $(ECR_ROLE_ARN)"
	@echo "Instance Role ARN: $(INSTANCE_ROLE_ARN)"
	@echo "ECR Image: $(ECR_IMAGE_URI)"
	@echo "Creating source configuration JSON..."
	@echo '{ \
		"AuthenticationConfiguration": { \
			"AccessRoleArn": "$(ECR_ROLE_ARN)" \
		}, \
		"AutoDeploymentsEnabled": false, \
		"ImageRepository": { \
			"ImageIdentifier": "$(ECR_IMAGE_URI)", \
			"ImageConfiguration": { \
				"Port": "$(APP_RUNNER_PORT)", \
				"RuntimeEnvironmentVariables": {' > /tmp/apprunner-source.json
	@# 全ての必須環境変数（値をエスケープ）
	@printf '"ENV_MODE": "%s",\n' "$$(echo '$(ENV_MODE)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"PORT": "%s",\n' "$$(echo '$(PORT)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"PSQL_HOST": "%s",\n' "$$(echo '$(PSQL_HOST)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"PSQL_PORT": "%s",\n' "$$(echo '$(PSQL_PORT)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"PSQL_USER": "%s",\n' "$$(echo '$(PSQL_USER)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"PSQL_PASS": "%s",\n' "$$(echo '$(PSQL_PASS)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"PSQL_DBNAME": "%s",\n' "$$(echo '$(PSQL_DBNAME)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"SUPABASE_URL": "%s",\n' "$$(echo '$(SUPABASE_URL)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"SUPABASE_KEY": "%s",\n' "$$(echo '$(SUPABASE_KEY)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"SUPABASE_PROJECT_ID": "%s",\n' "$$(echo '$(SUPABASE_PROJECT_ID)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"SUPABASE_BUCKET": "%s",\n' "$$(echo '$(SUPABASE_BUCKET)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"ADMIN_AUTH_REDIRECT_URL": "%s",\n' "$$(echo '$(ADMIN_AUTH_REDIRECT_URL)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"USER_AUTH_REDIRECT_URL": "%s",\n' "$$(echo '$(USER_AUTH_REDIRECT_URL)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"ENCRYPT_KEY": "%s",\n' "$$(echo '$(ENCRYPT_KEY)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"REDISHOST": "%s",\n' "$$(echo '$(REDISHOST)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"REDISPORT": "%s",\n' "$$(echo '$(REDISPORT)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"REDISUSER": "%s",\n' "$$(echo '$(REDISUSER)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"REDISPASS": "%s",\n' "$$(echo '$(REDISPASS)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"MAIL_HOST": "%s",\n' "$$(echo '$(MAIL_HOST)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"MAIL_PORT": "%s",\n' "$$(echo '$(MAIL_PORT)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"MAIL_FROM": "%s",\n' "$$(echo '$(MAIL_FROM)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"MAIL_PASS": "%s",\n' "$$(echo '$(MAIL_PASS)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"MAIL_USER": "%s",\n' "$$(echo '$(MAIL_USER)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"TLBOOKING_IS_TEST": "%s",\n' "$$(echo '$(TLBOOKING_IS_TEST)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"TLBOOKING_AVAIL_API_URL": "%s",\n' "$$(echo '$(TLBOOKING_AVAIL_API_URL)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"TLBOOKING_BOOKING_API_URL": "%s",\n' "$$(echo '$(TLBOOKING_BOOKING_API_URL)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"TLBOOKING_CANCEL_API_URL": "%s",\n' "$$(echo '$(TLBOOKING_CANCEL_API_URL)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"TLBOOKING_USERNAME": "%s",\n' "$$(echo '$(TLBOOKING_USERNAME)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"TLBOOKING_PASSWORD": "%s",\n' "$$(echo '$(TLBOOKING_PASSWORD)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"TEST_USER_MAIL": "%s",\n' "$$(echo '$(TEST_USER_MAIL)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"TEST_USER_PASS": "%s",\n' "$$(echo '$(TEST_USER_PASS)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"TEST_ADMIN_MAIL": "%s",\n' "$$(echo '$(TEST_ADMIN_MAIL)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"CRON_ACCESS_KEY": "%s",\n' "$$(echo '$(CRON_ACCESS_KEY)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"CRON_ACCESS_SECRET": "%s",\n' "$$(echo '$(CRON_ACCESS_SECRET)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"S3_REGION": "%s",\n' "$$(echo '$(S3_REGION)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"S3_BUCKET": "%s",\n' "$$(echo '$(S3_BUCKET)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"S3_ACCESS_KEY": "%s",\n' "$$(echo '$(S3_ACCESS_KEY)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@printf '"S3_SECRET_KEY": "%s"\n' "$$(echo '$(S3_SECRET_KEY)' | sed 's/"/\\"/g')" >> /tmp/apprunner-source.json
	@echo '} \
			}, \
			"ImageRepositoryType": "ECR" \
		} \
	}' >> /tmp/apprunner-source.json
	@echo '{ \
		"Cpu": "$(APP_RUNNER_CPU)", \
		"Memory": "$(APP_RUNNER_MEMORY)", \
		"InstanceRoleArn": "$(INSTANCE_ROLE_ARN)" \
	}' > /tmp/apprunner-instance.json
	@echo "Generated instance configuration:"
	@cat /tmp/apprunner-instance.json
	@echo ""
	@echo "Generated source configuration:"
	@cat /tmp/apprunner-source.json
	@echo ""
	$(AWS_CMD) apprunner create-service \
		--service-name $(APP_RUNNER_SERVICE_NAME) \
		--source-configuration file:///tmp/apprunner-source.json \
		--instance-configuration file:///tmp/apprunner-instance.json \
		--region $(AWS_REGION)
	@rm -f /tmp/apprunner-source.json /tmp/apprunner-instance.json

# サービスの更新（新しいイメージをデプロイ）
update-apprunner-service:
	@echo "Updating App Runner service with new image..."
	$(eval SERVICE_ARN := $(shell $(AWS_CMD) apprunner list-services --query "ServiceSummaryList[?ServiceName=='$(APP_RUNNER_SERVICE_NAME)'].ServiceArn" --output text))
	$(AWS_CMD) apprunner update-service \
		--service-arn $(SERVICE_ARN) \
		--source-configuration '{ \
			"ImageRepository": { \
				"ImageIdentifier": "'$(ECR_IMAGE_URI)'", \
				"ImageConfiguration": { \
					"Port": "'$(APP_RUNNER_PORT)'" \
				}, \
				"ImageRepositoryType": "ECR" \
			} \
		}' \
		--region $(AWS_REGION)

# App Runnerサービスの状態確認
check-apprunner-status:
	@echo "Checking App Runner service status..."
	$(eval SERVICE_ARN := $(shell $(AWS_CMD) apprunner list-services --query "ServiceSummaryList[?ServiceName=='$(APP_RUNNER_SERVICE_NAME)'].ServiceArn" --output text))
	@$(AWS_CMD) apprunner describe-service \
		--service-arn $(SERVICE_ARN) \
		--query '{Status: Service.Status, URL: Service.ServiceUrl, UpdatedAt: Service.UpdatedAt}' \
		--output table \
		--region $(AWS_REGION)

# App Runnerサービスの削除
delete-apprunner-service:
	@echo "Deleting App Runner service..."
	$(eval SERVICE_ARN := $(shell $(AWS_CMD) apprunner list-services --query "ServiceSummaryList[?ServiceName=='$(APP_RUNNER_SERVICE_NAME)'].ServiceArn" --output text))
	$(AWS_CMD) apprunner delete-service \
		--service-arn $(SERVICE_ARN) \
		--region $(AWS_REGION)

# 環境変数の更新
update-apprunner-env:
	@echo "Updating environment variables..."
	$(eval SERVICE_ARN := $(shell $(AWS_CMD) apprunner list-services --query "ServiceSummaryList[?ServiceName=='$(APP_RUNNER_SERVICE_NAME)'].ServiceArn" --output text))
	$(AWS_CMD) apprunner update-service \
		--service-arn $(SERVICE_ARN) \
		--source-configuration '{ \
			"ImageRepository": { \
				"ImageConfiguration": { \
					"RuntimeEnvironmentVariables": { \
						"PSQL_HOST": "'$(PSQL_HOST)'", \
						"PSQL_PORT": "'$(PSQL_PORT)'", \
						"PSQL_USER": "'$(PSQL_USER)'", \
						"PSQL_DBNAME": "'$(PSQL_DBNAME)'", \
						"REDIS_HOST": "'$(REDIS_HOST)'", \
						"REDIS_PORT": "'$(REDIS_PORT)'", \
						"SUPABASE_URL": "'$(SUPABASE_URL)'", \
						"SUPABASE_KEY": "'$(SUPABASE_KEY)'" \
					} \
				}, \
				"ImageRepositoryType": "ECR" \
			} \
		}' \
		--region $(AWS_REGION)

# カスタムドメインの設定
associate-custom-domain:
ifndef CUSTOM_DOMAIN
	$(error CUSTOM_DOMAIN is not set. Example: api.heiwadai-hotel.app)
endif
	@echo "Associating custom domain..."
	$(eval SERVICE_ARN := $(shell $(AWS_CMD) apprunner list-services --query "ServiceSummaryList[?ServiceName=='$(APP_RUNNER_SERVICE_NAME)'].ServiceArn" --output text))
	$(AWS_CMD) apprunner associate-custom-domain \
		--service-arn $(SERVICE_ARN) \
		--domain-name $(CUSTOM_DOMAIN) \
		--region $(AWS_REGION)

# 完全デプロイ（初回）
deploy-apprunner: create-apprunner-ecr-role create-apprunner-instance-role create-apprunner-service
	@echo "App Runner deployment completed!"
	@echo "Checking service status..."
	@sleep 10
	@$(MAKE) check-apprunner-status

# 更新デプロイ（ECRに新しいイメージがある場合）
deploy-apprunner-update: push update-apprunner-service
	@echo "App Runner update deployment completed!"
	@sleep 10
	@$(MAKE) check-apprunner-status

# ヘルプ
help-apprunner:
	@echo "App Runner Deployment Commands:"
	@echo "  deploy-apprunner       - 初回デプロイ（サービス作成）"
	@echo "  deploy-apprunner-update - 更新デプロイ（新しいイメージ）"
	@echo "  deploy-apprunner-with-fixed-ip - 固定IPアドレス付きデプロイ"
	@echo "  check-apprunner-status - サービス状態確認"
	@echo "  update-apprunner-env   - 環境変数の更新"
	@echo "  associate-custom-domain - カスタムドメイン設定"
	@echo "  delete-apprunner-service - サービス削除"
	@echo ""
	@echo "VPC/Fixed IP Commands:"
	@echo "  setup-vpc-for-apprunner - VPC環境のセットアップ（固定IP用）"
	@echo "  apply-vpc-connector    - App RunnerサービスにVPCコネクタを適用"
	@echo "  delete-vpc-resources   - VPCリソースの削除"
	@echo ""
	@echo "Required environment variables:"
	@echo "  AWS_ID                - AWS Account ID"
	@echo "  PSQL_* / REDIS_*      - Database settings"
	@echo "  SUPABASE_*            - Supabase settings"
	@echo ""
	@echo "Optional variables:"
	@echo "  APP_RUNNER_SERVICE_NAME - Service name (default: heiwadai-server)"
	@echo "  APP_RUNNER_CPU         - CPU allocation (default: 0.25 vCPU)"
	@echo "  APP_RUNNER_MEMORY      - Memory allocation (default: 0.5 GB)"
	@echo "  CUSTOM_DOMAIN          - Custom domain for the service"
	@echo "  VPC_NAME              - VPC name (default: heiwadai-vpc)"
	@echo "  VPC_CIDR              - VPC CIDR block (default: 10.0.0.0/16)"

# VPCの作成
create-vpc:
	@echo "Creating VPC..."
	$(eval VPC_ID := $(shell $(AWS_CMD) ec2 create-vpc \
		--cidr-block $(VPC_CIDR) \
		--tag-specifications 'ResourceType=vpc,Tags=[{Key=Name,Value=$(VPC_NAME)}]' \
		--query Vpc.VpcId \
		--output text \
		--region $(AWS_REGION)))
	@echo "VPC created: $(VPC_ID)"
	@echo "Enabling DNS hostnames..."
	@$(AWS_CMD) ec2 modify-vpc-attribute \
		--vpc-id $(VPC_ID) \
		--enable-dns-hostnames \
		--region $(AWS_REGION)
	@echo "VPC_ID=$(VPC_ID)" > /tmp/vpc-resources.env

# パブリックサブネットの作成
create-public-subnet:
	@echo "Creating public subnet..."
	$(eval VPC_ID := $(shell grep VPC_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	$(eval PUBLIC_SUBNET_ID := $(shell $(AWS_CMD) ec2 create-subnet \
		--vpc-id $(VPC_ID) \
		--cidr-block $(PUBLIC_SUBNET_CIDR) \
		--availability-zone $(AWS_REGION)a \
		--tag-specifications 'ResourceType=subnet,Tags=[{Key=Name,Value=$(VPC_NAME)-public-subnet}]' \
		--query Subnet.SubnetId \
		--output text \
		--region $(AWS_REGION)))
	@echo "Public subnet created: $(PUBLIC_SUBNET_ID)"
	@echo "PUBLIC_SUBNET_ID=$(PUBLIC_SUBNET_ID)" >> /tmp/vpc-resources.env

# プライベートサブネットの作成
create-private-subnet:
	@echo "Creating private subnet..."
	$(eval VPC_ID := $(shell grep VPC_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	$(eval PRIVATE_SUBNET_ID := $(shell $(AWS_CMD) ec2 create-subnet \
		--vpc-id $(VPC_ID) \
		--cidr-block $(PRIVATE_SUBNET_CIDR) \
		--availability-zone $(AWS_REGION)a \
		--tag-specifications 'ResourceType=subnet,Tags=[{Key=Name,Value=$(VPC_NAME)-private-subnet}]' \
		--query Subnet.SubnetId \
		--output text \
		--region $(AWS_REGION)))
	@echo "Private subnet created: $(PRIVATE_SUBNET_ID)"
	@echo "PRIVATE_SUBNET_ID=$(PRIVATE_SUBNET_ID)" >> /tmp/vpc-resources.env

# インターネットゲートウェイの作成と接続
create-internet-gateway:
	@echo "Creating Internet Gateway..."
	$(eval VPC_ID := $(shell grep VPC_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	$(eval IGW_ID := $(shell $(AWS_CMD) ec2 create-internet-gateway \
		--tag-specifications 'ResourceType=internet-gateway,Tags=[{Key=Name,Value=$(VPC_NAME)-igw}]' \
		--query InternetGateway.InternetGatewayId \
		--output text \
		--region $(AWS_REGION)))
	@echo "Internet Gateway created: $(IGW_ID)"
	@echo "Attaching to VPC..."
	@$(AWS_CMD) ec2 attach-internet-gateway \
		--vpc-id $(VPC_ID) \
		--internet-gateway-id $(IGW_ID) \
		--region $(AWS_REGION)
	@echo "IGW_ID=$(IGW_ID)" >> /tmp/vpc-resources.env

# Elastic IPの作成
create-elastic-ip:
	@echo "Allocating Elastic IP..."
	$(eval EIP_ALLOC_ID := $(shell $(AWS_CMD) ec2 allocate-address \
		--domain vpc \
		--tag-specifications 'ResourceType=elastic-ip,Tags=[{Key=Name,Value=$(VPC_NAME)-nat-eip}]' \
		--query AllocationId \
		--output text \
		--region $(AWS_REGION)))
	@echo "Elastic IP allocated: $(EIP_ALLOC_ID)"
	@echo "EIP_ALLOC_ID=$(EIP_ALLOC_ID)" >> /tmp/vpc-resources.env
	$(eval EIP_ADDRESS := $(shell $(AWS_CMD) ec2 describe-addresses \
		--allocation-ids $(EIP_ALLOC_ID) \
		--query 'Addresses[0].PublicIp' \
		--output text \
		--region $(AWS_REGION)))
	@echo "Elastic IP address: $(EIP_ADDRESS)"
	@echo "EIP_ADDRESS=$(EIP_ADDRESS)" >> /tmp/vpc-resources.env

# NAT Gatewayの作成
create-nat-gateway:
	@echo "Creating NAT Gateway..."
	$(eval PUBLIC_SUBNET_ID := $(shell grep PUBLIC_SUBNET_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	$(eval EIP_ALLOC_ID := $(shell grep EIP_ALLOC_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	$(eval NAT_GATEWAY_ID := $(shell $(AWS_CMD) ec2 create-nat-gateway \
		--subnet-id $(PUBLIC_SUBNET_ID) \
		--allocation-id $(EIP_ALLOC_ID) \
		--tag-specifications 'ResourceType=natgateway,Tags=[{Key=Name,Value=$(NAT_GATEWAY_NAME)}]' \
		--query NatGateway.NatGatewayId \
		--output text \
		--region $(AWS_REGION)))
	@echo "NAT Gateway created: $(NAT_GATEWAY_ID)"
	@echo "NAT_GATEWAY_ID=$(NAT_GATEWAY_ID)" >> /tmp/vpc-resources.env
	@echo "Waiting for NAT Gateway to become available..."
	@$(AWS_CMD) ec2 wait nat-gateway-available \
		--nat-gateway-ids $(NAT_GATEWAY_ID) \
		--region $(AWS_REGION)
	@echo "NAT Gateway is available"

# ルートテーブルの設定
configure-route-tables:
	@echo "Configuring route tables..."
	$(eval VPC_ID := $(shell grep VPC_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	$(eval IGW_ID := $(shell grep IGW_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	$(eval NAT_GATEWAY_ID := $(shell grep NAT_GATEWAY_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	$(eval PUBLIC_SUBNET_ID := $(shell grep PUBLIC_SUBNET_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	$(eval PRIVATE_SUBNET_ID := $(shell grep PRIVATE_SUBNET_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	# パブリックルートテーブルの作成
	$(eval PUBLIC_ROUTE_TABLE_ID := $(shell $(AWS_CMD) ec2 create-route-table \
		--vpc-id $(VPC_ID) \
		--tag-specifications 'ResourceType=route-table,Tags=[{Key=Name,Value=$(VPC_NAME)-public-rt}]' \
		--query RouteTable.RouteTableId \
		--output text \
		--region $(AWS_REGION)))
	@echo "Public route table created: $(PUBLIC_ROUTE_TABLE_ID)"
	# インターネットゲートウェイへのルート追加
	@$(AWS_CMD) ec2 create-route \
		--route-table-id $(PUBLIC_ROUTE_TABLE_ID) \
		--destination-cidr-block 0.0.0.0/0 \
		--gateway-id $(IGW_ID) \
		--region $(AWS_REGION)
	# パブリックサブネットとの関連付け
	@$(AWS_CMD) ec2 associate-route-table \
		--route-table-id $(PUBLIC_ROUTE_TABLE_ID) \
		--subnet-id $(PUBLIC_SUBNET_ID) \
		--region $(AWS_REGION)
	# プライベートルートテーブルの作成
	$(eval PRIVATE_ROUTE_TABLE_ID := $(shell $(AWS_CMD) ec2 create-route-table \
		--vpc-id $(VPC_ID) \
		--tag-specifications 'ResourceType=route-table,Tags=[{Key=Name,Value=$(VPC_NAME)-private-rt}]' \
		--query RouteTable.RouteTableId \
		--output text \
		--region $(AWS_REGION)))
	@echo "Private route table created: $(PRIVATE_ROUTE_TABLE_ID)"
	# NAT Gatewayへのルート追加
	@$(AWS_CMD) ec2 create-route \
		--route-table-id $(PRIVATE_ROUTE_TABLE_ID) \
		--destination-cidr-block 0.0.0.0/0 \
		--nat-gateway-id $(NAT_GATEWAY_ID) \
		--region $(AWS_REGION)
	# プライベートサブネットとの関連付け
	@$(AWS_CMD) ec2 associate-route-table \
		--route-table-id $(PRIVATE_ROUTE_TABLE_ID) \
		--subnet-id $(PRIVATE_SUBNET_ID) \
		--region $(AWS_REGION)
	@echo "Route tables configured successfully"

# VPCコネクタの作成
create-vpc-connector:
	@echo "Creating VPC connector for App Runner..."
	$(eval PRIVATE_SUBNET_ID := $(shell grep PRIVATE_SUBNET_ID /tmp/vpc-resources.env | cut -d'=' -f2))
	$(eval VPC_CONNECTOR_ARN := $(shell $(AWS_CMD) apprunner create-vpc-connector \
		--vpc-connector-name $(VPC_CONNECTOR_NAME) \
		--subnets $(PRIVATE_SUBNET_ID) \
		--query VpcConnector.VpcConnectorArn \
		--output text \
		--region $(AWS_REGION)))
	@echo "VPC Connector created: $(VPC_CONNECTOR_ARN)"
	@echo "VPC_CONNECTOR_ARN=$(VPC_CONNECTOR_ARN)" >> /tmp/vpc-resources.env

# VPC設定の完全実行
setup-vpc-for-apprunner: create-vpc create-public-subnet create-private-subnet create-internet-gateway create-elastic-ip create-nat-gateway configure-route-tables create-vpc-connector
	@echo "VPC setup completed!"
	@echo "Fixed Elastic IP Address:"
	@grep EIP_ADDRESS /tmp/vpc-resources.env | cut -d'=' -f2
	@echo "Resources saved to /tmp/vpc-resources.env"

# App RunnerサービスにVPCコネクタを適用
apply-vpc-connector:
	@echo "Applying VPC connector to App Runner service..."
	$(eval SERVICE_ARN := $(shell $(AWS_CMD) apprunner list-services --query "ServiceSummaryList[?ServiceName=='$(APP_RUNNER_SERVICE_NAME)'].ServiceArn" --output text))
	$(eval VPC_CONNECTOR_ARN := $(shell grep VPC_CONNECTOR_ARN /tmp/vpc-resources.env | cut -d'=' -f2))
	@$(AWS_CMD) apprunner update-service \
		--service-arn $(SERVICE_ARN) \
		--network-configuration '{ \
			"EgressConfiguration": { \
				"EgressType": "VPC", \
				"VpcConnectorArn": "$(VPC_CONNECTOR_ARN)" \
			} \
		}' \
		--region $(AWS_REGION)
	@echo "VPC connector applied successfully"

# App RunnerサービスからVPCコネクタを削除
remove-vpc-connector-from-apprunner:
	@echo "Removing VPC connector from App Runner service..."
	$(eval SERVICE_ARN := $(shell $(AWS_CMD) apprunner list-services --query "ServiceSummaryList[?ServiceName=='$(APP_RUNNER_SERVICE_NAME)'].ServiceArn" --output text))
	@if [ -n "$(SERVICE_ARN)" ]; then \
		$(AWS_CMD) apprunner update-service \
			--service-arn $(SERVICE_ARN) \
			--network-configuration '{ \
				"EgressConfiguration": { \
					"EgressType": "DEFAULT" \
				} \
			}' \
			--region $(AWS_REGION); \
		echo "VPC connector removed from App Runner. Service will use default internet access."; \
	else \
		echo "App Runner service not found"; \
	fi

# VPCリソースの削除
delete-vpc-resources:
	@echo "Deleting VPC resources..."
	@if [ -f /tmp/vpc-resources.env ]; then \
		echo "Loading resource IDs..."; \
		. /tmp/vpc-resources.env; \
		if [ -n "$$VPC_CONNECTOR_ARN" ]; then \
			echo "Deleting VPC Connector..."; \
			$(AWS_CMD) apprunner delete-vpc-connector --vpc-connector-arn $$VPC_CONNECTOR_ARN --region $(AWS_REGION) || true; \
			echo "Waiting for VPC Connector deletion..."; \
			sleep 30; \
		fi; \
		if [ -n "$$NAT_GATEWAY_ID" ]; then \
			echo "Deleting NAT Gateway..."; \
			$(AWS_CMD) ec2 delete-nat-gateway --nat-gateway-id $$NAT_GATEWAY_ID --region $(AWS_REGION) || true; \
			echo "Waiting for NAT Gateway deletion (this may take a few minutes)..."; \
			sleep 120; \
		fi; \
		if [ -n "$$EIP_ALLOC_ID" ]; then \
			echo "Releasing Elastic IP..."; \
			$(AWS_CMD) ec2 release-address --allocation-id $$EIP_ALLOC_ID --region $(AWS_REGION) || true; \
		fi; \
		if [ -n "$$IGW_ID" ] && [ -n "$$VPC_ID" ]; then \
			echo "Detaching Internet Gateway..."; \
			$(AWS_CMD) ec2 detach-internet-gateway --internet-gateway-id $$IGW_ID --vpc-id $$VPC_ID --region $(AWS_REGION) || true; \
			echo "Deleting Internet Gateway..."; \
			$(AWS_CMD) ec2 delete-internet-gateway --internet-gateway-id $$IGW_ID --region $(AWS_REGION) || true; \
		fi; \
		if [ -n "$$PUBLIC_SUBNET_ID" ]; then \
			echo "Deleting public subnet..."; \
			$(AWS_CMD) ec2 delete-subnet --subnet-id $$PUBLIC_SUBNET_ID --region $(AWS_REGION) || true; \
		fi; \
		if [ -n "$$PRIVATE_SUBNET_ID" ]; then \
			echo "Deleting private subnet..."; \
			$(AWS_CMD) ec2 delete-subnet --subnet-id $$PRIVATE_SUBNET_ID --region $(AWS_REGION) || true; \
		fi; \
		if [ -n "$$LAMBDA_SECURITY_GROUP_ID" ]; then \
			echo "Deleting Lambda security group..."; \
			$(AWS_CMD) ec2 delete-security-group --group-id $$LAMBDA_SECURITY_GROUP_ID --region $(AWS_REGION) || true; \
		fi; \
		echo "Waiting before VPC deletion..."; \
		sleep 30; \
		if [ -n "$$VPC_ID" ]; then \
			echo "Deleting VPC..."; \
			$(AWS_CMD) ec2 delete-vpc --vpc-id $$VPC_ID --region $(AWS_REGION) || true; \
		fi; \
		rm -f /tmp/vpc-resources.env; \
		echo "VPC resources deletion completed"; \
	else \
		echo "No VPC resources file found"; \
	fi

# NAT Gatewayのみ削除（Lambda用VPCは保持）
delete-nat-gateway-only:
	@echo "Deleting NAT Gateway only (keeping VPC for Lambda)..."
	@if [ -f /tmp/vpc-resources.env ]; then \
		echo "Loading resource IDs..."; \
		. /tmp/vpc-resources.env; \
		if [ -n "$$VPC_CONNECTOR_ARN" ]; then \
			echo "Deleting VPC Connector..."; \
			$(AWS_CMD) apprunner delete-vpc-connector --vpc-connector-arn $$VPC_CONNECTOR_ARN --region $(AWS_REGION) || true; \
			echo "Waiting for VPC Connector deletion..."; \
			sleep 30; \
		fi; \
		if [ -n "$$NAT_GATEWAY_ID" ]; then \
			echo "Deleting NAT Gateway..."; \
			$(AWS_CMD) ec2 delete-nat-gateway --nat-gateway-id $$NAT_GATEWAY_ID --region $(AWS_REGION) || true; \
			echo "Waiting for NAT Gateway deletion (this may take a few minutes)..."; \
			sleep 120; \
		fi; \
		if [ -n "$$EIP_ALLOC_ID" ]; then \
			echo "Releasing Elastic IP for NAT Gateway..."; \
			$(AWS_CMD) ec2 release-address --allocation-id $$EIP_ALLOC_ID --region $(AWS_REGION) || true; \
		fi; \
		echo "NAT Gateway deletion completed"; \
		echo "VPC, subnets, and security groups are kept for Lambda function"; \
		echo "Birthday Coupon Lambda can continue to work in the VPC"; \
	else \
		echo "No VPC resources file found"; \
	fi

# App Runner削除 + NAT Gateway削除（VPCは保持）
cleanup-apprunner-only: remove-vpc-connector-from-apprunner delete-apprunner-service delete-nat-gateway-only
	@echo "App Runner cleanup completed! VPC and Lambda are preserved."

# 完全クリーンアップ（App Runner + VPC）
cleanup-apprunner-and-vpc: remove-vpc-connector-from-apprunner delete-apprunner-service delete-vpc-resources
	@echo "App Runner and VPC cleanup completed!"

# 固定IPを使用したApp Runnerデプロイ（初回）
deploy-apprunner-with-fixed-ip: setup-vpc-for-apprunner deploy-apprunner apply-vpc-connector
	@echo "App Runner deployment with fixed IP completed!"
	@echo "Fixed Elastic IP Address:"
	@grep EIP_ADDRESS /tmp/vpc-resources.env | cut -d'=' -f2

.PHONY: create-apprunner-ecr-role create-apprunner-instance-role create-apprunner-config \
        create-apprunner-service update-apprunner-service check-apprunner-status \
        delete-apprunner-service update-apprunner-env associate-custom-domain \
        deploy-apprunner deploy-apprunner-update help-apprunner \
        create-vpc create-public-subnet create-private-subnet create-internet-gateway \
        create-elastic-ip create-nat-gateway configure-route-tables create-vpc-connector \
        setup-vpc-for-apprunner apply-vpc-connector delete-vpc-resources deploy-apprunner-with-fixed-ip \
        remove-vpc-connector-from-apprunner cleanup-apprunner-and-vpc delete-nat-gateway-only cleanup-apprunner-only
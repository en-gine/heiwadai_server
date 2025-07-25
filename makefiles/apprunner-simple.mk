# シンプルなApp Runner構成（VPC/NAT Gateway不要）
AWS_REGION ?= ap-northeast-1
AWS_PROFILE ?= default
AWS_CMD = aws --profile $(AWS_PROFILE)

# App Runner設定
APP_RUNNER_SERVICE_NAME ?= heiwadai-server
ECR_IMAGE_URI ?= $(AWS_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/heiwadai-server:latest
APP_RUNNER_PORT ?= 3000
APP_RUNNER_CPU ?= 0.25 vCPU
APP_RUNNER_MEMORY ?= 0.5 GB

# シンプルなApp Runnerサービス作成（VPCなし）
create-simple-apprunner:
ifndef AWS_ID
	$(error AWS_ID is not set)
endif
	@echo "Creating simple App Runner service without VPC..."
	@echo '{ \
		"ServiceName": "$(APP_RUNNER_SERVICE_NAME)", \
		"SourceConfiguration": { \
			"AuthenticationConfiguration": { \
				"AccessRoleArn": "arn:aws:iam::$(AWS_ID):role/service-role/AppRunnerECRAccessRole" \
			}, \
			"AutoDeploymentsEnabled": true, \
			"ImageRepository": { \
				"ImageIdentifier": "$(ECR_IMAGE_URI)", \
				"ImageConfiguration": { \
					"Port": "$(APP_RUNNER_PORT)" \
				}, \
				"ImageRepositoryType": "ECR" \
			} \
		} \
	}' > /tmp/apprunner-simple.json
	$(AWS_CMD) apprunner create-service \
		--cli-input-json file:///tmp/apprunner-simple.json \
		--region $(AWS_REGION)
	@rm -f /tmp/apprunner-simple.json

# 固定IP Lambda関数のデプロイ
deploy-fixed-ip-lambda:
	@echo "Deploying fixed IP proxy Lambda..."
	cd lambda/fixed-ip-proxy && \
	npm init -y && \
	npm install axios && \
	zip -r function.zip . && \
	$(AWS_CMD) lambda create-function \
		--function-name heiwadai-fixed-ip-proxy \
		--runtime nodejs18.x \
		--role arn:aws:iam::$(AWS_ID):role/lambda-vpc-execution-role \
		--handler index.handler \
		--zip-file fileb://function.zip \
		--timeout 60 \
		--vpc-config SubnetIds=subnet-xxx,SecurityGroupIds=sg-xxx \
		--region $(AWS_REGION) || \
	$(AWS_CMD) lambda update-function-code \
		--function-name heiwadai-fixed-ip-proxy \
		--zip-file fileb://function.zip \
		--region $(AWS_REGION)

# EC2 NATインスタンス作成オプション
create-nat-instance:
	@echo "Creating EC2 NAT instance (cheap alternative)..."
	@echo "1. Launch t4g.nano instance"
	@echo "2. Allocate Elastic IP"
	@echo "3. Configure as NAT"
	@echo "Monthly cost: ~$3-5"
	@echo ""
	@echo "Commands:"
	@echo "aws ec2 run-instances --instance-type t4g.nano --image-id ami-xxx"
	@echo "aws ec2 allocate-address --domain vpc"
	@echo "aws ec2 associate-address --instance-id i-xxx --allocation-id eipalloc-xxx"

# コスト比較
show-cost-comparison:
	@echo "=== Fixed IP Solutions Cost Comparison ==="
	@echo ""
	@echo "1. NAT Gateway:"
	@echo "   - Monthly: ~$45 + data transfer"
	@echo "   - Pros: Managed, highly available"
	@echo "   - Cons: Expensive"
	@echo ""
	@echo "2. Lambda with VPC + NAT Gateway:"
	@echo "   - Monthly: ~$45 (same as above)"
	@echo "   - Pros: Only for specific calls"
	@echo "   - Cons: Still need NAT Gateway"
	@echo ""
	@echo "3. EC2 NAT Instance:"
	@echo "   - Monthly: ~$3-5 (t4g.nano)"
	@echo "   - Pros: Very cheap"
	@echo "   - Cons: Manual management, single point of failure"
	@echo ""
	@echo "4. App Runner Direct (No fixed IP):"
	@echo "   - Monthly: $0 extra"
	@echo "   - Pros: Simplest, cheapest"
	@echo "   - Cons: Dynamic IP"
	@echo ""
	@echo "5. API Gateway + Lambda:"
	@echo "   - Monthly: ~$3.50/million requests"
	@echo "   - Pros: Pay per use"
	@echo "   - Cons: Cold starts"

.PHONY: create-simple-apprunner deploy-fixed-ip-lambda create-nat-instance show-cost-comparison
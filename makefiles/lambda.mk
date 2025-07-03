# AWS Lambda関連のコマンド
AWS_REGION ?= ap-northeast-1
AWS_PROFILE ?= default
LAMBDA_FUNCTION_NAME ?= heiwadai-birthday-coupon-issuer
LAMBDA_ROLE_NAME ?= heiwadai-lambda-execution-role
EVENTBRIDGE_RULE_NAME ?= heiwadai-birthday-coupon-schedule
LAMBDA_DIR = lambda/birthday-coupon

# AWS CLI profile設定
AWS_CMD = aws --profile $(AWS_PROFILE)

# AWS認証チェック
aws-check:
	@echo "=== AWS Authentication Check ==="
	@echo "Profile: $(AWS_PROFILE)"
	@echo "Region: $(AWS_REGION)"
	@echo ""
	@echo "Testing credentials..."
	@$(AWS_CMD) sts get-caller-identity 2>&1 || ( \
		echo "❌ Authentication failed!" && \
		echo "" && \
		echo "Try one of these solutions:" && \
		echo "1. SSO Login: aws sso login --profile $(AWS_PROFILE)" && \
		echo "2. Configure profile: aws configure --profile $(AWS_PROFILE)" && \
		echo "3. Use AssumeRole: make aws-assume-role" && \
		false \
	)

# SSO ログイン
aws-sso-login:
	@echo "Starting SSO login for profile: $(AWS_PROFILE)"
	aws sso login --profile $(AWS_PROFILE)
	@echo ""
	@echo "Testing authentication..."
	@$(MAKE) aws-check

include ./lambda/birthday-coupon/.env

# Lambda関数のデプロイ
deploy-lambda: lambda-package lambda-create lambda-update-env

# Lambda関数のパッケージング
lambda-package:
	@echo "Packaging Lambda function..."
	cd $(LAMBDA_DIR) && \
	npm install --production && \
	zip -r function.zip . -x "*.env*" "README*"

# Lambda実行ロールの作成
lambda-create-role:
	@echo "Creating Lambda execution role..."
	$(AWS_CMD) iam create-role \
		--role-name $(LAMBDA_ROLE_NAME) \
		--assume-role-policy-document '{ \
			"Version": "2012-10-17", \
			"Statement": [ \
				{ \
					"Effect": "Allow", \
					"Principal": { "Service": "lambda.amazonaws.com" }, \
					"Action": "sts:AssumeRole" \
				} \
			] \
		}' \
		--region $(AWS_REGION) || true
	@echo "Attaching basic execution policy..."
	$(AWS_CMD) iam attach-role-policy \
		--role-name $(LAMBDA_ROLE_NAME) \
		--policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole \
		--region $(AWS_REGION) || true

# Lambda関数の作成
lambda-create:
	@echo "Creating Lambda function..."
	$(eval ACCOUNT_ID := $(shell $(AWS_CMD) sts get-caller-identity --query Account --output text))
	$(AWS_CMD) lambda create-function \
		--function-name $(LAMBDA_FUNCTION_NAME) \
		--runtime nodejs18.x \
		--role arn:aws:iam::$(ACCOUNT_ID):role/$(LAMBDA_ROLE_NAME) \
		--handler index.handler \
		--zip-file fileb://$(LAMBDA_DIR)/function.zip \
		--timeout 60 \
		--memory-size 128 \
		--region $(AWS_REGION) || \
	$(AWS_CMD) lambda update-function-code \
		--function-name $(LAMBDA_FUNCTION_NAME) \
		--zip-file fileb://$(LAMBDA_DIR)/function.zip \
		--region $(AWS_REGION)

# Lambda環境変数の更新
lambda-update-env:
ifndef CRON_ACCESS_ENDPOINT
	$(error CRON_ACCESS_ENDPOINT is not set. Please set environment variables)
endif
ifndef CRON_ACCESS_SECRET
	$(error CRON_ACCESS_SECRET is not set. Please set environment variables)
endif
ifndef CRON_ACCESS_KEY
	$(error CRON_ACCESS_KEY is not set. Please set environment variables)
endif
	@echo "Updating Lambda environment variables..."
	@echo '{"Variables":{"CRON_ACCESS_ENDPOINT":"$(CRON_ACCESS_ENDPOINT)","CRON_ACCESS_SECRET":"$(CRON_ACCESS_SECRET)","CRON_ACCESS_KEY":"$(CRON_ACCESS_KEY)"}}' > /tmp/lambda-env.json
	$(AWS_CMD) lambda update-function-configuration \
		--function-name $(LAMBDA_FUNCTION_NAME) \
		--environment file:///tmp/lambda-env.json \
		--region $(AWS_REGION)
	@rm -f /tmp/lambda-env.json

# EventBridgeルールの作成（毎月1日午前9時JST = 午前0時UTC）
eventbridge-create-rule: aws-check
	@echo "Creating EventBridge rule..."
	$(AWS_CMD) events put-rule \
		--name $(EVENTBRIDGE_RULE_NAME) \
		--schedule-expression "cron(0 0 1 * ? *)" \
		--description "Monthly birthday coupon issuance on 1st day of every month" \
		--region $(AWS_REGION)

# EventBridgeにLambda権限を付与
eventbridge-add-permission:
	@echo "Adding EventBridge permission to Lambda..."
	$(eval ACCOUNT_ID := $(shell $(AWS_CMD) sts get-caller-identity --query Account --output text))
	$(AWS_CMD) lambda add-permission \
		--function-name $(LAMBDA_FUNCTION_NAME) \
		--statement-id birthday-coupon-eventbridge \
		--action lambda:InvokeFunction \
		--principal events.amazonaws.com \
		--source-arn arn:aws:events:$(AWS_REGION):$(ACCOUNT_ID):rule/$(EVENTBRIDGE_RULE_NAME) \
		--region $(AWS_REGION) || true

# EventBridgeターゲットの設定
eventbridge-add-target:
	@echo "Adding Lambda target to EventBridge rule..."
	$(eval ACCOUNT_ID := $(shell $(AWS_CMD) sts get-caller-identity --query Account --output text))
	$(AWS_CMD) events put-targets \
		--rule $(EVENTBRIDGE_RULE_NAME) \
		--targets "Id"="1","Arn"="arn:aws:lambda:$(AWS_REGION):$(ACCOUNT_ID):function:$(LAMBDA_FUNCTION_NAME)" \
		--region $(AWS_REGION)

# EventBridge完全セットアップ
deploy-eventbridge: eventbridge-create-rule eventbridge-add-permission eventbridge-add-target

# 完全デプロイ（初回用）
deploy-birthday-coupon: lambda-create-role deploy-lambda deploy-eventbridge
	@echo "Birthday coupon Lambda and EventBridge deployed successfully!"
	@echo "Schedule: Every 1st day of the month at 00:00 UTC (09:00 JST)"

# Lambda関数の手動実行（テスト用）
lambda-test:
	@echo "Testing Lambda function manually..."
	$(AWS_CMD) lambda invoke \
		--function-name $(LAMBDA_FUNCTION_NAME) \
		--region $(AWS_REGION) \
		--payload '{}' \
		response.json
	@cat response.json && rm response.json

# Lambda関数の削除
lambda-delete:
	@echo "Deleting Lambda function..."
	$(AWS_CMD) lambda delete-function \
		--function-name $(LAMBDA_FUNCTION_NAME) \
		--region $(AWS_REGION) || true

# EventBridgeルールの削除
eventbridge-delete:
	@echo "Removing targets from EventBridge rule..."
	$(AWS_CMD) events remove-targets \
		--rule $(EVENTBRIDGE_RULE_NAME) \
		--ids "1" \
		--region $(AWS_REGION) || true
	@echo "Deleting EventBridge rule..."
	$(AWS_CMD) events delete-rule \
		--name $(EVENTBRIDGE_RULE_NAME) \
		--region $(AWS_REGION) || true

# 完全削除
delete-birthday-coupon: lambda-delete eventbridge-delete
	@echo "Birthday coupon resources deleted successfully!"

# ヘルプ
help-lambda:
	@echo "Birthday Coupon Lambda Commands:"
	@echo "  deploy-birthday-coupon  - 初回デプロイ（Lambda + EventBridge）"
	@echo "  deploy-lambda          - Lambda関数のみデプロイ"
	@echo "  deploy-eventbridge     - EventBridgeのみデプロイ"
	@echo "  lambda-test           - Lambda関数をテスト実行"
	@echo "  delete-birthday-coupon - 全リソースを削除"
	@echo ""
	@echo "AWS Authentication:"
	@echo "  aws-check             - 認証状態を確認"
	@echo "  aws-sso-login         - AWS SSOでログイン"
	@echo ""
	@echo "Environment variables required:"
	@echo "  CRON_ACCESS_ENDPOINT   - APIエンドポイントURL"
	@echo "  CRON_ACCESS_SECRET     - Authorization token"
	@echo "  CRON_ACCESS_KEY       - X-Cron-Key value"
	@echo ""
	@echo "Optional variables:"
	@echo "  AWS_REGION            - AWS region (default: ap-northeast-1)"
	@echo "  AWS_PROFILE           - AWS CLI profile (default: default)"
	@echo "  LAMBDA_FUNCTION_NAME  - Lambda function name"

.PHONY: deploy-lambda lambda-package lambda-create-role lambda-create lambda-update-env \
        eventbridge-create-rule eventbridge-add-permission eventbridge-add-target \
        deploy-eventbridge deploy-birthday-coupon lambda-test lambda-delete \
        eventbridge-delete delete-birthday-coupon help-lambda \
        aws-check aws-sso-login
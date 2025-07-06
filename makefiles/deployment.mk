# デプロイメント関連のコマンド
AWS_PROFILE ?= heiwadai
AWS_CMD = aws --profile $(AWS_PROFILE)

# ECR関連
login:
	$(AWS_CMD) ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin https://${AWS_ID}.dkr.ecr.ap-northeast-1.amazonaws.com

build:
	docker build -t ${AWS_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/heiwadai-server:latest -f ./docker/Dockerfile/server/Dockerfile.prod .

push: build
	docker push ${AWS_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/heiwadai-server:latest

# プロト関連
init-proto:
	git submodule add git@github.com:en-gine/heiwadai_proto.git server/v1
	
update-proto:
	git submodule update --remote

# デプロイフロー
deploy-server: login build push
	@echo "Server image pushed to ECR successfully!"

# 本番環境確認用
check-deployment:
	@echo "Checking deployment status..."
	$(AWS_CMD) ecs describe-services --cluster heiwadai-cluster --services heiwadai-service --region ap-northeast-1 || echo "ECS service not found"

.PHONY: login build push init-proto update-proto deploy-server check-deployment
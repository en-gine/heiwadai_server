# Heiwadai プロジェクト Makefile
# Makefile用環境変数を読み込む
include .env.makefile

# 分割したMakefileを読み込む
include makefiles/docker.mk
include makefiles/server.mk
include makefiles/deployment.mk
include makefiles/lambda.mk
include makefiles/apprunner.mk
include makefiles/iam.mk

# デフォルトターゲット
.DEFAULT_GOAL := help

# プロジェクト共通のヘルプ
help:
	@echo "Heiwadai Project Commands:"
	@echo ""
	@echo "Docker Operations:"
	@echo "  up                     - Start all containers"
	@echo "  down                   - Stop and remove containers"
	@echo "  restart                - Restart all containers"
	@echo "  bash                   - Access server container shell"
	@echo "  logs                   - Show container logs"
	@echo ""
	@echo "Development:"
	@echo "  dev                    - Start development server with hot reload"
	@echo "  run                    - Run server"
	@echo "  buf                    - Generate protobuf code"
	@echo "  lint                   - Run linter"
	@echo "  test                   - Run all tests"
	@echo ""
	@echo "Database:"
	@echo "  migrate-up             - Run database migrations"
	@echo "  migrate-down           - Rollback migrations"
	@echo "  init-db                - Initialize database"
	@echo "  sqlboiler              - Generate ORM models"
	@echo ""
	@echo "Deployment:"
	@echo "  deploy-server          - Build and push server image to ECR"
	@echo "  deploy-birthday-coupon - Deploy Lambda birthday coupon function"
	@echo "  deploy-apprunner       - Deploy to AWS App Runner (first time)"
	@echo "  deploy-apprunner-update - Update App Runner with new image"
	@echo ""
	@echo "For detailed help on specific categories:"
	@echo "  make help-lambda       - Lambda deployment help"
	@echo "  make help-apprunner    - App Runner deployment help"
	@echo "  make help-iam          - IAM management help"
	@echo ""
	@echo "Environment variables required for Lambda:"
	@echo "  CRON_ACCESS_ENDPOINT, CRON_ACCESS_SECRET, CRON_ACCESS_KEY"
	@echo "  AWS_PROFILE (optional, default: default)"

# 全体初期化コマンド
init: up init-db
	@echo "Project initialized successfully!"

# 開発環境セットアップ
setup-dev: up buf sqlboiler
	@echo "Development environment setup complete!"

.PHONY: help init setup-dev
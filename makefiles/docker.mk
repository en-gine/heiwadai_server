# Docker関連のコマンド

# 基本的なDocker操作
up:
	docker compose up -d

down:
	docker compose down

stop:
	docker compose stop

restart:
	docker compose restart

logs:
	docker compose logs -f

# サーバーコンテナに接続
bash:
	docker compose exec server bash

# サーバー実行関連
run: 
	docker compose exec server go run .

dev:
	docker compose -f docker-compose.dev.yml exec server air

# .envファイルを再読み込みして起動
reload-env:
	docker-compose --env-file .env up -d

# コンテナ状態確認
ps:
	docker compose ps

# 使用していないDockerリソースをクリーンアップ
docker-clean:
	docker system prune -f
	docker volume prune -f

.PHONY: up down stop restart logs bash run dev reload-env ps docker-clean
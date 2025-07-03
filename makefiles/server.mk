# サーバー関連のコマンド

# コード生成
buf: 
	docker compose exec server buf generate

# ORM関連
sqlboiler: 
	docker compose exec server make sqlboiler

# データベース関連
migrate-up:
	docker compose exec server make migrate-up

migrate-down:
	docker compose exec server make migrate-down

migrate-create:
	docker compose exec server make migrate-create TABLE_NAME=$(TABLE_NAME)

init-db:
	docker compose exec server make init-db

# テスト実行
test:
	docker compose exec server make test-all

test-birth-coupon:
	docker compose exec server make birthCouponTest

test-booking:
	docker compose exec server make bookTest

test-checkin:
	docker compose exec server make checkinTest

test-coupon:
	docker compose exec server make couponTest

test-mail:
	docker compose exec server make mailTest

# リンター
lint:
	docker compose exec server make lint

# シーダー実行
seed-store:
	docker compose exec server make seeder-store

seed-user:
	docker compose exec server make seeder-user

seed-admin:
	docker compose exec server make seeder-admin

seed-coupon:
	docker compose exec server make seeder-coupon

# 開発用コマンド実行
server-exec:
	docker compose exec server $(CMD)

.PHONY: buf sqlboiler migrate-up migrate-down migrate-create init-db test \
        test-birth-coupon test-booking test-checkin test-coupon test-mail \
        lint seed-store seed-user seed-admin seed-coupon server-exec
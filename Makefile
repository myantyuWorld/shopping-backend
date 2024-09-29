up-build:
	docker compose up -d --build
	
up:
	docker-compose up -d

down:
	docker-compose down

restart:
	@make down
	@make up

exec-api:
	docker-compose exec api /bin/bash

# psql -d eb -U postgres
# \z // テーブル一覧表示
# \d TABLE_NAME テーブル定義確認
exec-db:
	docker compose exec db /bin/bash 
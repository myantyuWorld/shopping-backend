up:
	docker-compose up -d

down:
	docker-compose down

restart:
	@make down
	@make up

exec-api:
	docker-compose exec api /bin/bash
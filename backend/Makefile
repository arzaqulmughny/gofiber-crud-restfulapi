include .env

DB_URL="mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

migrate-create:
	migrate create -ext sql -dir database/migrations -seq $(name)

migrate-up:
	migrate -path database/migrations -database $(DB_URL) up

migrate-down:
	migrate -path database/migrations -database $(DB_URL) down

migrate-force:
	migrate -path database/migrations -database $(DB_URL) force $(version)
include .env

migrate-up:
	migrate -path ./db/migrations -database "${DB_URL}" up
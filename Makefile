DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

create-db:
	docker exec -it simple-bank-postgres createdb --username=root --owner=root simple_bank

drop-db:
	docker exec -it simple-bank-postgres dropdb simple_bank

migrate-up:
	migrate -path=./db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate-down:
	migrate -path=./db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres create-db drop-db migrate-up migrate-down
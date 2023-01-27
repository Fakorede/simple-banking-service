postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine 

createdb:
	docker exec -it postgres12 createdb --username=postgres --owner=postgres simple_banking_service

dropdb:
	docker exec -it postgres12 dropdb simple_banking_service

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5433/simple_banking_service?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5433/simple_banking_service?sslmode=disable" -verbose down

sqlc:
	sqlc generate 

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
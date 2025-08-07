postgres:
	 docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=psql -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres bank

dropdb:
	docker exec -it postgres dropdb -U postgres bank

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:psql@172.19.34.230:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:psql@172.19.34.230:5432/bank?sslmode=disable" -verbose down

start:
	docker start postgres

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown start sqlc test
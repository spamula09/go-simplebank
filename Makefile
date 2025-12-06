postgres:
	docker run --name postgres18 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:18-alpine

createdb:
	docker exec -it postgres18  createdb  --username root   --owner=root simplebank

dropdb:
	docker exec -it postgres18  dropdb  simplebank

migrateup:
	migrate  -path db/migration -database  "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate  -path db/migration -database  "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down 

migrateup1:
	migrate  -path db/migration -database  "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up 1

migratedown1:
	migrate  -path db/migration -database  "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run  main.go

mock:
	mockgen  -package mockdb -destination db/mock/store.go  github.com/spamula09/go-simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
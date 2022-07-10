postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secrect -d postgres:14.2-alpine
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank
migrateup:
	migrate -path "./db/migration" -database "postgresql://root:secrect@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path "./db/migration" -database "postgresql://root:secrect@localhost:5432/simple_bank?sslmode=disable" -verbose down
dropdb:
	docker exec -it postgres14 dropdb simple_bank
# sqlc generate
sqlc:
	docker run --rm -v "C:\WebLearning\Simplebank":/src -w /src kjconroy/sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/Al3xDo/simple_bank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
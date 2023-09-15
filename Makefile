postgres:
	docker run --name ps_name -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb: 
	docker exec -it ps_name createdb --username=root --owner=root gopherbank

dropdb:
	docker exec -it ps_name dropdb gopherbank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gopherbank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gopherbank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

# to run unit test in complete project
test:
	go test -v -cover ./...

server:
	go run main.go

# .PHONY is a way to ensure that Make knows when a target is not meant to produce a file but 
# rather execute specific commands or recipes.
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
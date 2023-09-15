package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/scortier/gopherbank/api"
	db "github.com/scortier/gopherbank/db/sqlc"
)

const (
	// DBDriverPostgres is the driver name for postgres
	DBDriverPostgres = "postgres"
	// DBSourcePostgres is the data source name for postgres
	DBSourcePostgres = "postgresql://root:secret@localhost:5432/gopherbank?sslmode=disable"
)

func main() {
	// create new db connection
	conn, err := sql.Open(DBDriverPostgres, DBSourcePostgres)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)     // create new store
	server := api.NewServer(store) // create new server

	err = server.Start(":8080") // start server
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}

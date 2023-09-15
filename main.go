package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/scortier/gopherbank/api"
	db "github.com/scortier/gopherbank/db/sqlc"
	"github.com/scortier/gopherbank/util"
)

func main() {
	config, err := util.LoadConfig(".") // load config from file
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	// create new db connection
	conn, err := sql.Open(config.DBDriverPostgres, config.DBSourcePostgres)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)     // create new store
	server := api.NewServer(store) // create new server

	err = server.Start(config.ServerAddress) // start server
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}

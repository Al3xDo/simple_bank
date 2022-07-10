package main

import (
	"database/sql"
	"log"

	"github.com/Al3xDo/simple_bank/api"
	db "github.com/Al3xDo/simple_bank/db/sqlc"
	"github.com/Al3xDo/simple_bank/util"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:secrect@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8000"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("you cannot start server", err)
	}
}

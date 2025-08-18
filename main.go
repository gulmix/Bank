package main

import (
	"context"
	"log"

	"github.com/gulmix/bank/api"
	db "github.com/gulmix/bank/db/sqlc"
	"github.com/gulmix/bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerName)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

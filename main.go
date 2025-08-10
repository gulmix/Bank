package main

import (
	"context"
	"log"

	"github.com/gulmix/bank/api"
	db "github.com/gulmix/bank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	connString   = "postgresql://postgres:psql@localhost:5432/bank?sslmode=disable"
	serverAddres = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddres)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

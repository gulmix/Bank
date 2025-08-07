package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const connString = "postgresql://postgres:psql@172.19.34.230:5432/bank?sslmode=disable"

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}

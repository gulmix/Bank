package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gulmix/bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}
	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}

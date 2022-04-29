package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbString = "postgresql://postgres:pass@localhost:5434/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbString)
	if err != nil {
		log.Fatalf("cannot open db: %v, as the error occured: %v", dbString, err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

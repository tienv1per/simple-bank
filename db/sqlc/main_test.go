package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"simple-bank/util"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error when connecting to DB:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}

package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/pjserol/api-rest-user/common/logs"

	// lib/pq
	_ "github.com/lib/pq"
	"github.com/pjserol/api-rest-user/config"
)

// DB connection
var DB *sql.DB

// Connected connected or not
var Connected bool

///////////////////
// Connection DB
///////////////////

// Initialise initialise the database connection.
func Initialise(ctx context.Context) {
	env := config.InitEnvironment()

	appEnv := env.AppEnvironment
	awsRegion := env.AWSRegion
	dbEndpoint := env.DBEndpoint
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbName := env.DBName
	testDatabaseURL := env.TestDatabaseURL

	if env.IsTestMode {
		// display the line and file in the logs
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	newDB, err := connectToPostgreSQL(ctx, appEnv, awsRegion, dbEndpoint, dbUser, dbName, dbPort, testDatabaseURL)
	if err != nil {
		log.Fatalf("Couldn't initialise the DB")
	}

	logs.Log(ctx, "Database started!")

	Connected = true
	DB = newDB.DB
}

func connectToPostgreSQL(ctx context.Context, appEnv, awsRegion, dbEndpoint, dbUser, dbName string, dbPort int, testDatabaseURL string) (ServiceDB, error) {
	if appEnv == "local" {

		if testDatabaseURL == "" {
			logs.Log(ctx, "testDatabaseURL should not be empty")
		}

		conn, err := sql.Open("postgres", testDatabaseURL)
		if err != nil {
			log.Println(err)
			return ServiceDB{}, err
		}
		return ServiceDB{
			DB: conn,
		}, nil
	}

	// TODO: connection to an AWS DB

	log.Fatalf("AWS DB not implemented")

	return ServiceDB{}, nil
}

# REST API

## Environment Variables

* APP_ENVIRONMENT
* IS_TEST_MODE
* SUPPRESS_LOGS
* PORT_NUMBER

* TEST_DATABASE_URL
* AWS_REGION
* USER_DATABASE_ENDPOINT
* USER_DATABASE_USER
* USER_DATABASE_PORT
* USER_DATABASE_NAME

## Environment Variables to test locally

* Assuming default homebrew postgres installation:
* APP_ENVIRONMENT=local
* TEST_DATABASE_URL="host=localhost port=54320 user=postgres password=admin dbname=postgres sslmode=disable"
* SUPPRESS_LOGS = "true"

## Run the application

Run postgres locally.

* ./scripts/script.sh

Init environment variable related to the correct environment (for the moment only local).

* APP_ENVIRONMENT=local
* TEST_DATABASE_URL="host=localhost port=54320 user=postgres password=admin dbname=postgres sslmode=disable"

Run the application.

* go run main.go

## Test locally

Run postgres locally.

* ./scripts/script.sh


Run integration and unit test.

* ./scripts/test.sh

Integration test only.

* go test -tags integration

Unit test only.

* go test -v 

## postgres

Execute command into mysql

* docker exec -it my_postgres bash
* psql -U postgres
* CREATE DATABASE appdb;
* postgres-# \q

## Tools

* Postman
* DBVisualizer

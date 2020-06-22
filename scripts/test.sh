#!/bin/bash

export ENVIRONMENT=local
export TEST_DATABASE_URL="host=localhost port=54320 user=postgres password=admin dbname=postgres sslmode=disable"
export SUPPRESS_LOGS="false"

cd ..

go test -tags=integration

go test ./...
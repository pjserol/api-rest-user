#!/bin/bash

echo "Removing docker resources..."
docker rm -f my_postgres

echo "Creating postgres container..."
docker run -d --name my_postgres -v my_dbdata:/var/lib/postgresql/data -e POSTGRES_PASSWORD=admin -p 54320:5432 postgres:11
# be careful we don't use the default port of postgres 5432 -> 54210

# docker cp ./localfile.sql containername:/container/path/file.sql
docker cp ./dump.sql my_postgres:/docker-entrypoint-initdb.d/dump.sql

# docker exec -u postgresuser containername psql dbname postgresuser -f /container/path/file.sql
docker exec -u postgres my_postgres psql postgres postgres -f docker-entrypoint-initdb.d/dump.sql
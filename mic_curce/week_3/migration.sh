#!/bin/bash
source .env


export MIGRATION_DSN="host=pg port=5432 dbname=$PG_DATABASE_NAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"

echo "Waiting 2 seconds for database to be ready..."
sleep 2
echo "Starting migrations..."
goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v
#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE DATABASE gin_gorm;
	GRANT ALL PRIVILEGES ON DATABASE gin_gorm TO postgres;
EOSQL
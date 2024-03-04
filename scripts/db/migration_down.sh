#!/bin/bash
BASEDIR=$(dirname $0)
migrate -path=${BASEDIR}../../../_db/postgreSQL/migrations/ -database=$DB_URL down
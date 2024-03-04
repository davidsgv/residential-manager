if [ "$1" = "" ]
then
    echo "Usage: $0 <migration name>"
    exit
fi

MIGRATE_DIR=$(dirname $0)../../../_db/postgreSQL/migrations/
migrate create -seq -ext=.sql -dir=${MIGRATE_DIR} $1
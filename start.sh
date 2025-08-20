#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$CONN_STRING" -verbose up

echo "start the app"
exec "$@"
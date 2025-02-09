#!/bin/sh
set -e

echo "Waiting for database to be ready..."
until nc -z -v -w30 "$DB_HOST" "$DB_PORT"; do
  echo "Waiting for database..."
  sleep 5
done

echo "Database is up!"
exec "$@"

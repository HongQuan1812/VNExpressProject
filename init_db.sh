#!/bin/bash

set -e

# Start MySQL server
docker-entrypoint.sh mysqld &

# Wait for MySQL server to start (max 30 seconds)
timeout=30
while ! mysqladmin ping -h"localhost" --silent; do
  timeout=$((timeout - 1))
  if [ $timeout -le 0]; then
    echo "MySQL server startup timed out"
    exit 1
  fi
  sleep 1
done

# Restore the dump file
if [ -f /docker-entrypoint-initdb.d/dump.sql ]; then
  mysql -uroot -p"$MYSQL_ROOT_PASSWORD" < /docker-entrypoint-initdb.d/dump.sql
fi

# Wait for the MySQL server process to exit
wait

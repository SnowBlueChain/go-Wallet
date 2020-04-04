#!/bin/sh
# wait_mysql.sh

#command: ["./wait_mysql.sh", "mysql", "/go/bin/api", "--host", "0.0.0.0", "--port", "8081", "--config", "config/example.yml"]
set -e

host="$1"
shift
cmd="$@"

count=0
while ! mysqladmin ping -h"$host" --silent; do
    >&2 echo "Database is unavailable - sleeping"
    count=$((++count))
    if [ "$count" -gt 30 ]; then
        >&2 echo "timeout"
        exit 1
    fi
    sleep 1
done

>&2 echo "Database is up - executing command"
exec $cmd

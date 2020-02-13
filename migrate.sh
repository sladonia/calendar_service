#!/bin/sh

#export POSTGRESQL_URL='postgres://user:password@localhost:5432/calendar_development?sslmode=disable'

print_usage()
{
  echo "usage":
  echo "./migrate.sh (up | down) [N]"
}


if [ "$POSTGRESQL_URL" = "" ]
then
  echo "POSTGRESQL_URL env var should be provided"
  exit
fi

MIGRATION_COMMAND="$1"
if [ "$MIGRATION_COMMAND" = "" ]
then
  print_usage
  exit
fi

migrate -database ${POSTGRESQL_URL} -path migrations $MIGRATION_COMMAND "$2"

#!/bin/sh

print_usage()
{
  echo "usage":
  echo "./create_migrations.sh migration_name"
}

create_migration()
{
  migrate create -ext sql -dir migrations -seq $MIGRATION_NMAE
}

MIGRATION_NMAE="$1"
if [ "$MIGRATION_NMAE" = "" ]
then
  print_usage
  exit
fi

create_migration

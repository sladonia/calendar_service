# calendar_service

A simple http REST API service for creating and managing user calendars and events. Made in self-education purpose.

### build dependencies

* go v1.13
* GNU Make
* docker v19.03.3
* docker-compose v1.25.0

### service dependencies

Uses postgres v10 as a primary data store

uuid-ossp extension should be installed

calendar service postgres configuration env vars and the defaults:
```.env
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=bookshelf_db
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_MAX_OPEN_CONNECTIONS=25
POSTGRES_MAX_IDLE_CONNECTIONS=25
POSTGRES_CONNECTION_MAX_LIFETIME=5
```

### db migrations

use github.com/golang-migrate/migrate V4.8.0 for migrations management
* Export POSTGRESQL_URL env var with postgres connection string. Example:
```sh
export POSTGRESQL_URL='postgres://user:password@localhost:5432/calendar_development?sslmode=disable'
``` 

* Run migrations:
```sh
./migrate.sh up
``` 

* to generate empty migration files run:
```sh
./create_migrations.sh migration_name
```

### build

* build executable
```sh
make build
```

* build docker container
```sh
make docker_build
```

### run locally with docker-compose

* build services
```sh
make build
docker-compose build postgres calendar
```

* run migrations
```shell script
export POSTGRESQL_URL='postgres://user:password@localhost:5432/calendar_development?sslmode=disable'
./migrate.sh up
```

* run services
```sh
docker-compose up calendar
```

* check service
```sh
curl localhost:8080
# expected response:
{"message":"welcome to calendar api"}
```

### tests
test are executed against the test db. Ensure running postgres instance and the calendar_test db existence. Tests use .env.test file for service configuration.
run tests:
```sh
make test
```

### configuration

service is configurable with env vars and the defaults are

```.env
SERVICE_NAME=calendar
ENV=dev
LOG_LEVEL=debug
PORT=:8080

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=bookshelf_db
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_MAX_OPEN_CONNECTIONS=25
POSTGRES_MAX_IDLE_CONNECTIONS=25
POSTGRES_CONNECTION_MAX_LIFETIME=5
```

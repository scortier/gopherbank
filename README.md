# GopherBank

- Create and manage account

  - Owner, balance, currency

- Record all balance changes

  - Create an account entry for each change

- Money transfer transaction
  - Perform money transfer between 2 accounts consistently within a transaction.

# Database Design

- Design a SQL DB schema using dbdiagram.io
- Save it in PNG/PDF diagram
- Generate code to create a schema in a target db engine: Postgres/MySQL/SQL Server

# Install Docker & Postgres + TablePlus

- install docker desktop
- find official docker postgress image
-

# DB Migration

- golang-migrate library to write and read db migration : https://github.com/golang-migrate/migrate
- Up and down script for migration schema , why ?
- Up for any change in current migration
- Down for revert changes in migration

```
╭─scortier@Infinity ~
╰─$ docker stop ps_name                                                           127 ↵
ps_name
╭─scortier@Infinity ~
╰─$ docker start ps_name
ps_name
╭─scortier@Infinity ~
╰─$ docker ps
CONTAINER ID   IMAGE             COMMAND                  CREATED          STATUS         PORTS                    NAMES
2d1f2ae8b71f   postgres:latest   "docker-entrypoint.s…"   57 minutes ago   Up 5 seconds   0.0.0.0:5432->5432/tcp   ps_name
╭─scortier@Infinity ~
╰─$ docker exec -it ps_name /bin/sh
# psql gopherbank
psql: error: connection to server on socket "/var/run/postgresql/.s.PGSQL.5432" failed: FATAL:  database "gopherbank" does not exist
# createdb --username=root --owner=root gopherbank
# psql gopherbank
psql (15.4 (Debian 15.4-1.pgdg120+1))
Type "help" for help.

gopherbank=# \q
# dropdb gopherbank
# exit
╭─scortier@Infinity ~
╰─$ docker exec -it ps_name createdb --username=root --owner=root gopherbank
╭─scortier@Infinity ~
╰─$ docker exec -it ps_name psql -U root  gopherbank
psql (15.4 (Debian 15.4-1.pgdg120+1))
Type "help" for help.

gopherbank=# \q
```

# Generate Golang code from SQLC

- create sqlc.yaml
- use this : https://docs.sqlc.dev/en/v1.21.0/tutorials/getting-started-postgresql.html

```
version: "2"
sql:
  - engine: "postgresql"
    queries: "./db/query/"
    schema: "./db/migration/"
    gen:
      go:
        package: "db"
        out: "./db/sqlc"

```
- use create, get, update queries in query.sql and `make sqlc` to create corresponding code files.

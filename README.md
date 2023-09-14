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

# DB Transaction

- Each Query can do only 1 operation on 1 specific table
- storing query inside store is called composition, it is preferred way to extend struct functionality in go instead of inheritance.
- All individual function provided by queries will be available to store and we can support tx by adding more func to that new struct(Store)
- Issue : after one trans, fromAcc balance was not updated immedialtely which was leading issue in simultaneous tx,
  Sol: Add `For UPDATE` clause at the end of GetList

```
-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;
```

- After above changes, that problem is resolved but now deadlock issue will arrive
  ![For debugging deadlock process](image-1.png)
- so the deadlock happends from `SELECT statement`: https://wiki.postgresql.org/wiki/Lock_Monitoring
  ![Alt text](image.png)
- ![Alt text](image-2.png)

- the only conn between account and transfer schema is foreign key
  - ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
  - from_acc_id and to_acc_id of transfer table are referring to the id column of accounts table. So any affect on account id will ffect the foreign key constraint.
  - That'swhy when we select an account for update. it needs to acquire a lock to prevent conflict and ensure onsistency in the data.
  - To check again just run , create entry for acc 1 and 2 and select acc1 for update in tx1, you will geta deadlock,
    ![Alt text](image-3.png)  
    becoz both tx1 and tx2 has to wait for each other.
- HOW TO FIX IT ?

  - Remove foreign key
  - update db schema , just by not changing id while updating. Tell postgres to udpate acc but u will not touch its primary key. hence no tx lock , so no deadlock.

```
  -- name: GetAccountForUpdate :one
  SELECT * FROM accounts
  WHERE id = $1 LIMIT 1
  FOR UPDATE;

- Below NO KEY is added in last line.
-- name: GetAccountForUpdate :one
SELECT \* FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

```

- To avoid duplication of code of fetchign and updating balance in both accounts

```
		fromAcc, err := q.GetAccountForUpdate(ctx, arg.FromAccountID) // get account
		if err != nil {
			return err
		}

		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: fromAcc.Balance - arg.Amount,
		}) // update account
		if err != nil {
			return err
		}
```

- we can simplify this by adding the update query of account

```
FROM:
-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING *;

TO:
-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

// SET balance = balance + $2 is also fine. but we want to update struct field to new value name , here it will be amount that's why added sqlc.arg(amount)
```

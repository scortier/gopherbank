CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_email_verified" bool NOT NULL DEFAULT false,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- EACH OWNER HAS ATMOST 1 ACC FOR A SPECIFIC CURRENCY.
-- CHOOSE AMONG BELOW 2 OPTIONS : 
-- 1. Using direct unique index
-- 2. Add a unique constraint for the pair of owner and currency on the accounts table. It will be very similar 
-- to the command to add foreign key constraint above. In Summary, addign this unique constriant will automatically
-- create the same unique composite index for owner and currency.

-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");
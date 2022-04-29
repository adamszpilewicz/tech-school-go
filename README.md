# dbdiagrams
```
//// -- LEVEL 1
//// -- Tables and References

// Creating tables
Table accounts as A {
id bigserial [pk] // auto-increment
owner varchar [not null]
balance bigint [not null]
currency varchar [not null]
created_at timestamptz [not null, default: `now()`]

Indexes{
owner
}
}

Table countries {
id bigserial [pk]
account_id bigint [ref: > A.id]
amount bigint [not null, note: 'can be negtive or positive']

Indexes {
account_id
}

}

Table transfers {
id bigint [pk]
from_account_id bigint [ref: > A.id]
to_account_id bigint [ref: > A.id]
amount bigint [not null, note: 'must be positive']
created_at timestamptz [not null, default: `now()`]

Indexes {
from_account_id
to_account_id
(from_account_id, to_account_id)
}
}
```

```sql
CREATE TABLE "accounts" (
                            "id" bigserial PRIMARY KEY,
                            "owner" varchar NOT NULL,
                            "balance" bigint NOT NULL,
                            "currency" varchar NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "countries" (
                             "id" bigserial PRIMARY KEY,
                             "account_id" bigint NOT NULL,
                             "amount" bigint NOT NULL
);

CREATE TABLE "transfers" (
                             "id" bigint PRIMARY KEY,
                             "from_account_id" bigint NOT NULL,
                             "to_account_id" bigint NOT NULL,
                             "amount" bigint NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "countries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "countries"."amount" IS 'can be negtive or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "countries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");


```
 
CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "password" varchar NOT NULL,
  "is_admin" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "details" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "category" varchar NOT NULL,
  "cost" bigint NOT NULL,
  "date" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "details" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

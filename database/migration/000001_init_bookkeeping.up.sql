CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "is_admin" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categorys" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "details" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  "cost" bigint NOT NULL,
  "date" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "details" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "details" ADD FOREIGN KEY ("category_id") REFERENCES "categorys" ("id");

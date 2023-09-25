CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "is_admin" bool,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "categorys" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "details" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "category_id" bigint,
  "cost" bigint NOT NULL,
  "date" timestamptz NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

ALTER TABLE "details" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "details" ADD FOREIGN KEY ("category_id") REFERENCES "categorys" ("id");

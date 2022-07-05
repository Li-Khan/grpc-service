CREATE TABLE IF NOT EXISTS "event"(
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL UNIQUE,
    "date" TIMESTAMP NOT NULL UNIQUE
);
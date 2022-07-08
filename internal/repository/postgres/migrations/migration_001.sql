CREATE TABLE IF NOT EXISTS "event"(
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL UNIQUE,
    "date" TIMESTAMP NOT NULL UNIQUE
);

INSERT INTO "event" (
    "name",
    "date"
) VALUES ('Thor: Love and Thunder', '2022-07-07');

INSERT INTO "event" (
    "name",
    "date"
) VALUES ('Nirvana concert', '1994-02-12');

INSERT INTO "event" (
    "name",
    "date"
) VALUES ('Birthday', '2023-04-30');

INSERT INTO "event" (
    "name",
    "date"
) VALUES ('GopherCon', '2022-10-06');

INSERT INTO "event" (
    "name",
    "date"
) VALUES ('Day of the Capital', '2022-07-06');
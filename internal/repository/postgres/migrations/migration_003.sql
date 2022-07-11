ALTER TABLE "event"
    ADD COLUMN IF NOT EXISTS test TEXT;

INSERT INTO "event" (
    "name",
    "date",
    "test"
) VALUES ('Day of the Capital', '2022-07-06', 'test alter table');
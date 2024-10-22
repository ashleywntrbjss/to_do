CREATE SCHEMA todo

CREATE TABLE "todo".todo_items (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    is_complete BOOLEAN NOT NULL DEFAULT FALSE
);
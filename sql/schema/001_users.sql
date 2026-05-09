-- +goose Up
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    created_at DATE NOT NULL,
    updated_at DATE NOT NULL
    name TEXT UNIQUE NOT NULL,
);

-- -goose Down
DROP TABLE users;
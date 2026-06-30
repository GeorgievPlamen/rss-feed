-- +goose Up
CREATE TABLE
    IF NOT EXISTS users (
        id UUID NOT NULL Unique,
        created_at TIMESTAMP,
        updated_at TIMESTAMP,
        name text Unique NOT NULL
    );

-- +goose Down
DROP TABLE users;
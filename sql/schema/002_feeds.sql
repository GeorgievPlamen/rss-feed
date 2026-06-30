-- +goose Up
CREATE TABLE
    IF NOT EXISTS feeds (
        id UUID PRIMARY KEY,
        created_at TIMESTAMP,
        updated_at TIMESTAMP,
        name TEXT,
        url TEXT Unique,
        user_id UUID  NOT NULL references users(id) ON DELETE CASCADE
    );

-- +goose Down
DROP TABLE feeds;
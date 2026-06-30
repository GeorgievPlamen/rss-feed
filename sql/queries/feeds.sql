-- name: AddFeed :one
INSERT INTO
    feeds (id, created_at, updated_at, name, url, user_id)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetAllFeedsWithUserName :many
SELECT
    feeds.id,
    feeds.created_at,
    feeds.updated_at,
    feeds.name,
    feeds.url,
    feeds.user_id,
    users.name AS user_name
FROM feeds
LEFT JOIN users
    ON feeds.user_id = users.id;

-- name: GetFeedByUrl :one
SELECT *
FROM feeds
WHERE url = ($1);

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = ($1), updated_at = ($1)
WHERE id = ($2);

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
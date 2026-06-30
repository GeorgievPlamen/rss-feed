-- name: CreateFeedFollow :one
WITH cte AS (
    INSERT INTO
    feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT 
    cte.id,
    cte.created_at,
    cte.updated_at,
    cte.user_id,
    cte.feed_id,
    users.name AS user_name,
    feeds.name AS feed_name
FROM cte
INNER  JOIN users
    ON cte.user_id = users.id
INNER  JOIN feeds
    on cte.feed_id = feeds.id;

-- name: GetFeedFollowsByUserId :many
SELECT 
    feed_follows.id,
    feed_follows.created_at,
    feed_follows.updated_at,
    feeds.name AS feed_name
FROM feed_follows
LEFT JOIN feeds
    ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = ($1);

-- name: Unfollow :exec
DELETE
FROM feed_follows
WHERE 
    user_id = ($1) AND feed_id = ($2);
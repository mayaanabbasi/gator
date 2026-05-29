-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: FetchFeeds :many
SELECT * FROM feeds;

-- name: FetchFeedByUrl :one
SELECT * FROM feeds WHERE url=$1;

-- name: FetchFeedUsers :many
SELECT * FROM users INNER JOIN feeds ON users.id = feeds.user_id;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;

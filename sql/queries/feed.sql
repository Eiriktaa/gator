-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING  *;

-- name: GetFeeds :many
SELECT feeds.name as name, url, users.name as username
FROM feeds
INNER JOIN users ON users.id = feeds.user_id;

-- name: GetFeed :one
SELECT feeds.id as id, feeds.name as name, url, users.name as username
FROM feeds
INNER JOIN users ON users.id = feeds.user_id
WHERE url = $1
;

-- name: MarkFeedFetch :one
UPDATE feeds
    SET last_fetched_at = $1, updated_at = $1
    WHERE feeds.id = $2
RETURNING *
;

-- name: GetNextFeedToFetch :one
SELECT * 
FROM feeds
ORDER BY last_fetched_at 
ASC NULLS FIRST
LIMIT 1;


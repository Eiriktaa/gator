-- name: CreatePosts :one
INSERT INTO posts (id, created_at, updated_at,published_at,title, description,url,feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING  *;
-- name: GetPosts :many
SELECT * 
FROM posts
WHERE feed_id in (
    SELECT feeds.id
    FROM feeds
    INNER join feed_follow on feed_follow.feed_id = feeds.id
    INNER join users on feed_follow.user_id = $1
);

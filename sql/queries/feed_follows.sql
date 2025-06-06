-- name: CreateFeedFollow :one
WITH
  inserted_feed as (
    INSERT INTO
      feed_follow (id, created_at, updated_at, user_id, feed_id)
    VALUES
      ($1, $2, $3, $4, $5)
    RETURNING
      *
  )
SELECT
  inserted_feed.id,
  inserted_feed.created_at,
  inserted_feed.updated_at,
  inserted_feed.user_id,
  inserted_feed.feed_id,
  users.name as user_name,
  feeds.name as feed_name
FROM
  inserted_feed
  INNER JOIN users on users.id = inserted_feed.user_id
  INNER JOIN feeds on feeds.id = inserted_feed.feed_id;


-- name: GetFollowings :many
SELECT
  feeds.name as feed_name
FROM
  feed_follow
  INNER JOIN users on users.id = feed_follow.user_id
  INNER JOIN feeds on feeds.id = feed_follow.feed_id
WHERE users.id = $1
;

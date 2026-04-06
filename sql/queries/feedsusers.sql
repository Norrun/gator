-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO users_feeds (id, created_at, updated_at, user_id, feed_id)
    VALUES (
            $1,
            $2,
            $3,
            $4,
            $5
        )
    RETURNING *
)
SELECT inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
    INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
    INNER JOIN users ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
 SELECT users_feeds.*, feeds.name AS feed_name, users.name AS user_name FROM users_feeds
 INNER JOIN feeds ON users_feeds.feed_id = feeds.id
 INNER JOIN users ON users_feeds.user_id = users.id
 WHERE users.name = $1;
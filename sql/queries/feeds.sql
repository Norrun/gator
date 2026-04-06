-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name AS user_name FROM users_feeds
INNER JOIN users ON users_feeds.user_id = users.id
INNER JOIN feeds ON users_feeds.feed_id = feeds.id;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;
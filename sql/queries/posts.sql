-- name: CreatePost :one
INSERT INTO posts (id, created_at , updated_at , title , url, description, published_at, feed_id )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)Returning *;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts
INNER JOIN users_feeds
ON posts.feed_id = users_feeds.feed_id
INNER JOIN users
ON users.id = user_id
WHERE users.id = $1
ORDER BY published_at DESC NULLS LAST
LIMIT $2;
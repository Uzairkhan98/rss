-- name: CreatePost :one
WITH created_post AS (
    INSERT INTO posts (id, created_at, updated_at, title, url, feed_id, description, published_at)
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
    RETURNING *
) 
SELECT created_post.*, 
feed.name AS feed_name
FROM created_post 
LEFT JOIN feeds feed 
    ON created_post.feed_id = feed.id;

-- name: GetUserPosts :many
SELECT * FROM posts p 
LEFT JOIN feeds f
ON p.feed_id = f.id
LEFT JOIN users u
ON f.user_id = u.id
WHERE u.id = $1
ORDER BY p.published_at DESC
LIMIT $2;
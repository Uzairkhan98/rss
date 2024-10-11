-- name: CreateFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
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
feed.name AS feed_name, "user".name AS user_name
FROM inserted_feed_follow 
LEFT JOIN feeds feed 
    ON inserted_feed_follow.feed_id = feed.id
LEFT JOIN users "user" 
    ON inserted_feed_follow.user_id = "user".id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*,
feed.name AS feed_name, "user".name AS user_name
FROM feed_follows ff 
LEFT JOIN feeds feed 
    ON ff.feed_id = feed.id
LEFT JOIN users "user" 
    ON ff.user_id = "user".id
WHERE "user".name = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows ff
WHERE ff.user_id = $1 
AND ff.feed_id = (SELECT id FROM feeds where url = $2); 
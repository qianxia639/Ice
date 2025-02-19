-- name: ExistsFriendship :one
SELECT COUNT(*) FROM friend_requests 
WHERE user_id = $1 AND friend_id = $2 AND status = 2;

-- name: CreateFriendship :one
INSERT INTO friendships (
    user_id, friend_id, comment
) VALUES (
    $1, $2, $3
)
RETURNING *;
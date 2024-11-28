-- name: GetUser :one
SELECT
    u.id,
    u.name,
    u.picture,
    u.description,
    u.mail,
    u.createdAt,
    COALESCE(
        array_agg(f.follower_id) FILTER (WHERE f.follower_id IS NOT NULL),
        '{}'::uuid[]
    ) as followers
FROM users u 
LEFT JOIN user_followers f ON u.id = f.user_id
WHERE u.id = $1
GROUP BY u.id;

-- name: AddFollower :exec
INSERT INTO user_followers (user_id, follower_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveFollower :exec
DELETE FROM user_followers
WHERE user_id = $1 AND follower_id = $2;

-- name: SaveUser :exec
INSERT INTO users(
  id,
  name,
  description,
  picture,
  mail
) VALUES ( $1, $2, $3, $4, $5 );

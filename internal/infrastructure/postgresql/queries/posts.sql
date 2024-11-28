-- name: SavePosts :exec
INSERT INTO posts (
  id,
  comment,
  postedBy
) VALUES ( $1, $2, $3 );

-- name: FindPost :one
SELECT 
  id, 
  comment, 
  postedBy, 
  createdAt
FROM posts
WHERE id = $1;

-- name: FindRecentPost :many
SELECT 
  id, 
  comment, 
  postedBy, 
  createdAt 
FROM posts 
ORDER BY createdAt DESC;

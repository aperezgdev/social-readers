-- name: SaveComments :exec
INSERT INTO comments (
  id,
  content,
  postId,
  commentedBy
) VALUES ( $1, $2, $3, $4 );

-- name: FindCommentsByPost :many
SELECT 
  id, 
  content, 
  postId, 
  commentedBy, 
  createdAt 
FROM comments 
WHERE postId = $1;

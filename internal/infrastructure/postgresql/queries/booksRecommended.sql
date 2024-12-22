-- name: SaveBooksRecommended :exec
INSERT INTO booksRecommended (
  id,
  isbn,
  title,
  description,
  picture,
  userId
) VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetBooksRecommendedByUser :many
SELECT 
  id,
  isbn,
  title,
  description,
  picture,
  userId,
  createdAt
FROM booksRecommended
WHERE userId = $1;
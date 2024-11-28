-- name: SaveBooksToRead :exec
INSERT INTO booksToRead (
  id,
  isbn,
  title,
  description,
  picture,
  userId
) VALUES ( $1, $2, $3, $4, $5, $6 );


// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const findPost = `-- name: FindPost :one
SELECT 
  id, 
  comment, 
  postedBy, 
  createdAt
FROM posts
WHERE id = $1
`

func (q *Queries) FindPost(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRow(ctx, findPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Comment,
		&i.Postedby,
		&i.Createdat,
	)
	return i, err
}

const findRecentPost = `-- name: FindRecentPost :many
SELECT 
  id, 
  comment, 
  postedBy, 
  createdAt 
FROM posts 
ORDER BY createdAt DESC
`

func (q *Queries) FindRecentPost(ctx context.Context) ([]Post, error) {
	rows, err := q.db.Query(ctx, findRecentPost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Comment,
			&i.Postedby,
			&i.Createdat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const savePosts = `-- name: SavePosts :exec
INSERT INTO posts (
  id,
  comment,
  postedBy
) VALUES ( $1, $2, $3 )
`

type SavePostsParams struct {
	ID       uuid.UUID
	Comment  string
	Postedby pgtype.UUID
}

func (q *Queries) SavePosts(ctx context.Context, arg SavePostsParams) error {
	_, err := q.db.Exec(ctx, savePosts, arg.ID, arg.Comment, arg.Postedby)
	return err
}

package repository

import (
	"context"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	domain_repo "github.com/aperezgdev/social-readers-api/internal/domain/repository"
	comment_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/comment"
	post_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
	"github.com/aperezgdev/social-readers-api/internal/infrastructure/postgresql/sqlc"
	"github.com/aperezgdev/social-readers-api/pkg"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CommentPostgresRepository struct {
	queries *sqlc.Queries
}

func NewCommentRepository(queries *sqlc.Queries) domain_repo.CommentRepository {
	return CommentPostgresRepository{queries: queries}
}

func (cp CommentPostgresRepository) FindByPost(ctx context.Context, postId post_vo.PostId) ([]models.Comment, error) {
	result, err := cp.queries.FindCommentsByPost(ctx, pgtype.UUID{Bytes: uuid.MustParse(string(postId))})
	if err != nil {
		return nil, err
	}

	comments := pkg.Map(result, func(t sqlc.Comment) models.Comment {
		return models.Comment{
			Id:          comment_vo.ComentId(t.ID.String()),
			Content:     comment_vo.CommentContent(t.Content),
			CommentedBy: user_vo.UserId(uuid.UUID(t.Commentedby.Bytes).String()),
			PostId:      post_vo.PostId(uuid.UUID(t.Postid.Bytes).String()),
			CreatedAt:   shared_vo.CreatedAt(t.Createdat.Time),
		}
	})

	return comments, nil
}

func (cp CommentPostgresRepository) Save(ctx context.Context, comment models.Comment) error {
	saveCommentParams := sqlc.SaveCommentsParams{
		ID:      uuid.MustParse(string(comment.Id)),
		Content: string(comment.Content),
		Postid: pgtype.UUID{
			Bytes: uuid.MustParse(string(comment.PostId)),
			Valid: true,
		},
		Commentedby: pgtype.UUID{
			Bytes: uuid.MustParse(string(comment.CommentedBy)),
			Valid: true,
		},
	}

	return cp.queries.SaveComments(ctx, saveCommentParams)
}

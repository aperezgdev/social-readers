package repository

import (
	"context"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	post_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
	"github.com/aperezgdev/social-readers-api/internal/infrastructure/postgresql/sqlc"
	"github.com/aperezgdev/social-readers-api/pkg"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostPostgresRepository struct {
	queries *sqlc.Queries
}

func NewPostPostgresRepository(queries *sqlc.Queries) repository.PostRepository {
	return PostPostgresRepository{queries}
}

func (pp PostPostgresRepository) Find(ctx context.Context, id post_vo.PostId) (models.Post, error) {
	result, err := pp.queries.FindPost(ctx, uuid.MustParse(string(id)))
	if err != nil {
		return models.Post{}, err
	}

	post := models.Post{
		Id:        post_vo.PostId(result.ID.String()),
		Comment:   post_vo.PostComment(result.Comment),
		PostedBy:  user_vo.UserId(uuid.UUID(result.Postedby.Bytes).String()),
		CreatedAt: shared_vo.CreatedAt(result.Createdat.Time),
	}

	return post, nil
}

func (pp PostPostgresRepository) Save(ctx context.Context, post models.Post) error {
	savePostParams := sqlc.SavePostsParams{
		ID:      uuid.MustParse(string(post.Id)),
		Comment: string(post.Comment),
		Postedby: pgtype.UUID{
			Bytes: uuid.MustParse(string(post.PostedBy)),
			Valid: true,
		},
	}

	return pp.queries.SavePosts(ctx, savePostParams)
}

func (pp PostPostgresRepository) FindRecent(ctx context.Context) ([]models.Post, error) {
	result, err := pp.queries.FindRecentPost(ctx)
	if err != nil {
		return nil, err
	}

	posts := pkg.Map(result, func(t sqlc.Post) models.Post {
		return models.Post{
			Id:        post_vo.PostId(t.ID.String()),
			Comment:   post_vo.PostComment(t.Comment),
			PostedBy:  user_vo.UserId(uuid.UUID(t.Postedby.Bytes).String()),
			CreatedAt: shared_vo.CreatedAt(t.Createdat.Time),
		}
	})

	return posts, nil
}

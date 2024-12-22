package repository

import (
	"context"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
	"github.com/aperezgdev/social-readers-api/internal/infrastructure/postgresql/sqlc"
	"github.com/aperezgdev/social-readers-api/pkg"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type BookRecommendedRepository struct {
	queries *sqlc.Queries
}

func NewBookRecommendedRepository(queries *sqlc.Queries) BookRecommendedRepository {
	return BookRecommendedRepository{queries: queries}
}
func (br BookRecommendedRepository) FindByUser(ctx context.Context, userId user_vo.UserId) ([]models.BookRecommended, error) {
	result, err := br.queries.GetBooksRecommendedByUser(ctx, pgtype.UUID{Bytes: uuid.MustParse(string(userId))})

	if err != nil {
		return nil, err
	}

	return pkg.Map(result, func(t sqlc.Booksrecommended) models.BookRecommended {
		return models.BookRecommended{
			Isbn:        shared_vo.Isbn(t.Isbn),
			Title:       shared_vo.BookTitle(t.Title),
			Description: shared_vo.BookDescription(t.Description.String),
			Picture:     shared_vo.BookPicture(t.Picture.String),
			UserId:      user_vo.UserId(uuid.UUID(t.Userid.Bytes).String()),
			CreatedAt:   shared_vo.CreatedAt(t.Createdat.Time),
		}
	}), nil
}

func (br BookRecommendedRepository) Save(ctx context.Context, book models.BookRecommended) error {
	err := br.queries.SaveBooksRecommended(ctx, sqlc.SaveBooksRecommendedParams{
		ID: 		 uuid.MustParse(string(book.Id)),
		Isbn:        string(book.Isbn),
		Title:       string(book.Title),
		Description: pgtype.Text{String: string(book.Description)},
		Picture:     pgtype.Text{String: string(book.Picture)},
		Userid:      pgtype.UUID{Bytes: uuid.MustParse(string(book.UserId))},
	})

	if err != nil {
		return err
	}

	return nil
}
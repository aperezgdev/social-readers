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

type BookToReadRepository struct {
	queries *sqlc.Queries
}

func NewBookToReadRepository(queries *sqlc.Queries) BookToReadRepository {
	return BookToReadRepository{queries: queries}
}

func (br BookToReadRepository) FindByUser(ctx context.Context, userId user_vo.UserId) ([]models.BookToRead, error) {
	result, err := br.queries.GetBooksToReadByUser(ctx, pgtype.UUID{Bytes: uuid.MustParse(string(userId))})

	if err != nil {
		return nil, err
	}

	return pkg.Map(result, func(t sqlc.Bookstoread) models.BookToRead {
		return models.BookToRead{
			Isbn:        shared_vo.Isbn(t.Isbn),
			Title:       shared_vo.BookTitle(t.Title),
			Description: shared_vo.BookDescription(t.Description.String),
			Picture:     shared_vo.BookPicture(t.Picture.String),
			UserId:      user_vo.UserId(uuid.UUID(t.Userid.Bytes).String()),
			CreatedAt:   shared_vo.CreatedAt(t.Createdat.Time),
		}
	}), nil
}

func (br BookToReadRepository) Save(ctx context.Context, book models.BookToRead) error {
	err := br.queries.SaveBooksToRead(ctx, sqlc.SaveBooksToReadParams{
		ID:          uuid.MustParse(string(book.Id)),
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
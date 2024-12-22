package finder

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type BookRecommendedFinderByUser struct {
	slog *slog.Logger
	bookRecommendedRepository repository.BookRecommendedRepository
}

func NewBookRecommendedFinderByUser(
	slog *slog.Logger,
	bookRecommendedRepository repository.BookRecommendedRepository,
) BookRecommendedFinderByUser {
	return BookRecommendedFinderByUser{
		slog:                     slog,
		bookRecommendedRepository: bookRecommendedRepository,
	}
}

func (bf *BookRecommendedFinderByUser) Run(ctx context.Context, userId string) ([]models.BookRecommended, error) {
	bookRecommendeds, err := bf.bookRecommendedRepository.FindByUser(ctx, user_vo.UserId(userId))
	if err != nil {
		return nil, err
	}

	return bookRecommendeds, nil
}
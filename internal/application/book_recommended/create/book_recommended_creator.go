package create

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	finder "github.com/aperezgdev/social-readers-api/internal/domain/use_case/user"
)

type BookRecommendedCreator struct {
	slog                      *slog.Logger
	bookRecommendedRepository repository.BookRecommendedRepository
	userFinder                finder.UserFinder
}

func NewBookRecommendedCreator(
	slog *slog.Logger,
	bookRecommendedRepository repository.BookRecommendedRepository,
	userRepository repository.UserRepository,
) BookRecommendedCreator {
	return BookRecommendedCreator{
		slog:                      slog,
		bookRecommendedRepository: bookRecommendedRepository,
		userFinder:                finder.NewUserFinder(slog, userRepository),
	}
}

func (br *BookRecommendedCreator) Run(
	ctx context.Context,
	isbn, title, description, picture, userId string,
) error {
	_, err := br.userFinder.Run(ctx, userId)
	if err != nil {
		return err
	}

	bookRecommended := models.NewBookRecommended(isbn, title, description, userId, picture)

	return br.bookRecommendedRepository.Save(ctx, bookRecommended)
}

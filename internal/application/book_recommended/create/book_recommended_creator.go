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
	br.slog.Info(
		"BookRecommendedCreator - Run - Params into: ",
		slog.Any("isbn", isbn),
		slog.Any("title", title),
		slog.Any("description", description),
		slog.Any("userId", userId),
		slog.Any("picture", picture),
	)

	bookRecommended, validationError := models.NewBookRecommended(isbn, title, description, userId, picture)
	if validationError != nil {
		br.slog.Info("BookRecommendedCreator - Run - Validation error: ", slog.Any("error", validationError))
		return validationError
	}

	_, err := br.userFinder.Run(ctx, userId)
	if err != nil {
		br.slog.Info("BookRecommendedCreator - Run - User not found: ", slog.Any("error", err))
		return err
	}

	return br.bookRecommendedRepository.Save(ctx, bookRecommended)
}

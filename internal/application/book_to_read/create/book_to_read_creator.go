package create

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	finder "github.com/aperezgdev/social-readers-api/internal/domain/use_case/user"
)

type BookToReadCreator struct {
	slog                 *slog.Logger
	bookToReadRepository repository.BookToReadRepository
	userFinder           finder.UserFinder
}

func NewBookToReadCreator(
	slog *slog.Logger,
	bookToReadRepository repository.BookToReadRepository,
	userRepository repository.UserRepository,
) BookToReadCreator {
	return BookToReadCreator{
		slog:                 slog,
		bookToReadRepository: bookToReadRepository,
		userFinder:           finder.NewUserFinder(slog, userRepository),
	}
}

func (bc *BookToReadCreator) Run(
	ctx context.Context,
	isbn, title, description, picture, userId string,
) error {
	bc.slog.Info(
		"BookToReadCreator - Run - Params into: ",
		slog.Any("isbn", isbn),
		slog.Any("title", title),
		slog.Any("description", description),
		slog.Any("userId", userId),
		slog.Any("picture", picture),
	)

	bookToRead, validationError := models.NewBookToRead(isbn, title, description, picture, userId)
	if validationError != nil {
		bc.slog.Info("BookToReadCreator - Run - Validation error: ", slog.Any("error", validationError))
		return validationError
	}

	_, err := bc.userFinder.Run(ctx, userId)
	if err != nil {
		bc.slog.Info("BookToReadCreator - Run - User not found: ", slog.Any("error", err))
		return err
	}

	return bc.bookToReadRepository.Save(ctx, bookToRead)
}

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
	isbn, title, description, userId, picture string,
) error {
	_, err := bc.userFinder.Run(ctx, userId)
	if err != nil {
		return err
	}

	bookToRead := models.NewBookToRead(isbn, title, description, userId, picture)

	bc.bookToReadRepository.Save(ctx, bookToRead)

	return nil
}

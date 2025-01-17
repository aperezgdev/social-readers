package finder

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type BookToReadFinderByUser struct {
	slog                 *slog.Logger
	bookToReadRepository repository.BookToReadRepository
}

func NewBookToReadFinderByUser(
	slog *slog.Logger,
	bookToReadRepository repository.BookToReadRepository,
) BookToReadFinderByUser {
	return BookToReadFinderByUser{
		slog:                 slog,
		bookToReadRepository: bookToReadRepository,
	}
}

func (bf *BookToReadFinderByUser) Run(
	ctx context.Context,
	userId string,
) ([]models.BookToRead, error) {
	bf.slog.Info(
		"BookToReadFinderByUser - Run - Params into: ",
		slog.Any("userId", userId),
	)

	bookToReads, err := bf.bookToReadRepository.FindByUser(ctx, user_vo.UserId(userId))
	if err != nil {
		bf.slog.Info(
			"BookToReadFinderByUser - Run - Error: ")
		return nil, err
	}

	return bookToReads, nil
}
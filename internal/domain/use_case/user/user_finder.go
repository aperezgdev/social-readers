package finder

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type UserFinder struct {
	slog           *slog.Logger
	userRepository repository.UserRepository
}

func NewUserFinder(slog *slog.Logger, userRepository repository.UserRepository) UserFinder {
	return UserFinder{slog, userRepository}
}

func (uf UserFinder) Run(ctx context.Context, id string) (models.User, error) {
	uf.slog.Info("UserFinder - Run - Params into: ", slog.Any("id", id))
	return uf.userRepository.Find(ctx, user_vo.UserId(id))
}

package create

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
)

type UserCreator struct {
	slog           *slog.Logger
	userRepository repository.UserRepository
}

func NewUserCreator(slog *slog.Logger, userRepository repository.UserRepository) UserCreator {
	return UserCreator{slog, userRepository}
}

func (uc UserCreator) Run(ctx context.Context, name, mail, picture string) error {
	uc.slog.Info(
		"UserCreator - Run - Params into: ",
		slog.Any("name", name),
		slog.Any("mail", mail),
		slog.Any("picture", picture),
	)
	user := models.NewUser(name, picture, mail)

	return uc.userRepository.Save(ctx, user)
}

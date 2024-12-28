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
	user, err := models.NewUser(name, picture, mail)

	if err != nil {
		uc.slog.Info("UserCreator - Run - Validation error: ", slog.Any("error", err))
		return err
	}

	return uc.userRepository.Save(ctx, user)
}

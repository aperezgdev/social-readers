package create

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	finder "github.com/aperezgdev/social-readers-api/internal/domain/use_case/user"
)

type PostCreator struct {
	slog           *slog.Logger
	postRepository repository.PostRepository
	userFinder     finder.UserFinder
}

func NewPostCreator(
	slog *slog.Logger,
	postRepository repository.PostRepository,
	userRepository repository.UserRepository,
) PostCreator {
	return PostCreator{
		slog:           slog,
		postRepository: postRepository,
		userFinder:     finder.NewUserFinder(slog, userRepository),
	}
}

func (pc PostCreator) Run(ctx context.Context, comment, postedBy string) error {
	pc.slog.Info("PostCreator - Run - Params into: ", slog.Any("comment", comment))

	post, errValidation := models.NewPost(comment, postedBy)
	if errValidation != nil {
		pc.slog.Info("PostCreator - Run - Validation error: ", slog.Any("error", errValidation))
		return errValidation
	}

	_, err := pc.userFinder.Run(ctx, postedBy)
	if err != nil {
		pc.slog.Info("PostCreator - Run - User not exists")
		return err
	}

	return pc.postRepository.Save(ctx, post)
}

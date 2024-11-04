package finder

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	post_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
)

type PostFinder struct {
	slog           *slog.Logger
	postRepository repository.PostRepository
}

func NewPostFinder(slog *slog.Logger, postRepository repository.PostRepository) PostFinder {
	return PostFinder{slog, postRepository}
}

func (pf *PostFinder) Run(ctx context.Context, id string) (models.Post, error) {
	pf.slog.Info("PostFinder - Run - Params into: ", slog.Any("id", id))
	return pf.postRepository.Find(ctx, post_vo.PostId(id))
}

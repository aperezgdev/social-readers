package finder

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
)

type PostRecentFinder struct {
	slog           *slog.Logger
	postRepository repository.PostRepository
}

func NewPostRecentFinder(
	slog *slog.Logger,
	postRepository repository.PostRepository,
) PostRecentFinder {
	return PostRecentFinder{
		slog,
		postRepository,
	}
}

func (psl *PostRecentFinder) Run(ctx context.Context) ([]models.Post, error) {
	psl.slog.Info("PostSelectorLasted - Run - ")
	return psl.postRepository.FindRecent(ctx)
}

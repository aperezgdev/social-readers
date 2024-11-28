package finder

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	finder "github.com/aperezgdev/social-readers-api/internal/domain/use_case/post"
	post_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
)

type CommentFinderByPost struct {
	slog              *slog.Logger
	commentRepository repository.CommentRepository
	postFinder        finder.PostFinder
}

func NewCommentFinderByPost(
	slog *slog.Logger,
	commentRepository repository.CommentRepository,
	postRepository repository.PostRepository,
) CommentFinderByPost {
	return CommentFinderByPost{
		slog:              slog,
		commentRepository: commentRepository,
		postFinder:        finder.NewPostFinder(slog, postRepository),
	}
}

func (cf *CommentFinderByPost) Run(ctx context.Context, postId string) ([]models.Comment, error) {
	cf.slog.Info("CommentFinderByPost - Run - Params into: ", slog.Any("postId", postId))
	_, err := cf.postFinder.Run(ctx, postId)
	if err != nil {
		cf.slog.Warn("CommentFinderByPost - Run - Post not exists", slog.Any("postId", postId))
		return make([]models.Comment, 0), err
	}

	return cf.commentRepository.FindByPost(ctx, post_vo.PostId(postId))
}

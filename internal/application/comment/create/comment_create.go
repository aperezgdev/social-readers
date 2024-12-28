package create

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	finder_post "github.com/aperezgdev/social-readers-api/internal/domain/use_case/post"
	finder "github.com/aperezgdev/social-readers-api/internal/domain/use_case/user"
)

type CommentCreator struct {
	slog              *slog.Logger
	commentRepository repository.CommentRepository
	userFinder        finder.UserFinder
	postFinder        finder_post.PostFinder
}

func NewCommentCreator(
	slog *slog.Logger,
	commentRepository repository.CommentRepository,
	userRepository repository.UserRepository,
	postRepository repository.PostRepository,
) CommentCreator {
	return CommentCreator{
		slog,
		commentRepository,
		finder.NewUserFinder(slog, userRepository),
		finder_post.NewPostFinder(slog, postRepository),
	}
}

func (cc *CommentCreator) Run(ctx context.Context, content, postId, commentedBy string) error {
	comment, validationError := models.NewComment(content, postId, commentedBy)
	if validationError != nil {
		cc.slog.Info("CommentCreator - Run - Validation error: ", slog.Any("error", validationError))
		return validationError
	}
	
	cc.slog.Info(
		"CommentCreator - Run - Params into: ",
		slog.Any("content", content),
		slog.Any("postId", postId),
		slog.Any("commentedBy", commentedBy),
	)
	_, errUser := cc.userFinder.Run(ctx, commentedBy)
	if errUser != nil {
		cc.slog.Info("CommentCreator - Run - User not exists")
		return errUser
	}

	_, errPost := cc.postFinder.Run(ctx, postId)
	if errPost != nil {
		cc.slog.Info("CommentCreator - Run - Post not exists")
		return errPost
	}

	return cc.commentRepository.Save(ctx, comment)
}

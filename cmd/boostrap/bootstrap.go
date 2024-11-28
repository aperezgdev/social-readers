package boostrap

import (
	"context"
	"log/slog"
	"os"

	create_comment "github.com/aperezgdev/social-readers-api/internal/application/comment/create"
	finder_comment "github.com/aperezgdev/social-readers-api/internal/application/comment/find"
	"github.com/aperezgdev/social-readers-api/internal/application/post/create"
	finder_post "github.com/aperezgdev/social-readers-api/internal/application/post/find"
	create_user "github.com/aperezgdev/social-readers-api/internal/application/user/create"
	finder_user "github.com/aperezgdev/social-readers-api/internal/domain/use_case/user"
	"github.com/aperezgdev/social-readers-api/internal/infrastructure/config"
	server "github.com/aperezgdev/social-readers-api/internal/infrastructure/http"
	"github.com/aperezgdev/social-readers-api/internal/infrastructure/http/controller"
	"github.com/aperezgdev/social-readers-api/internal/infrastructure/postgresql/repository"
	"github.com/aperezgdev/social-readers-api/internal/infrastructure/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

func Run() error {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	conf := config.NewConfig(logger)

	conn, err := pgx.Connect(context.Background(), conf.DatabaseUrl)
	if err != nil {
		panic(err)
	}

	healthController := controller.NewHealthController()
	queries := sqlc.New(conn)

	userRepository := repository.NewUserRepository(queries)
	commentRepository := repository.NewCommentRepository(queries)
	postRepository := repository.NewPostPostgresRepository(queries)

	userCreator := create_user.NewUserCreator(logger, userRepository)
	userFinder := finder_user.NewUserFinder(logger, userRepository)
	commentCreator := create_comment.NewCommentCreator(logger, commentRepository, userRepository, postRepository)
	commentFinderByPost := finder_comment.NewCommentFinderByPost(logger, commentRepository, postRepository)
	postCreator := create.NewPostCreator(logger, postRepository, userRepository)
	postFinderRecent := finder_post.NewPostRecentFinder(logger, postRepository)

	userController := controller.NewUserController(userCreator, userFinder)
	commentController := controller.NewCommentController(commentCreator, commentFinderByPost)
	postController := controller.NewPostController(postCreator, postFinderRecent)

	httpServer := server.NewHttpServer(logger, conf)
	httpServer.AddHandler("/health", healthController.GetHealth)
	httpServer.AddHandler("/users/{id}", userController.GetUser)
	httpServer.AddHandler("POST /users", userController.PostUser)
	httpServer.AddHandler("/posts", postController.GetPost)
	httpServer.AddHandler("POST /posts", postController.PostPost)
	httpServer.AddHandler("/posts/{postId}/comments", commentController.GetCommentByPost)
	httpServer.AddHandler("POST /comments", commentController.PostComment)

	return httpServer.Start()
}

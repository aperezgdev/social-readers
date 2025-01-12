package boostrap

import (
	"context"
	"log/slog"
	"os"

	create_book_recommended "github.com/aperezgdev/social-readers-api/internal/application/book_recommended/create"
	finder_book_recommended "github.com/aperezgdev/social-readers-api/internal/application/book_recommended/find"
	"github.com/aperezgdev/social-readers-api/internal/application/book_to_read/create"
	finder "github.com/aperezgdev/social-readers-api/internal/application/book_to_read/find"
	create_comment "github.com/aperezgdev/social-readers-api/internal/application/comment/create"
	finder_comment "github.com/aperezgdev/social-readers-api/internal/application/comment/find"
	create_post "github.com/aperezgdev/social-readers-api/internal/application/post/create"
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
	bookRecommendedRepository := repository.NewBookRecommendedRepository(queries)
	bookToReadRepository := repository.NewBookToReadRepository(queries)

	userCreator := create_user.NewUserCreator(logger, userRepository)
	userFinder := finder_user.NewUserFinder(logger, userRepository)
	commentCreator := create_comment.NewCommentCreator(logger, commentRepository, userRepository, postRepository)
	commentFinderByPost := finder_comment.NewCommentFinderByPost(logger, commentRepository, postRepository)
	postCreator := create_post.NewPostCreator(logger, postRepository, userRepository)
	postFinderRecent := finder_post.NewPostRecentFinder(logger, postRepository)
	bookRecommendedCreator := create_book_recommended.NewBookRecommendedCreator(logger, bookRecommendedRepository, userRepository)
	bookRecommendedFinderByUser := finder_book_recommended.NewBookRecommendedFinderByUser(logger, bookRecommendedRepository)
	bookToReadFinderByUser := finder.NewBookToReadFinderByUser(logger, bookToReadRepository)
	bookToReadCreator := create.NewBookToReadCreator(logger, bookToReadRepository, userRepository)

	userController := controller.NewUserController(userCreator, userFinder)
	commentController := controller.NewCommentController(commentCreator, commentFinderByPost)
	postController := controller.NewPostController(postCreator, postFinderRecent)
	bookRecommendedController := controller.NewBookRecommendedController(bookRecommendedFinderByUser, bookRecommendedCreator)
	bookToReadController := controller.NewBookToReadsController(bookToReadFinderByUser, bookToReadCreator)

	httpServer := server.NewHttpServer(logger, conf)
	httpServer.AddHandler("/health", healthController.GetHealth)
	httpServer.AddHandler("/users/{id}", userController.GetUser)
	httpServer.AddHandler("POST /users", userController.PostUser)
	httpServer.AddHandler("/posts", postController.GetPost)
	httpServer.AddHandler("POST /posts", postController.PostPost)
	httpServer.AddHandler("/posts/{postId}/comments", commentController.GetCommentByPost)
	httpServer.AddHandler("POST /comments", commentController.PostComment)
	httpServer.AddHandler("/users/{userId}/book-recommended", bookRecommendedController.GetBookRecommendedByUser)
	httpServer.AddHandler("POST /book-recommended", bookRecommendedController.PostBookRecommended)
	httpServer.AddHandler("/users/{userId}/book-to-read", bookToReadController.GetBooksToReadByUser)
	httpServer.AddHandler("POST /book-to-read", bookToReadController.PostBookToRead)

	return httpServer.Start()
}

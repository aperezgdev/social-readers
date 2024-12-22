package finder_test

import (
	"context"
	"log/slog"
	"testing"

	finder "github.com/aperezgdev/social-readers-api/internal/application/book_to_read/find"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testSuite struct {
	mockCommentRepository *repository.MockBookToReadRepository
	bookToReadFinderByUser *finder.BookToReadFinderByUser
}

func setupTest() *testSuite {
	mockBookToReadRepository := repository.MockBookToReadRepository{}
	logger := slog.Default()
	bookToReadFinderByUser := finder.NewBookToReadFinderByUser(
		logger,
		&mockBookToReadRepository,
	)

	return &testSuite{
		mockCommentRepository: &mockBookToReadRepository,
		bookToReadFinderByUser: bookToReadFinderByUser,
	}
}

func TestBookToReadFinderByUser(t *testing.T) {
	t.Parallel()

	t.Run("should find book to read by user", func(t *testing.T) {
		suite := setupTest()
		suite.mockCommentRepository.On("FindByUser", mock.Anything, mock.Anything).
			Return([]models.BookToRead{
				{
					Isbn:        "123123",
					Title:       "Hobbit",
					Description: "Description",
					Picture:     "picture",
					UserId:      "1",
				},
			}, nil).
			Once()

		bookToReads, err := suite.bookToReadFinderByUser.Run(
			context.Background(),
			"1",
		)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(bookToReads))
		suite.mockCommentRepository.AssertExpectations(t)
	})
}
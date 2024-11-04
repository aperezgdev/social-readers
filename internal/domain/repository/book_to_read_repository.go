package repository

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type BookToReadRepository interface {
	FindByUser(context.Context, user_vo.UserId) ([]models.BookToRead, error)
	Save(context.Context, models.BookToRead) error
}

type MockBookToReadRepository struct {
	mock.Mock
}

func (mb *MockBookToReadRepository) FindByUser(
	ctx context.Context,
	id user_vo.UserId,
) ([]models.BookToRead, error) {
	args := mb.Called(ctx, id)

	return args.Get(0).([]models.BookToRead), args.Error(1)
}

func (mb *MockBookToReadRepository) Save(ctx context.Context, bookToRead models.BookToRead) error {
	args := mb.Called(ctx, bookToRead)

	return args.Error(0)
}

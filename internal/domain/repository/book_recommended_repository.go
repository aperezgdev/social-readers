package repository

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type BookRecommendedRepository interface {
	FindByUser(context.Context, user_vo.UserId) ([]models.BookRecommended, error)
	Save(context.Context, models.BookRecommended) error
}

type MockBookRecommendedRepository struct {
	mock.Mock
}

func (mb *MockBookRecommendedRepository) FindByUser(
	ctx context.Context,
	userId user_vo.UserId,
) ([]models.BookRecommended, error) {
	args := mb.Called(ctx, userId)

	return args.Get(0).([]models.BookRecommended), args.Error(1)
}

func (mb *MockBookRecommendedRepository) Save(
	ctx context.Context,
	bookRecommended models.BookRecommended,
) error {
	args := mb.Called(ctx, bookRecommended)

	return args.Error(0)
}

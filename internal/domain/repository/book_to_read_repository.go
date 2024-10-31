package repository

import (
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type BookToReadRepository interface {
	FindByUser(user_vo.UserId) []models.BookToRead
	Save(models.BookToRead) error
}

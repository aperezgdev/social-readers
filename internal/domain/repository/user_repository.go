package repository

import (
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type UserRepository interface {
	Find(user_vo.UserId)
	Save(models.User)
}

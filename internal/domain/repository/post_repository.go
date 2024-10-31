package repository

import (
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	post_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
)

type PostRepository interface {
	Find(post_vo.PostId)
	Save(models.Post)
}

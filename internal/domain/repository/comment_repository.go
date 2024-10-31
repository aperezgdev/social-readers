package repository

import (
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	comment_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/comment"
)

type CommentRepositoryRepository interface {
	Find(comment_vo.ComentId) models.Comment
	Save(models.Comment) error
}

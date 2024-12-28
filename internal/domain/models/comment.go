package models

import (
	. "github.com/aperezgdev/social-readers-api/internal/domain/value_object/comment"
	post_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
	"github.com/aperezgdev/social-readers-api/pkg"
)

type Comment struct {
	Id          ComentId
	Content     CommentContent
	PostId      post_vo.PostId
	CommentedBy user_vo.UserId
	CreatedAt   shared_vo.CreatedAt
}

func NewComment(content, postId, commentedBy string) (Comment, error) {
	contentVO, errValidationCommentContent := NewCommentContent(content)
	if errValidationCommentContent != nil {
		return Comment{}, errValidationCommentContent
	}

	errValidationPostId := pkg.ValidUUID(postId, "postId")
	if errValidationPostId != nil {
		return Comment{}, errValidationPostId
	}

	errValidationCommentedBy := pkg.ValidUUID(commentedBy, "commentedBy")
	if errValidationCommentedBy != nil {
		return Comment{}, errValidationCommentedBy
	}

	return Comment{
		Id:          NewCommentId(),
		Content:     contentVO,
		PostId:      post_vo.PostId(postId),
		CommentedBy: user_vo.UserId(commentedBy),
		CreatedAt:   shared_vo.NewCreatedAt(),
	}, nil
}

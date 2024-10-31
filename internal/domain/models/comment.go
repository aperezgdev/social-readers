package models

import (
	. "github.com/aperezgdev/social-readers-api/internal/domain/value_object/comment"
	post_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type Comment struct {
	Id          ComentId
	Content     CommentContent
	PostId      post_vo.PostId
	CommentedBy user_vo.UserId
	CreatedAt   shared_vo.CreatedAt
}

func NewComment(content, postId, commentedBy string) Comment {
	return Comment{
		Id:          NewCommentId(),
		Content:     NewCommentContent(content),
		PostId:      post_vo.PostId(postId),
		CommentedBy: user_vo.UserId(commentedBy),
		CreatedAt:   shared_vo.NewCreatedAt(),
	}
}

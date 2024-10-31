package models

import (
	. "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type Post struct {
	Id        PostId
	Comment   PostComment
	PostedBy  user_vo.UserId
	CreatedAt shared_vo.CreatedAt
}

func NewPost(comment, postedBy string) Post {
	return Post{
		Id:        NewPostId(),
		Comment:   NewPostComment(comment),
		PostedBy:  user_vo.UserId(postedBy),
		CreatedAt: shared_vo.NewCreatedAt(),
	}
}

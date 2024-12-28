package models

import (
	. "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
	"github.com/aperezgdev/social-readers-api/pkg"
)

type Post struct {
	Id        PostId
	Comment   PostComment
	PostedBy  user_vo.UserId
	CreatedAt shared_vo.CreatedAt
}

func NewPost(comment, postedBy string) (Post, error) {
	commentVO, err :=  NewPostComment(comment)
	if err != nil {
		return Post{}, err
	}

	validError := pkg.ValidUUID(postedBy, "postedBy")
	if validError != nil {
		return Post{}, validError
	}

	return Post{
		Id:        NewPostId(),
		Comment:   commentVO,
		PostedBy:  user_vo.UserId(postedBy),
		CreatedAt: shared_vo.NewCreatedAt(),
	}, nil
}

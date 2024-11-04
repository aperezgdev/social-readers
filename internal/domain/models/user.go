package models

import (
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	. "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type User struct {
	Id          UserId
	Name        UserName
	Picture     UserPicture
	Description UserDescription
	Followers   UserFollowers
	Mail        UserMail
	CreatedAt   shared_vo.CreatedAt
}

func NewUser(name, picture, mail string) User {
	return User{
		Id:          NewUserId(),
		Name:        NewUserName(name),
		Picture:     NewUserPicture(picture),
		Description: NewUserDescription(),
		Followers:   NewUserFollower(),
		Mail:        NewUserMail(mail),
		CreatedAt:   shared_vo.NewCreatedAt(),
	}
}

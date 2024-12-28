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

func NewUser(name, picture, mail string) (User, error) {
	nameVO, errName := NewUserName(name)
	if errName != nil {
		return User{}, errName
	}
	mailVO, errMail := NewUserMail(mail)
	if errMail != nil {
		return User{}, errMail
	}

	return User{
		Id:          NewUserId(),
		Name:        nameVO,
		Picture:     NewUserPicture(picture),
		Description: NewUserDescription(),
		Followers:   NewUserFollower(),
		Mail:        mailVO,
		CreatedAt:   shared_vo.NewCreatedAt(),
	}, nil
}

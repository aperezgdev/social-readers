package models

import (
	. "github.com/aperezgdev/social-readers-api/internal/domain/value_object/book_to_read"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type BookToRead struct {
	Id          BookToReadId
	Isbn        shared_vo.Isbn
	Title       shared_vo.BookTitle
	Description shared_vo.BookDescription
	Picture     shared_vo.BookPicture
	UserId      user_vo.UserId
	CreatedAt   shared_vo.CreatedAt
}

func NewBookToRead(isbn, title, description, userId, picture string) BookToRead {
	return BookToRead{
		Id:          NewBookToReadId(),
		Isbn:        shared_vo.NewIsbn(isbn),
		Title:       shared_vo.NewBookTitle(title),
		Description: shared_vo.NewBookDescription(description),
		Picture:     shared_vo.NewBookPicture(picture),
		UserId:      user_vo.UserId(userId),
		CreatedAt:   shared_vo.NewCreatedAt(),
	}
}

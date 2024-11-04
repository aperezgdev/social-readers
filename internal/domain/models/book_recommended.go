package models

import (
	book_recommended_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/book_recommended"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type BookRecommended struct {
	Id          book_recommended_vo.BookRecommendedId
	Isbn        shared_vo.Isbn
	Title       shared_vo.BookTitle
	Description shared_vo.BookDescription
	Picture     shared_vo.BookPicture
	UserId      user_vo.UserId
	CreatedAt   shared_vo.CreatedAt
}

func NewBookRecommended(isbn, title, description, userId, picture string) BookRecommended {
	return BookRecommended{
		Id:          book_recommended_vo.NewBookRecommendedId(),
		Isbn:        shared_vo.NewIsbn(isbn),
		Title:       shared_vo.NewBookTitle(title),
		Description: shared_vo.NewBookDescription(description),
		Picture:     shared_vo.NewBookPicture(picture),
		UserId:      user_vo.UserId(userId),
		CreatedAt:   shared_vo.NewCreatedAt(),
	}
}

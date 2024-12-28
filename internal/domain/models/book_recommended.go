package models

import (
	book_recommended_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/book_recommended"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
	"github.com/aperezgdev/social-readers-api/pkg"
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

func NewBookRecommended(isbn, title, description, userId, picture string) (BookRecommended, error) {
	isbnVO, errValidationIsbn := shared_vo.NewIsbn(isbn)
	if errValidationIsbn != nil {
		return BookRecommended{}, errValidationIsbn
	}

	titleVO, errValidationTitle := shared_vo.NewBookTitle(title)
	if errValidationTitle != nil {
		return BookRecommended{}, errValidationTitle
	}

	errValidationUserId := pkg.ValidUUID(userId, "UserId")
	if errValidationUserId != nil {
		return BookRecommended{}, errValidationUserId
	}

	return BookRecommended{
		Id:          book_recommended_vo.NewBookRecommendedId(),
		Isbn:        isbnVO,
		Title:       titleVO,
		Description: shared_vo.NewBookDescription(description),
		Picture:     shared_vo.NewBookPicture(picture),
		UserId:      user_vo.UserId(userId),
		CreatedAt:   shared_vo.NewCreatedAt(),
	}, nil
}

package models

import (
	. "github.com/aperezgdev/social-readers-api/internal/domain/value_object/book_to_read"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
	"github.com/aperezgdev/social-readers-api/pkg"
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

func NewBookToRead(isbn, title, description, picture, userId string) (BookToRead, error) {
	isbnVO, errValidationIsbn := shared_vo.NewIsbn(isbn)
	if errValidationIsbn != nil {
		return BookToRead{}, errValidationIsbn
	}

	titleVO, errValidationTitle := shared_vo.NewBookTitle(title)
	if errValidationTitle != nil {
		return BookToRead{}, errValidationTitle
	}

	errValidationUserId := pkg.ValidUUID(userId, "UserId")
	if errValidationUserId != nil {
		return BookToRead{}, errValidationUserId
	}

	return BookToRead{
		Id:          NewBookToReadId(),
		Isbn:        isbnVO,
		Title:       titleVO,
		Description: shared_vo.NewBookDescription(description),
		Picture:     shared_vo.NewBookPicture(picture),
		UserId:      user_vo.UserId(userId),
		CreatedAt:   shared_vo.NewCreatedAt(),
	}, nil
}

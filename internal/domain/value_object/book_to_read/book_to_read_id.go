package book_to_read_vo

import "github.com/google/uuid"

type BookToReadId string

func NewBookToReadId() BookToReadId {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return BookToReadId(id.String())
}

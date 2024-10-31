package book_recommended_vo

import "github.com/google/uuid"

type BookRecommendedId string

func NewBookRecommendedId() BookRecommendedId {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return BookRecommendedId(id.String())
}

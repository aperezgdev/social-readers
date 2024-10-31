package post_vo

import "github.com/google/uuid"

type PostId string

func NewPostId() PostId {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return PostId(id.String())
}

package comment_vo

import "github.com/google/uuid"

type ComentId string

func NewCommentId() ComentId {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return ComentId(id.String())
}

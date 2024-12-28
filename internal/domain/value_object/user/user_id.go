package user_vo

import "github.com/google/uuid"

type UserId string

func NewUserId() UserId {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return UserId(id.String())
}
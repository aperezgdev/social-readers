package user_vo

import "github.com/aperezgdev/social-readers-api/internal/domain/errors"

type UserName string

func NewUserName(userName string) (UserName, error) {
	return UserName(userName), validateName(userName)
}

func validateName(un string) error {
	if len(un) == 0 {
		return errors.FieldRequired("Name")
	}
	if len(string(un)) < 2 || len(string(un)) > 20 {
		return errors.OutRange("Name", 2, 20)
	}

	return nil
}

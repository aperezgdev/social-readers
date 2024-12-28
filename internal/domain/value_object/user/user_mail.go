package user_vo

import (
	"net/mail"

	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
)

type UserMail string

func NewUserMail(userMail string) (UserMail, error) {
	return UserMail(userMail), validateMail(userMail)
}

func validateMail(um string) error {
	_, err := mail.ParseAddress(string(um))

	if err != nil {
		return errors.FormatInvalidad("Mail")
	}

	return nil
}

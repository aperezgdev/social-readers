package user_vo

import "net/mail"

type UserMail string

func NewUserMail(userMail string) UserMail {
	return UserMail(userMail)
}

func (um UserMail) Validate() bool {
	_, err := mail.ParseAddress(string(um))

	return err == nil
}

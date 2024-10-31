package user_vo

type UserDescription string

func NewUserDescription(userDescription string) UserDescription {
	return UserDescription(userDescription)
}

func (ud UserDescription) Validate() bool {
	return len(string(ud)) < 240
}

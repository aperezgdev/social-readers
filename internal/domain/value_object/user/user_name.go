package user_vo

type UserName string

func NewUserName(userName string) UserName {
	return UserName(userName)
}

func (un UserName) Validate() bool {
	return len(string(un)) > 2 && len(string(un)) <= 20
}

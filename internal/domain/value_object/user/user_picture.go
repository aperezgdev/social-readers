package user_vo

type UserPicture string

func NewUserPicture(userPicture string) UserPicture {
	return UserPicture(userPicture)
}

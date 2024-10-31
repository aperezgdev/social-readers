package user_vo

type UserFollowers []UserId

func NewUserFollower() UserFollowers {
	return make(UserFollowers, 0)
}

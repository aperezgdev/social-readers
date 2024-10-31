package shared_vo

type BookDescription string

func NewBookDescription(bookDescription string) BookDescription {
	return BookDescription(bookDescription)
}

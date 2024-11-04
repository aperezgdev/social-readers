package shared_vo

type BookAuthor string

func NewBookAuthor(author string) BookAuthor {
	return BookAuthor(author)
}

func (ba BookAuthor) Validate() bool {
	return len(ba) > 1 && len(ba) <= 50
}

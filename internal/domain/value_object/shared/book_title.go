package shared_vo

type BookTitle string

func NewBookTitle(bookTitle string) BookTitle {
	return BookTitle(bookTitle)
}

func (bt BookTitle) Validate() bool {
	return len(bt) > 1 && len(bt) <= 50
}

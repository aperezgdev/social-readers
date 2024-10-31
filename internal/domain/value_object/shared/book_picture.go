package shared_vo

type BookPicture string

func NewBookPicture(bookPicture string) BookPicture {
	return BookPicture(bookPicture)
}

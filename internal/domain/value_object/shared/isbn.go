package shared_vo

type Isbn string

func NewIsbn(isbn string) Isbn {
	return Isbn(isbn)
}

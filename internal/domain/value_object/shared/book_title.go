package shared_vo

import (
	"strings"

	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
)

type BookTitle string

func NewBookTitle(bookTitle string) (BookTitle, error) {
	return BookTitle(bookTitle), validateBookTitle(bookTitle)
}

func validateBookTitle(bt string) error {
	title := strings.ReplaceAll(bt, " ", "")
	if len(title) < 1 || len(title) >= 50 {
		return errors.OutRange("Title", 1, 50)
	} 
	return nil
}

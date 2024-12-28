package shared_vo

import (
	"strconv"
	"strings"

	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
)

type Isbn string

func NewIsbn(isbn string) (Isbn, error) {
    if err := validateIsbn(isbn); err != nil {
        return "", err
    }
    return Isbn(isbn), nil
}

func validateIsbn(isbn string) error {
    isbn = strings.ReplaceAll(isbn, "-", "")
    if len(isbn) == 10 {
        return validateIsbn10(isbn)
    } else if len(isbn) == 13 {
        return validateIsbn13(isbn)
    }
    return errors.New("Isbn", "Invalid ISBN length")
}

func validateIsbn10(isbn string) error {
    if len(isbn) != 10 {
        return errors.New("Isbn","invalid ISBN-10 length")
    }
    sum := 0
    for i := 0; i < 9; i++ {
        num, err := strconv.Atoi(string(isbn[i]))
        if err != nil {
            return errors.New("Isbn","invalid character in ISBN-10")
        }
        sum += num * (10 - i)
    }
    checkDigit := isbn[9]
    if checkDigit == 'X' {
        sum += 10
    } else {
        num, err := strconv.Atoi(string(checkDigit))
        if err != nil {
            return errors.New("Isbn","invalid check digit in ISBN-10")
        }
        sum += num
    }
    if sum%11 != 0 {
        return errors.New("Isbn","invalid ISBN-10 checksum")
    }
    return nil
}

func validateIsbn13(isbn string) error {
    if len(isbn) != 13 {
        return errors.New("Isbn","invalid ISBN-13 length")
    }
    sum := 0
    for i := 0; i < 12; i++ {
        num, err := strconv.Atoi(string(isbn[i]))
        if err != nil {
            return errors.New("Isbn","invalid character in ISBN-13")
        }
        if i%2 == 0 {
            sum += num
        } else {
            sum += num * 3
        }
    }
    checkDigit, err := strconv.Atoi(string(isbn[12]))
    if err != nil {
        return errors.New("Isbn","invalid check digit in ISBN-13")
    }
    if (sum+checkDigit)%10 != 0 {
        return errors.New("Isbn","invalid ISBN-13 checksum")
    }
    return nil
}
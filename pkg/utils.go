package pkg

import (
	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/google/uuid"
)

func Map[T any, K any](array []T, fun func(t T) K) []K {
	var mappedArray []K

	for _, v := range array {
		mappedArray = append(mappedArray, fun(v))
	}

	return mappedArray
}

func Reduce[T any, K any](array []T, fun func(acc K, t T) K, init K) K {
	acc := init
	for _, v := range array {
		acc = fun(acc, v)
	}
	return acc
}

func ValidUUID(value, field string) error {
    _, err := uuid.Parse(value)
    if err != nil {
		return errors.InvalidUUID(field)
	}
	return nil
}
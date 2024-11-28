package pkg

func Map[T any, K any](array []T, fun func(t T) K) []K {
	var mappedArray []K

	for _, v := range array {
		mappedArray = append(mappedArray, fun(v))
	}

	return mappedArray
}

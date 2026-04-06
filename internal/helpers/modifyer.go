package helpers

func ToTakePointer[T any](run func(T) T) func(*T) {
	return func(t *T) {
		*t = run(*t)
	}
}

func Map[TIn any, TOut any](fn func(TIn) TOut, collection []TIn) []TOut {
	result := make([]TOut, 0, len(collection))
	for _, v := range collection {
		result = append(result, fn(v))
	}

	return result
}

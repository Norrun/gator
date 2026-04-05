package helpers

func ToTakePointer[T any](run func(T) T) func(*T) {
	return func(t *T) {
		*t = run(*t)
	}
}

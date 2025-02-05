package utils

func Filter[T any](s []T, f func(T) bool) (res []T) {
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return
}

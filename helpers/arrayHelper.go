package helpers

func Filter[T any](arr []T, test func(T) bool) (ret []T) {
	for _, s := range arr {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

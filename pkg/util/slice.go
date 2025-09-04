package util

func SliceMap[T any, M any](i []T, f func(highway T) M) []M {
	o := make([]M, len(i))

	for idx, _ := range i {
		o[idx] = f(i[idx])
	}

	return o
}

func SliceFilter[T any](i []T, f func(highway T) bool) []T {
	o := make([]T, 0)

	for idx, _ := range i {
		if f(i[idx]) {
			o = append(o, i[idx])
		}
	}

	return o
}

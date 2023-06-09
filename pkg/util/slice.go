package util

func SliceMap[T any, M any](i []T, f func(highway T) M) []M {
	o := make([]M, len(i))

	for idx, _ := range i {
		o[idx] = f(i[idx])
	}

	return o
}

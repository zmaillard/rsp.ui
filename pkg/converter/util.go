package converter

func getFromTo(prev []int32, next []int32) []int32 {
	if len(prev) != len(next) {
		return []int32{}
	}

	if len(prev) == 0 {
		return []int32{}
	}

	allFeatCount := len(prev) + 1
	out := make([]int32, allFeatCount)
	for i, _ := range prev {
		out[i] = prev[i]
	}

	out[allFeatCount-1] = next[len(next)-1]
	return out
}

package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	mArr := make([]int, max-min)
	for i := min; i < max; i++ {
		mArr[i-min] = i
	}

	return mArr
}

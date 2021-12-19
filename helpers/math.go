package helpers

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Sum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func Avg(arr []int) int {
	return Sum(arr) / len(arr)
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

package helpers

func ForEachNeighborInt(i, j, maxI, maxJ int, diag bool, f func(int, int)) {
	if i > 0 && j > 0 && diag {
		f(i-1, j-1)
	}
	if i > 0 {
		f(i-1, j)
	}
	if i > 0 && j+1 < maxJ && diag {
		f(i-1, j+1)
	}
	if j+1 < maxJ {
		f(i, j+1)
	}
	if i+1 < maxI && j+1 < maxJ && diag {
		f(i+1, j+1)
	}
	if i+1 < maxI {
		f(i+1, j)
	}
	if i+1 < maxI && j > 0 && diag {
		f(i+1, j-1)
	}
	if j > 0 {
		f(i, j-1)
	}
}

func InInt(v int, arr []int) bool {
	for _, v2 := range arr {
		if v == v2 {
			return true
		}
	}
	return false
}

func InStr(v string, arr []string) bool {
	for _, v2 := range arr {
		if v == v2 {
			return true
		}
	}
	return false
}

func CharInStr(v rune, s string) bool {
	for _, v2 := range s {
		if v == v2 {
			return true
		}
	}
	return false
}

func CountStr(v string, arr []string) int {
	total := 0
	for _, v2 := range arr {
		if v == v2 {
			total += 1
		}
	}
	return total
}

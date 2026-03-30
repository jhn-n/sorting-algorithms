package selection

func SortInt(a []int) {
	for i := 0; i < len(a)-1; i++ {
		j := findMinAbove(a, i)
		a[i], a[j] = a[j], a[i]
	}
}

func findMinAbove(a []int, i int) int {
	min_ind := i
	min_val := a[i]

	for j := i + 1; j < len(a); j++ {
		if a[j] < min_val {
			min_ind = j
			min_val = a[j]
		}
	}

	return min_ind
}

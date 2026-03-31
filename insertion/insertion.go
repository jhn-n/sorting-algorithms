package insertion

func SortInt(a []int) {
	for i, val := range a {
		insert(a, i, val)
	}
}

// maintain sorted list at start of array and insert accordingly
func insert(a []int, j, val int) {
	for j > 0 && val < a[j-1] {
		a[j] = a[j-1]
		j--
	}
	a[j] = val
}

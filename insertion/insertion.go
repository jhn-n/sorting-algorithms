package insertion

func SortInt(a []int) {
	for i, x := range a {
		insert(a, i, x)
	}
}

// maintain sorted list at start of array and insert
func insert(a []int, j, x int) {
	for j > 0 && x < a[j-1] {
		a[j] = a[j-1]
		j--
	}
	a[j] = x
}

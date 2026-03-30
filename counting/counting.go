package counting

func SortInt(a []int) {
	// first find max value
	max_int := 0
	for _, x := range a {
		if x > max_int {
			max_int = x
		}
	}
	// use array slice to count entities
	counts := make([]int, max_int+1)
	for _, x := range a {
		counts[x] += 1
	}

	// write ordered values back to source
	val := 0
	for i := 0; i < len(a); {
		for range counts[val] {
			a[i] = val
			i++
		}
		val++
	}

}

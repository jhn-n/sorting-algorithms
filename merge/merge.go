package merge

var insertThreshold = 0

func SortInt2(a []int) {
	insertThreshold = 8
	SortInt(a)
}

func SortInt(a []int) {
	temp := make([]int, len(a))
	mergeRecur(a, temp, 0, len(a))
}

func mergeRecur(a, temp []int, low, high int) {
	if high-low <= 1 {
		return
	}

	if high-low <= insertThreshold {
		insertSort(a, low, high)
		return
	}

	mid := (low + high) / 2
	mergeRecur(a, temp, low, mid)
	mergeRecur(a, temp, mid, high)
	mergeArrays(a, temp, low, mid, high)
}

func mergeArrays(a, temp []int, low, mid, high int) {
	x := low // first array index
	y := mid // second array index
	i := low // temp index to copy into

	// merge the two arrays by in order
	for x < mid && y < high {
		if a[x] <= a[y] {
			temp[i] = a[x]
			x++
		} else {
			temp[i] = a[y]
			y++
		}
		i++
	}

	// copy everything that's left in at most one of the two half-lists
	for x < mid {
		temp[i] = a[x]
		x++
		i++
	}

	for y < high {
		temp[i] = a[y]
		y++
		i++
	}

	// copy back from temp to main array
	for i := low; i < high; i++ {
		a[i] = temp[i]
	}
}

func insertSort(a []int, low, high int) {
	for i := low; i < high; i++ {
		j := i
		x := a[i]
		for j > low && x < a[j-1] {
			a[j] = a[j-1]
			j--
		}
		a[j] = x
	}
}

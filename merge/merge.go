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

func mergeRecur(a, temp []int, lowInd, highInd int) {
	if highInd-lowInd <= 1 {
		return
	}

	if highInd-lowInd <= insertThreshold {
		insertSort(a, lowInd, highInd)
		return
	}

	midInd := (lowInd + highInd) / 2
	mergeRecur(a, temp, lowInd, midInd)
	mergeRecur(a, temp, midInd, highInd)
	mergeArrays(a, temp, lowInd, midInd, highInd)
}

func mergeArrays(a, temp []int, lowInd, midInd, highInd int) {
	x := lowInd // first array index
	y := midInd // second array index
	i := lowInd // temp index to copy into

	// merge the two arrays by in order
	for x < midInd && y < highInd {
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
	for x < midInd {
		temp[i] = a[x]
		x++
		i++
	}

	for y < highInd {
		temp[i] = a[y]
		y++
		i++
	}

	// copy back from temp to main array
	for i := lowInd; i < highInd; i++ {
		a[i] = temp[i]
	}
}

func insertSort(a []int, lowInd, highInd int) {
	for i := lowInd; i < highInd; i++ {
		j := i
		val := a[i]
		for j > lowInd && val < a[j-1] {
			a[j] = a[j-1]
			j--
		}
		a[j] = val
	}
}

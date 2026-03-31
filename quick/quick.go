package quick

var pivotChoice = false

func SortInt2(a []int) {
	pivotChoice = true
	quicksort(a, 0, len(a))
}

func SortInt(a []int) {
	quicksort(a, 0, len(a))
}

func quicksort(a []int, lowInd, highInd int) {

	// Step 0: check for completion
	if lowInd >= highInd {
		return
	}

	// Step 1: select pivot
	if pivotChoice {
		swapPivot(a, lowInd, highInd)
	}
	pivot := a[lowInd]

	// Step 2: partition items to lower or upper portions of the temp list
	pivotX := partition(a, lowInd, highInd, pivot)

	// Step 3: sort both halves of the list
	quicksort(a, lowInd, pivotX)
	quicksort(a, pivotX+1, highInd)
}

func swapPivot(a []int, lowInd, highInd int) {
	if highInd-lowInd < 3 {
		return
	}
	midInd := (lowInd + highInd) / 2

	// find median of three indices
	highInd--
	if a[lowInd] < a[midInd] {
		if a[midInd] <= a[highInd] {
			a[lowInd], a[midInd] = a[midInd], a[lowInd]
		} else if a[lowInd] < a[highInd] {
			a[lowInd], a[highInd] = a[highInd], a[lowInd]
		} else {
		}
	} else {
		if a[lowInd] <= a[highInd] {
		} else if a[midInd] < a[highInd] {
			a[lowInd], a[highInd] = a[highInd], a[lowInd]
		} else {
			a[lowInd], a[midInd] = a[midInd], a[lowInd]
		}
	}
}

func partition(a []int, lowInd, highInd, pivot int) int {
	x := lowInd
	y := highInd - 1
	for x < y {
		// Advance x until we find a too-big value (or overshoot the end of the list)
		for x < highInd && a[x] <= pivot {
			x++
		}
		// Advance y (backwards) until we find a too-small value (or undershoot start)
		for y >= lowInd && a[y] > pivot {
			y--
		}
		if x < y {
			a[x], a[y] = a[y], a[x]
		}
	}
	// place pivot in remaining spot
	a[lowInd], a[y] = a[y], a[lowInd]
	return y
}

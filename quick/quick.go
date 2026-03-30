package quick

var pivotChoice = false

func SortInt2(a []int) {
	pivotChoice = true
	quicksort(a, 0, len(a))
}

func SortInt(a []int) {
	quicksort(a, 0, len(a))
}

func quicksort(a []int, lowX, highX int) {

	// Step 0: check for completion
	if lowX >= highX {
		return
	}

	// Step 1: select pivot
	if pivotChoice {
		swapPivot(a, lowX, highX)
	}
	pivot := a[lowX]

	// Step 2: partition items to lower or upper portions of the temp list
	pivotX := partition(a, lowX, highX, pivot)

	// Step 3: sort both halves of the list
	quicksort(a, lowX, pivotX)
	quicksort(a, pivotX+1, highX)
}

func swapPivot(a []int, low, high int) {
	if high-low < 3 {
		return
	}
	mid := (low + high) / 2

	// find median of three indices
	high--
	if a[low] < a[mid] {
		if a[mid] <= a[high] {
			a[low], a[mid] = a[mid], a[low]
		} else if a[low] < a[high] {
			a[low], a[high] = a[high], a[low]
		} else {
		}
	} else {
		if a[low] <= a[high] {
		} else if a[mid] < a[high] {
			a[low], a[high] = a[high], a[low]
		} else {
			a[low], a[mid] = a[mid], a[low]
		}
	}
}

func partition(a []int, lowX, highX, pivot int) int {
	x := lowX
	y := highX - 1
	for x < y {
		// Advance x until we find a too-big value (or overshoot the end of the list)
		for x < highX && a[x] <= pivot {
			x++
		}
		// Advance y (backwards) until we find a too-small value (or undershoot start)
		for y >= lowX && a[y] > pivot {
			y--
		}
		if x < y {
			a[x], a[y] = a[y], a[x]
		}
	}
	// place pivot in remaining spot
	a[lowX], a[y] = a[y], a[lowX]
	return y
}

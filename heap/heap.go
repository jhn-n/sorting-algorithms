package heap

func SortInt(a []int) {
	// setup heap (insert at end and bubble up)
	for i := 1; i < len(a); i++ {
		val := a[i]
		insertInd := i

		for insertInd > 0 {
			parentInd := (insertInd - 1) / 2
			if a[parentInd] >= val {
				break
			}
			a[insertInd] = a[parentInd]
			insertInd = parentInd
		}
		a[insertInd] = val
	}

	// swap max element with end element and rebalance (bubble down)
	for i := len(a) - 1; i > 0; i-- {
		val := a[i]
		a[i] = a[0]
		destInd := 0

		for {
			child1Ind := destInd*2 + 1
			child2Ind := child1Ind + 1

			if child2Ind < i && a[child2Ind] > val && a[child2Ind] > a[child1Ind] {
				a[destInd] = a[child2Ind]
				destInd = child2Ind
			} else if child1Ind < i && a[child1Ind] > val {
				a[destInd] = a[child1Ind]
				destInd = child1Ind
			} else {
				break
			}
		}
		a[destInd] = val
	}
}

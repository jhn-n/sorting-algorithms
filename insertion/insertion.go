package insertion

func Sort(a []int) []int {
	b := make([]int, len(a))

	for i, x := range a {
		j := len(a) - 1 - i
		for j < len(b)-1 && x > b[j+1] {
			b[j] = b[j+1]
			j++
		}
		b[j] = x
	}
	return b
}

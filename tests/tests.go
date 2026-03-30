package main

import (
	"fmt"
	"math"
	"math/rand"
	"sorting/counting"
	"sorting/insertion"
	"sorting/merge"
	"sorting/quick"
	"sorting/selection"
	"time"
)

func main() {
	alg := map[string]func([]int){
		"i":  insertion.SortInt,
		"s":  selection.SortInt,
		"m":  merge.SortInt,
		"m2": merge.SortInt2,
		"c":  counting.SortInt,
		"q":  quick.SortInt,
		"q2": quick.SortInt2}

	// testAlg(alg["q"], 10)
	compareAlg(alg["q2"], alg["m2"], 1)
}

func numIts(n int) int {
	return min(1000, 5e6/(n*int(math.Log2(float64(n)))))
}

func testAlg(alg func([]int), maxIntMultiplier int) {
	n := 5
	its := numIts(n)
	var prior_av float64

	for i := range 16 {
		fmt.Printf("n:%-6v its:%-6v ", n, its)
		durations := make([]float64, its)
		for i := range its {
			unsorted := randomArray(n, maxIntMultiplier*n)
			durations[i] = timeSort(alg, unsorted)
		}
		fastest, slowest, average := fastSlowAv(durations)
		fmt.Printf("av:%.2ems [%.2e, %.2e]", average, fastest, slowest)
		if i > 0 {
			fmt.Printf(" x %.3v", average/prior_av)
		}
		fmt.Println()
		n *= 2
		its = numIts(n)
		prior_av = average
	}
	fmt.Println()
}

func compareAlg(alg1, alg2 func([]int), max_mult int) {
	fmt.Println("Relative speed of second algorithm compared with first")
	n := 5
	its := numIts(n)
	for range 16 {
		fmt.Printf("n:%-6v its:%-6v ", n, its)
		durations1 := make([]float64, its)
		durations2 := make([]float64, its)
		for i := range its {
			unsorted1 := randomArray(n, max_mult*n)
			unsorted2 := make([]int, len(unsorted1))
			copy(unsorted2, unsorted1)
			durations1[i] = timeSort(alg1, unsorted1)
			durations2[i] = timeSort(alg2, unsorted2)
		}
		_, _, av1 := fastSlowAv(durations1)
		_, _, av2 := fastSlowAv(durations2)

		fmt.Printf("speedup x %.3v\n", av1/av2)
		n *= 2
		its = numIts(n)
	}
}

func randomArray(size, max int) []int {
	a := make([]int, size)
	for i := range size {
		a[i] = rand.Intn(max) + 1
	}
	return a
}

func timeSort(alg func([]int), array []int) float64 {
	// return time in milliseconds
	oldValues := recordValues(array)
	start := time.Now()
	alg(array)
	elapsed := time.Since(start)
	if !inOrder(array) {
		panic("Uh oh, incorrect sort!")
	}
	if !sameData(array, oldValues) {
		panic("Uh oh, different values!")
	}
	return float64(elapsed.Nanoseconds()) / 1e6
}

func recordValues(a []int) map[int]int {
	record := make(map[int]int)
	for _, x := range a {
		record[x] = record[x] + 1
	}
	return record
}

func sameData(a []int, oldValues map[int]int) bool {
	newValues := recordValues(a)
	for k := range oldValues {
		if oldValues[k] != newValues[k] {
			return false
		}
	}
	for k := range newValues {
		if oldValues[k] != newValues[k] {
			return false
		}
	}
	return true
}

func inOrder(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func fastSlowAv(data []float64) (float64, float64, float64) {
	fast := data[0]
	slow := data[0]
	sum := 0.0
	for _, x := range data {
		if x < fast {
			fast = x
		}
		if x > slow {
			slow = x
		}
		sum += x
	}
	return fast, slow, sum / float64(len(data))
}

package main

import (
	"fmt"
	"math/rand"
	"sorting/insertion"
	"time"
)

func main() {
	testAlg(insertion.Sort)
}

func randomSlice(size, max int) []int {
	a := make([]int, size)
	for i := range size {
		a[i] = rand.Intn(max + 1)
	}
	return a
}

func inOrder(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func testAlg(alg func([]int) []int) {
	n := 100
	for range 11 {
		fmt.Printf("n = %v, ", n)
		testSort(alg, n, n, 10)
		n *= 2
	}
}

func testSort(alg func([]int) []int, size, max, iters int) {
	durations := make([]int, iters)
	for i := range iters {
		unsorted := randomSlice(size, max)
		durations[i] = timeSort(alg, unsorted)
	}
	summarise(durations)
}

func timeSort(alg func([]int) []int, unsorted []int) int {
	start := time.Now()
	sorted := alg(unsorted)
	elapsed := time.Since(start)
	if !inOrder(sorted) {
		panic("Uh oh, incorrect sort")
	}
	return int(elapsed.Microseconds())
}

func summarise(data []int) {
	fastest := data[0]
	slowest := data[0]
	sum := 0
	for _, x := range data {
		if x < fastest {
			fastest = x
		} else if x > slowest {
			slowest = x
		}
		sum += x
	}
	average := sum / len(data)
	fmt.Printf("fastest: %v, slowest: %v, average: %v\n", fastest, slowest, average)
}

package main

import (
	"fmt"
)

func countNumberPairs(numbers []int) map[[2]int]int {
	pairCount := make(map[[2]int]int)

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := [2]int{numbers[i], numbers[j]}
			pairCount[pair]++
		}
	}

	return pairCount
}

func main() {
	numberSlice := []int{1, 2, 3, 2, 4, 3, 5, 1, 3, 4}

	pairFrequencies := countNumberPairs(numberSlice)

	fmt.Println("Frequência de pares de números:")
	for pair, frequency := range pairFrequencies {
		fmt.Printf("%v: %d\n", pair, frequency)
	}
}

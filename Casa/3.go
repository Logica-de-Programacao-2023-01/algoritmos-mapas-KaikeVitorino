package main

import (
	"fmt"
)

func sumValues(m map[string]int) int {
	sum := 0

	for _, value := range m {
		sum += value
	}

	return sum
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	total := sumValues(m)

	fmt.Printf("Soma dos valores: %d\n", total)
}

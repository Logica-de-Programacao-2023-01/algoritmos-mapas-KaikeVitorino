package main

import (
	"fmt"
)

func getCharacterFrequency(text string) map[rune]int {
	frequency := make(map[rune]int)

	for _, char := range text {
		frequency[char]++
	}

	return frequency
}

func main() {
	text := "abracadabra"

	charFrequency := getCharacterFrequency(text)

	fmt.Println("Frequência dos caracteres:")
	for char, count := range charFrequency {
		fmt.Printf("%c: %d\n", char, count)
	}
}

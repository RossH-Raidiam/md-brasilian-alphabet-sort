package main

import (
	"fmt"
)

//TODO: Implement further, can extend as needed
var _map = map[rune]int{
	'a': 1,
	'A': 1,
	'Á': 2,
	'á': 2,
	'b': 3,
	'B': 3,
	'c': 4,
	'C': 4,
	'd': 5,
	'D': 5,
	'e': 6,
	'E': 6,
	'f': 7,
	'F': 7,
	'g': 8,
	'G': 8,
	'h': 9,
	'H': 9,
	'j': 10,
	'J': 10,
	'k': 11,
	'K': 11,
	'l': 12,
	'L': 12,
	'm': 13,
	'M': 13,
	'n': 14,
	'N': 14,
	'o': 15,
	'O': 15,
	'Ô': 16,
	'p': 17,
	'P': 17,
	'q': 18,
	'Q': 18,
	'r': 19,
	'R': 19,
	's': 20,
	'S': 20,
	't': 21,
	'T': 21,
	'u': 22,
	'U': 22,
	'ú': 23,
	'v': 24,
	'V': 24,
	'w': 25,
	'W': 25,
	'x': 26,
	'X': 26,
	'y': 27,
	'Y': 27,
	'z': 28,
	'Z': 28,
}

func main() {
	test := "test string"

	// For testing
	for _, char := range test {
		fmt.Println(_map[char])
	}

}

//TODO: implement an alphabet sort using the map

package brazil_merge_sort

//Can extend this as needed
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

func MergeSort(strings []string) []string {
	if len(strings) == 1 {
		return strings
	}

	var midpoint int
	if len(strings)%2 == 0 {
		midpoint = len(strings) / 2
	} else {
		midpoint = (len(strings) + 1) / 2
	}

	var left []string = strings[:midpoint]
	var right []string = strings[midpoint:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []string, right []string) []string {
	var result []string
	var index_left int = 0
	var index_right int = 0

	for index_left < len(left) && index_right < len(right) {
		switch evalStrings(left[index_left], right[index_right]) {
		case 1:
			result = append(result, left[index_left])
			index_left = index_left + 1
		case 2:
			result = append(result, right[index_right])
			index_right = index_right + 1
		}
	}

	left_slice := left[index_left:]
	right_slice := right[index_right:]
	result = append(result, left_slice...)
	result = append(result, right_slice...)

	return result
}

func evalStrings(first string, second string) int {
	var a, b = []rune(first), []rune(second)
	for i, char := range a {
		switch {
		// first string larger, return second
		case _map[char] > _map[b[i]]:
			return 2
		// second string larger, return first
		case _map[char] < _map[b[i]]:
			return 1
		default:
		}
	}
	return 0
}

package mdbrasilianalphabetsort

//Can extend this as needed
var _map = GenerateMap()

func MergeSort(strings []string) []string {
	if len(strings) < 2 {
		return strings
	}

	var left []string = MergeSort(strings[:len(strings)/2])
	var right []string = MergeSort(strings[len(strings)/2:])

	return merge(left, right)
}

func merge(left []string, right []string) []string {
	result := []string{}
	index_left := 0
	index_right := 0

	for index_left < len(left) && index_right < len(right) {
		switch evalStrings(left[index_left], right[index_right]) {
		case 1:
			result = append(result, left[index_left])
			index_left++
		default:
			result = append(result, right[index_right])
			index_right++
		}
	}
	for index_left < len(left) {
		result = append(result, left[index_left])
		index_left++
	}
	for index_right < len(right) {
		result = append(result, right[index_right])
		index_right++
	}
	return result
}

func evalStrings(first string, second string) int {

	var a, b = []rune(first), []rune(second)
	for i, char := range a {
		switch {
		// first string larger, return second
		case i == len(a) && i == len(b):
			return 0
		case i == len(a):
			return 1
		case i == len(b):
			return 2
		case _map[char] > _map[b[i]]:
			return 2
		// second string larger, return first
		case _map[char] < _map[b[i]]:
			return 1
		default:
			continue
		}
	}
	return 0
}

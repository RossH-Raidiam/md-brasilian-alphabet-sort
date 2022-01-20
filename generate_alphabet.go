package mdbrasilianalphabetsort

func GenerateMap() map[rune]int {
	var _map = map[rune]int{}
	alphabet := [][]rune{
		{' '},
		{'a', 'A'},
		{'á', 'Á'},
		{'b', 'B'},
		{'c', 'C'},
		{'d', 'D'},
		{'e', 'E'},
		{'f', 'F'},
		{'g', 'G'},
		{'h', 'H'},
		{'i', 'I'},
		{'j', 'J'},
		{'k', 'K'},
		{'l', 'L'},
		{'m', 'M'},
		{'n', 'N'},
		{'o', 'O'},
		{'Ô'},
		{'p', 'P'},
		{'q', 'Q'},
		{'r', 'R'},
		{'s', 'S'},
		{'t', 'T'},
		{'u', 'U'},
		{'ú'},
		{'v', 'V'},
		{'w', 'W'},
		{'x', 'X'},
		{'y', 'Y'},
		{'z', 'Z'},
	}

	for i, outer := range alphabet {
		for _, inner := range outer {
			_map[inner] = i
		}
	}

	return _map
}

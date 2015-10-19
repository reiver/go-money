package money


import (
	"fmt"
)


func readable(n int64, sep rune) string {

	s := fmt.Sprintf("%d", n)

	runes := []rune(s)
	if 3 >= len(runes) {
		return s
	}

	result := []rune{}

	initialSep := false
	switch remainder := len(runes) % 3; remainder {
		case 1:
			result = append(result, runes[0])
			runes = runes[1:]
			initialSep = true
		case 2:
			result = append(result, runes[0], runes[1])
			runes = runes[2:]
			initialSep = true
	}

	first := true
	for 0 < len(runes) {
		if first && initialSep  {
			result = append(result, sep)
		}
		if !first {
			result = append(result, sep)
		}

		result = append(result, runes[0], runes[1], runes[2])

		runes = runes[3:]

		first = false
	}


	return fmt.Sprintf("%s", string(result))
}

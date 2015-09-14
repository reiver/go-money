package money


import (
	"errors"
	"strconv"
	"strings"
)


// ParseCAD converts a string to a CAD.
func ParseCAD(str string) (CAD, error) {
	cents, err := parseDollarsAndCents(str)
	return CAD(cents), err
}


// ParseCAD converts a string to a USD.
func ParseUSD(str string) (USD, error) {
	cents, err := parseDollarsAndCents(str)
	return USD(cents), err
}


func parseDollarsAndCents(str string) (int64, error) {

	// Trim leading and ending whitespace.
	str = strings.TrimSpace(str)

	// If resulting string is empty, return an error.
	if "" == str {
//@TODO: Make better error.
		return 0, errors.New("Nothing in the string.")
	}

	// Convert to runes.
	runes := []rune(str)

	// Check if it begins with a dollar symbol
	// or ends with a cents symbol.
	hasDollarSymbol := false
	hasCentsSymbol := false

	switch runes[0] {
	case '$',     // normal dollar sign
	'\uFE69',     // small dollar sign
	'\uFF04',     // full-width dollar sign
	'\U0001F4B2': // heavy dollar sign
		hasDollarSymbol = true
	}
	if !hasDollarSymbol {
		switch runes[0] {
		case '\u00A2', // cent sign
		     '\uFFE0': // full-width cent sign
			hasCentsSymbol = true
		}
	}


	switch {
	case hasDollarSymbol:
		return parseDollars(runes[1:])
	case hasCentsSymbol:
		return parseCents(runes[:len(runes)-1])
	default:
//@TODO: Make better error.
		return 0, errors.New("Parse error.")
	}
}

func parseDollars(runes []rune) (int64, error) {

	// Money is stored in cents not dollars.
	// So if the string (which we converted to []rune elsewhere)
	// has a decimal point (".") in it (ex: "12.34") then we
	// deal with that situation.
	//
	// So, for example:
	//
	//	"124.45" -> "12345"
	//
	//	"12" -> "1200"
	//
	//	"0.10" -> "010"
	if hasDecimalPoint := 3 <= len(runes) && '.' == runes[len(runes)-3]; hasDecimalPoint {

		i := len(runes)-3

		runes = append(runes[:i], runes[i+1:]...)
	} else {

		runes = append(runes, '0', '0')
	}


	// Return.
	return parseCents(runes)
}

func parseCents(runes []rune) (int64, error) {

	// Convert the string into an int64.
	// And then convert that int64 into money.
	const base = 10
	const bitSize = 64
	if i, err := strconv.ParseInt(string(runes), base, bitSize); nil != err {
		return 0, err
	} else {
		return i, nil
	}

}

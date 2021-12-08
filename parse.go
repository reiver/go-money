package money

import (
	"github.com/reiver/go-numeric"

	"errors"
	"fmt"
	"strconv"
	"strings"
)

// MustParseCAD converts a string to a CAD, but panic()s if there is an error.
func MustParseCAD(str string) CAD {
	cad, err := ParseCAD(str)
	if nil != err {
		panic(err)
	}

	return cad
}

// MustParseUSD converts a string to a USD, but panic()s if there is an error.
func MustParseUSD(str string) USD {
	usd, err := ParseUSD(str)
	if nil != err {
		panic(err)
	}

	return usd
}

// ParseCAD converts a string to a CAD.
func ParseCAD(str string) (CAD, error) {
	cents, err := parseDollarsAndCents(str)
	if nil != err {
		return nothingCAD(), err
	}

	return CADCents(cents), nil
}


// ParseUSD converts a string to a USD.
func ParseUSD(str string) (USD, error) {
	cents, err := parseDollarsAndCents(str)
	if nil != err {
		return nothingUSD(), err
	}

	return USDCents(cents), nil
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

	// Check if it begins with a negative symbol.
	// (I.e., a hyphen.)
	//
	// If it does have a negative symbol we
	// truncate the rune slice by removing
	// the first character (i.e., removing
	// the negative symbol).
	hasNegativeSymbol := false
	if '-' == runes[0] {
		hasNegativeSymbol = true
		runes = runes[1:]
	}
	if 1 > len(runes) {
//@TODO: Make better error.
		return 0, fmt.Errorf("Parse error on %q.", str)
	}

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
		runes = runes[1:]
	}
	if !hasDollarSymbol {
		switch runes[0] {
		case '\u00A2', // cent sign
		     '\uFFE0': // full-width cent sign
			hasCentsSymbol = true
			runes = runes[:len(runes)-1]
		}
	}
	if 1 > len(runes) {
//@TODO: Make better error.
		return 0, fmt.Errorf("Parse error on %q.", str)
	}


//@TODO: This is a bit hacky. This was added to make it so could parse things
//       like "84.27". I.e., no dollar or cent symbol.
//       Make this better later.
	if !hasDollarSymbol && !hasCentsSymbol {
		hasDollarSymbol = true
		for i:=0; i<len(runes); i++ {
			if r := runes[i]; !numeric.IsNumeric(r) && '.' != r && ',' != r {
				hasDollarSymbol = false
			}
		}
	}

	switch {
	case hasDollarSymbol:
		cents, err := parseDollars(runes)
		if hasNegativeSymbol {
			cents = -1 * cents
		}
		return cents, err
	case hasCentsSymbol:
		cents, err := parseCents(runes)
		if hasNegativeSymbol {
			cents = -1 * cents
		}
		return cents, err
	default:
//@TODO: Make better error.
		return 0, fmt.Errorf("Parse error on %q.", str)
	}
}

func parseDollars(runes []rune) (int64, error) {

	// Money is stored in cents not dollars.
	// So if the string (which we converted to []rune elsewhere)
	// has a decimal point (either "." or ",") in it then we deal
	// with that situation.
	//
	// So, for example:
	//
	//	"12345.67" -> "1234567"
	//
	//	"12345,67" -> "1234567"
	//
	//	"12,345.67" -> "12,34567"
	//
	//	"12.345,67" -> "12.34567"
	//
	//	"12345" -> "1234500"
	//
	//	"0.10" -> "010"
	//
	//	"0,10" -> "010"
    //
    //  "0,1"  -> "010"
	//
	// After that then if the decimal point is a dot (".") then we remove comma (",") separators.
	// Alternatively, if the decimal point is a comma (",") then we remove dot (".") separators.
	//
	// So, for example:
	//
	//	"12345.67" -> "1234567" -> "1234567"
	//
	//	"12345,67" -> "1234567" -> "1234567"
	//
	//	"12,345.67" -> "12,34567" -> "1234567"
	//
	//	"12.345,67" -> "12.34567" -> "12.34567"
	//
	//	"12345" -> "1234500" -> "1234500"
	//
	//	"0.10" -> "010" -> "010"
	//
	//	"0,10" -> "010" -> "010"
	//
	//  "0.1" ->  "010"  -> "010"

	if hasSingleDecimalPlace := 2 <= len(runes) && ('.' == runes[len(runes)-2] || ',' == runes[len(runes)-2]); hasSingleDecimalPlace {
		runes = append(runes, '0')
	}

	if hasDotDecimalPoint := 3 <= len(runes) && '.' == runes[len(runes)-3]; hasDotDecimalPoint {

		i := len(runes)-3

		runes = append(runes[:i], runes[i+1:]...)

//@TODO: This is not the proper way to remove the separator ",".
		for ii:=0; ii<len(runes); ii++ {

			if ',' == runes[ii] {
				runes = append(runes[:ii], runes[1+ii:]...)
			}
		}
	} else if hasCommaDecimalPoint := 3 <= len(runes) && ',' == runes[len(runes)-3]; hasCommaDecimalPoint {

		i := len(runes)-3

		runes = append(runes[:i], runes[i+1:]...)

//@TODO: This is not the proper way to remove the separator ".".
		for ii:=0; ii<len(runes); ii++ {

			if '.' == runes[ii] {
				runes = append(runes[:ii], runes[1+ii:]...)
			}
		}
	} else {

		if 4 <= len(runes) && ',' == runes[len(runes)-4] {
//@TODO: This is not the proper way to remove the separator ",".
			for ii:=0; ii<len(runes); ii++ {

				if ',' == runes[ii] {
					runes = append(runes[:ii], runes[1+ii:]...)
				}
			}

		} else if 4 <= len(runes) && '.' == runes[len(runes)-4] {
//@TODO: This is not the proper way to remove the separator ".".
			for ii:=0; ii<len(runes); ii++ {

				if '.' == runes[ii] {
					runes = append(runes[:ii], runes[1+ii:]...)
				}
			}
		}

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

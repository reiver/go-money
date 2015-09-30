package money


import (
	"fmt"
)


// Format returns a textual representation of the CAD money value formatted according to layout.
//
// Options for layout include:
//
// "$"
//	$123.45
//
// "\uFE69" or "﹩"
//	﹩123.45
//
// "\uFF04" or "＄"
//	＄123.45
//
// "\U0001F4B2" or "💲"
//	💲123.45
//
// "\u00A2" or "¢"
//      12345¢
//
// "\uFFE0" or "￠"
//      12345￠
func (m CAD) Format(layout string) string {
	cents := int64(m)
	return formatDollarsAndCents(cents, layout)
}

// Format returns a textual representation of the USD money value formatted according to layout.
//
// Options for layout include:
//
// "$"
//	$123.45
//
// "\uFE69" or "﹩"
//	﹩123.45
//
// "\uFF04" or "＄"
//	＄123.45
//
// "\U0001F4B2" or "💲"
//	💲123.45
//
// "\u00A2" or "¢"
//      12345¢
//
// "\uFFE0" or "￠"
//      12345￠

func (m USD) Format(layout string) string {
	cents := int64(m)
	return formatDollarsAndCents(cents, layout)
}


func formatDollarsAndCents(cents int64, layout string) string {

	maybeNegative := ""
	if 0 > cents {
		maybeNegative = "-"
		cents = -1 * cents // We make it positive.
	}

	centsPart := int64(cents) % 100
	dollarsPart := (int64(cents) - centsPart) / 100

	switch layout {
	default:
		return fmt.Sprintf("%s$%d.%02d",          maybeNegative, dollarsPart, centsPart)
	case "$":          // normal dollar sign
		return fmt.Sprintf("%s$%d.%02d",          maybeNegative, dollarsPart, centsPart)
	case "\uFE69":     // small dollar sign
		return fmt.Sprintf("%s\uFE69%d.%02d",     maybeNegative, dollarsPart, centsPart)
	case "\uFF04":     // full-width dollar sign
		return fmt.Sprintf("%s\uFF04%d.%02d",     maybeNegative, dollarsPart, centsPart)
	case "\U0001F4B2": // heavy dollar sign
		return fmt.Sprintf("%s\U0001F4B2%d.%02d", maybeNegative, dollarsPart, centsPart)
	case "\u00A2": // cent sign
		return fmt.Sprintf("%s%d\u00A2", maybeNegative, cents)
	case "\uFFE0": // full-width cent sign
		return fmt.Sprintf("%s%d\uFFE0", maybeNegative, cents)
	}
}


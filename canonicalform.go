package money

// CADCanonicalForm takes a dollars, and cents amount and returns a money.CAD with that amount.
//
// For example:
//
//	var value money.CAD = money.CADCanonicalForm(123, 45)
//
// This will return $123.45
func CADCanonicalForm(dollars int32, cents int8) CAD {
	var incents int64 = (int64(dollars) * 100) + int64(cents)

	return CADCents(incents)
}

// USDCanonicalForm takes a dollars, and cents amount and returns a money.USD with that amount.
//
// For example:
//
//	var value money.USD = money.USDCanonicalForm(123, 45)
//
// This will return $123.45
func USDCanonicalForm(dollars int32, cents int8) USD {
	var incents int64 = (int64(dollars) * 100) + int64(cents)

	return USDCents(incents)
}

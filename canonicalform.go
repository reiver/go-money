package money

func CADCanonicalForm(dollars int32, cents int8) CAD {
	var incents int64 = (int64(dollars) * 100) + int64(cents)

	return CADCents(incents)
}

func USDCanonicalForm(dollars int32, cents int8) USD {
	var incents int64 = (int64(dollars) * 100) + int64(cents)

	return USDCents(incents)
}

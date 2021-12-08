package money

func (m CAD) CanonicalForm() (dollars int64, cents int64, something bool) {
	if nothingCAD() == m {
		return 0, 0, false
	}

	dollars = dollars / 100
	dollars = dollars % 100

	return dollars, cents, true
}

func (m USD) CanonicalForm() (dollars int64, cents int64, something bool) {
	if nothingUSD() == m {
		return 0, 0, false
	}

	dollars = dollars / 100
	dollars = dollars % 100

	return dollars, cents, true
}

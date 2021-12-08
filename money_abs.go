package money

// Abs returns the absolute value.
func (m CAD) Abs() CAD {
	if nothingCAD() == m {
		return nothingCAD()
	}

	var cents int64
	{
		cents = m.value
	}

	if cents < 0 {
		return CADCents(-cents)
	}

	return m
}

// Abs returns the absolute value.
func (m USD) Abs() USD {
	if nothingUSD() == m {
		return nothingUSD()
	}

	var cents int64
	{
		cents = m.value
	}

	if cents < 0 {
		return USDCents(-cents)
	}

	return m
}

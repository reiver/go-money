package money

func (m CAD) Mul(scalar int64) CAD {
	if nothingCAD() == m {
		return nothingCAD()
	}

	var cents int64 = m.value

	var result int64 = cents * scalar

	return CADCents(result)
}

func (m USD) Mul(scalar int64) USD {
	if nothingUSD() == m {
		return nothingUSD()
	}

	var cents int64 = m.value

	var result int64 = cents * scalar

	return USDCents(result)
}

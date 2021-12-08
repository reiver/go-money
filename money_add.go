package money

func (m CAD) Add(val CAD) CAD {
	if nothingCAD() == m {
		return nothingCAD()
	}
	if nothingCAD() == val {
		return nothingCAD()
	}

	var cents1 int64 = m.value
	var cents2 int64 = val.value

	var result = cents1 + cents2

	return CADCents(result)
}

func (m USD) Add(val USD) USD {
	if nothingUSD() == m {
		return nothingUSD()
	}
	if nothingUSD() == val {
		return nothingUSD()
	}

	var cents1 int64 = m.value
	var cents2 int64 = val.value

	var result = cents1 + cents2

	return USDCents(result)
}

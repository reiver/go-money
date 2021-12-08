package money

func (m CAD) Sub(other CAD) CAD {
	if nothingCAD() == m {
		return nothingCAD()
	}
	if nothingCAD() == other {
		return nothingCAD()
	}

	var cents1 int64 = m.value
	var cents2 int64 = other.value

	var result int64 = cents1 - cents2

	return CADCents(result)
}

func (m USD) Sub(other USD) USD {
	if nothingUSD() == m {
		return nothingUSD()
	}
	if nothingUSD() == other {
		return nothingUSD()
	}

	var cents1 int64 = m.value
	var cents2 int64 = other.value

	var result int64 = cents1 - cents2

	return USDCents(result)
}

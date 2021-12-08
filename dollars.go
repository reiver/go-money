package money

func CADDollars(value int64) CAD {
	return CADCents(value * 100)
}

func USDDollars(value int64) USD {
	return USDCents(value * 100)
}

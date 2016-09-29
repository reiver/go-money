package money

func (m CAD) MulCAD(val int64) CAD {

	return CAD(int64(m) * val)
}

func (m USD) MulUSD(val int64) USD {

	return USD(int64(m) * val)
}

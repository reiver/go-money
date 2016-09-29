package money

func (m CAD) Mul(val int64) CAD {

	return CAD(int64(m) * val)
}

func (m USD) Mul(val int64) USD {

	return USD(int64(m) * val)
}

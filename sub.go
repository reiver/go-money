package money

func (m CAD) SubCAD(m1 CAD) CAD {
	return m - m1
}

func (m USD) SubUSD(m1 USD) USD {
	return m - m1
}

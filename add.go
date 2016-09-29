package money

func (m CAD) AddCAD(m1 CAD) CAD {
	return m + m1
}

func (m USD) AddUSD(m1 USD) USD {
	return m + m1
}

package money

func (m CAD) Sub(m1 CAD) CAD {
	return m - m1
}

func (m USD) Sub(m1 USD) USD {
	return m - m1
}

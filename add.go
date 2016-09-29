package money

func (m CAD) Add(m1 CAD) CAD {
	return m + m1
}

func (m USD) Add(m1 USD) USD {
	return m + m1
}

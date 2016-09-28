package money

func (m CAD) Abs() CAD {
	if m < 0 {
		return -m
	}

	return m
}

func (m USD) Abs() USD {
	if m < 0 {
		return -m
	}

	return m
}

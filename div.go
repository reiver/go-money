package money

func (m CAD) FloatDiv(val int64) float64 {

	return float64(m) / float64(val) / 100.0
}

func (m USD) FloatDiv(val int64) float64 {

	return float64(m) / float64(val) / 100.0
}

package money


// Deprecated: The functionality that this method provides will eventually come from a currency agnostic method name.
//
// Cents gets the amount of cents for an amount of CAD
func (m CAD) Cents() int64 {
	return int64(m)
}


// Deprecated: The functionality that this method provides will eventually come from a currency agnostic method name.
//
// Cents gets the amount of cents for an amount of USD
func (m USD) Cents() int64 {
	return int64(m)
}


// Deprecated: The functionality that this method provides will eventually come from a currency agnostic method name.
//
// Dollars gets the amount of dollars for an amount of CAD rounding to the nearest dollar
func (m CAD) Dollars() int64 {
	return round(float64(m) / float64(CAD_DOLLAR))
}


// Deprecated: The functionality that this method provides will eventually come from a currency agnostic method name.
//
// Dollars gets the amount of dollars for an amount of USD rounding to the nearest dollar
func (m USD) Dollars() int64 {
	return round(float64(m) / float64(USD_DOLLAR))
}


func round(val float64) int64 {
	if val < 0 {
		return int64(val - 0.5)
	}
	return int64(val + 0.5)
}

package money

// CADEightBits returns eight-bits in Canadian Dollars.
//
// Eight-bits means 100¢ (i.e., $1.00).
//
// A bit means ⅛ of a dollar.
// And thus eight-bits is 1 dollar (i.e., 100¢ or $1.00).
func CADEightBits() CAD {
	return CADCents(100)
}

// USDEightBits returns eight-bits in U.S. Dollars.
//
// Eight-bits means 100¢ (i.e., $1.00).
//
// A bit means ⅛ of a dollar.
// And thus eight-bits is 1 dollar (i.e., 100¢ or $1.00).
func USDEightBits() USD {
	return USDCents(100)
}

package money

// CADSixBits returns six-bits in Canadian Dollars.
//
// Six-bits means 75¢ (i.e., $0.75).
//
// A bit means ⅛ of a dollar.
// And thus six-bits is ¾ of a dollar (i.e., 75¢ or $0.75).
func CADSixBits() CAD {
	return CADCents(75)
}

// USDSixBits returns six-bits in U.S. Dollars.
//
// Six-bits means 75¢ (i.e., $0.75).
//
// A bit means ⅛ of a dollar.
// And thus six-bits is ¾ of a dollar (i.e., 75¢ or $0.75).
func USDSixBits() USD {
	return USDCents(75)
}

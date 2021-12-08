package money

// USDFourBits returns four-bits in U.S. Dollars.
//
// Four-bits means 50¢ (i.e., $0.50).
//
// A bit means ⅛ of a dollar.
// And thus four-bits is ½ of a dollar (i.e., 50¢ or $0.50).
func CADFourBits() CAD {
	return CADCents(50)
}

// CADFourBits returns four-bits in U.S. Dollars.
//
// Four-bits means 50¢ (i.e., $0.50).
//
// A bit means ⅛ of a dollar.
// And thus four-bits is ½ of a dollar (i.e., 50¢ or $0.50).
func USDFourBits() USD {
	return USDCents(50)
}

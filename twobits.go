package money

// CADTwoBits returns two-bits in Canadian Dollars.
//
// Two-bits means 25¢ (i.e., $0.25).
//
// A bit means ⅛ of a dollar.
// And thus two-bits is ¼ of a dollar (i.e., 25¢ or $0.25).
func CADTwoBits() CAD {
	return CADCents(25)
}

// CADTwoBits returns two-bits in U.S. Dollars.
//
// Two-bits means 25¢ (i.e., $0.25).
//
// A bit means ⅛ of a dollar.
// And thus two-bits is ¼ of a dollar (i.e., 25¢ or $0.25).
func USDTwoBits() USD {
	return USDCents(25)
}

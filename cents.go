package money

func CADCents(value int64) CAD {
	return CAD{
		loaded:true,
		value:value,
	}
}

func USDCents(value int64) USD {
	return USD{
		loaded:true,
		value:value,
	}
}

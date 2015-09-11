package money


// Comment money amounts in different currencies.
const (
	CAD_CENT      CAD = 1
	CAD_NICKEL    CAD = 5
	CAD_DIME      CAD = 10
	CAD_QUARTER   CAD = 25
	CAD_TWO_BITS  CAD = 25
	CAD_FOUR_BITS CAD = 50
	CAD_DOLLAR    CAD = 100

	USD_CENT      USD = 1
	USD_NICKEL    USD = 5
	USD_DIME      USD = 10
	USD_QUARTER   USD = 25
	USD_TWO_BITS  USD = 25
	USD_FOUR_BITS USD = 50
	USD_DOLLAR    USD = 100
)


// A CAD represents a Canadian Dollar amount.
type CAD int64


// A USD represents a U.S. Dollar amount.
type USD int64


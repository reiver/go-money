package money

// CAD represents a Canadian Dollar amount.
//
// Note that CAD is an option-type, and thus its uninitialized form is nothing (and not zero).
type CAD struct {
	value int64
	loaded bool
}

// USD represents a U.S. Dollar amount.
//
// Note that USD is an option-type, and thus its uninitialized form is nothing (and not zero).
type USD struct {
	value int64
	loaded bool
}


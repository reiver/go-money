package money


import (
	"errors"
)


var (
	errCannotScanNonString = errors.New("Cannot scan non-string.")
)


// Scan is to make it so *money.CAD fits the sql.Scanner interface.
func (m *CAD) Scan(src interface{}) error {
	s, ok := src.(string)
	if !ok {
		return errCannotScanNonString
	}

	mm, err := ParseCAD(s)
	if nil != err {
		return err
	}
	*m = mm

	return nil
}


// Scan is to make it so *money.USD fits the sql.Scanner interface.
func (m *USD) Scan(src interface{}) error {
	s, ok := src.(string)
	if !ok {
		return errCannotScanNonString
	}

	mm, err := ParseUSD(s)
	if nil != err {
		return err
	}
	*m = mm

	return nil
}

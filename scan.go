package money


import (
	"fmt"
)


// Scan is to make it so *money.CAD fits the sql.Scanner interface.
func (m *CAD) Scan(src interface{}) error {
	if bs, ok := src.([]byte); ok {
		src = string(bs)
	}

	s, ok := src.(string)
	if !ok {
		return fmt.Errorf("Cannot scan something that is not a string or []byte: %T", src)
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
	if bs, ok := src.([]byte); ok {
		src = string(bs)
	}

	s, ok := src.(string)
	if !ok {
		return fmt.Errorf("Cannot scan something that is not a string or []byte: %T", src)
	}

	mm, err := ParseUSD(s)
	if nil != err {
		return err
	}
	*m = mm

	return nil
}

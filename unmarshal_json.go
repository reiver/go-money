package money


import (
	"encoding/json"
)


// UnmarshalJSON is to make it so money.CAD fits the json.Marshaler interface.
func (m *CAD) UnmarshalJSON(src []byte) error {

	var s string
	if err := json.Unmarshal(src, &s); nil != err {
		return err
	}

	if err := m.Scan(s); nil != err {
		return err
	}

	return nil
}


// UnmarshalJSON is to make it so money.USD fits the json.Marshaler interface.
func (m *USD) UnmarshalJSON(src []byte) error {

	var s string
	if err := json.Unmarshal(src, &s); nil != err {
		return err
	}

	if err := m.Scan(s); nil != err {
		return err
	}

	return nil
}

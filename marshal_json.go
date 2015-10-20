package money


import (
	"encoding/json"
)


// MarshalJSON is to make it so money.CAD fits the json.Marshaler interface.
func (m CAD) MarshalJSON() ([]byte, error) {
	s := m.String()
	return json.Marshal(s)
}


// MarshalJSON is to make it so money.USD fits the json.Marshaler interface.
func (m USD) MarshalJSON() ([]byte, error) {
	s := m.String()
	return json.Marshal(s)
}

package money

import (
	"encoding/json"
	"strings"
)

// UnmarshalJSON is to make it so money.CAD fits the json.Marshaler interface.
func (m *CAD) UnmarshalJSON(src []byte) error {

	var s string
	if err := json.Unmarshal(src, &s); nil != err {
		switch err.(type) {
		case *json.UnmarshalTypeError:
			var n json.Number
			if err := json.Unmarshal(src, &n); nil != err {
				return err
			}

			s = trimDecimalPoints(n.String())

		default:
			return err
		}
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
		switch err.(type) {
		case *json.UnmarshalTypeError:
			var n json.Number
			if err := json.Unmarshal(src, &n); nil != err {
				return err
			}
			s = trimDecimalPoints(n.String())

		default:
			return err
		}
	}

	if err := m.Scan(s); nil != err {
		return err
	}

	return nil
}

func trimDecimalPoints(s string) string {

	if strings.Contains(s, ".") == false {
		return s
	}

	index := strings.Index(s, ".00")

	if index == -1 {
		return s
	}

	return s[0 : index+3]

}

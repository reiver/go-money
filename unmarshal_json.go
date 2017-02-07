package money

import (
	"encoding/json"
	"github.com/travis/errors"
	"strings"
)

var ErrTrimmingNonZeroDecimalPoint = errors.New("trimming non zero decimal point")

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

			s, err = trimDecimalPoints(n.String())
			if nil != err {
				return err
			}

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
			s, err = trimDecimalPoints(n.String())
			if nil != err {
				return err
			}

		default:
			return err
		}
	}

	if err := m.Scan(s); nil != err {
		return err
	}

	return nil
}

func trimDecimalPoints(s string) (string, error) {

	if strings.Contains(s, ".") == false {
		return s, nil
	}

	index := strings.Index(s, ".")

	if len(s) < index+3 {
		return s, nil
	}

	//check that all of the decimal points after 2 places are  zeros
	for i := index + 3; i < len(s); i++ {
		if s[i] != '0' {
			return "", ErrTrimmingNonZeroDecimalPoint
		}

	}
	return s[0 : index+3], nil

}

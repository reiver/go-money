package money

import (
	"testing"
)

func testsForUnmashalJson() []struct {
	Json     string
	Expected string
	Err      error
} {

	tests := []struct {
		Json     string
		Expected string
		Err      error
	}{
		{
			Json:     `"$0"`,
			Expected: "$0.00",
		},
		{
			Json:     `"$0.00"`,
			Expected: "$0.00",
		},

		{
			Json:     `"$1"`,
			Expected: "$1.00",
		},
		{
			Json:     `"$1.00"`,
			Expected: "$1.00",
		},

		{
			Json:     `"$123.45"`,
			Expected: "$123.45",
		},

		{
			Json:     `"$1234.45"`,
			Expected: "$1234.45",
		},
		{
			Json:     `"$1,234.45"`,
			Expected: "$1234.45",
		},
		{
			Json:     `"$1.234,45"`,
			Expected: "$1234.45",
		},

		{
			Json:     `"$1234567"`,
			Expected: "$1234567.00",
		},
		{
			Json:     `"$1234567.00"`,
			Expected: "$1234567.00",
		},
		{
			Json:     `"$1234567,00"`,
			Expected: "$1234567.00",
		},
		{
			Json:     `"$1,234,567"`,
			Expected: "$1234567.00",
		},
		{
			Json:     `"$1,234,567.00"`,
			Expected: "$1234567.00",
		},
		{
			Json:     `"$1.234.567"`,
			Expected: "$1234567.00",
		},
		{
			Json:     `"$1.234.567,00"`,
			Expected: "$1234567.00",
		},

		{
			Json:     `"-$1234567"`,
			Expected: "-$1234567.00",
		},
		{
			Json:     `"-$1234567.00"`,
			Expected: "-$1234567.00",
		},
		{
			Json:     `"-$1234567,00"`,
			Expected: "-$1234567.00",
		},
		{
			Json:     `"-$1,234,567"`,
			Expected: "-$1234567.00",
		},
		{
			Json:     `"-$1,234,567.00"`,
			Expected: "-$1234567.00",
		},
		{
			Json:     `"-$1.234.567"`,
			Expected: "-$1234567.00",
		},
		{
			Json:     `"-$1.234.567,00"`,
			Expected: "-$1234567.00",
		},

		{
			Json:     `"$1234567.89"`,
			Expected: "$1234567.89",
		},
		{
			Json:     `"$1234567,89"`,
			Expected: "$1234567.89",
		},
		{
			Json:     `"$1,234,567.89"`,
			Expected: "$1234567.89",
		},
		{
			Json:     `"$1.234.567,89"`,
			Expected: "$1234567.89",
		},
		{
			Json:     `0`,
			Expected: "$0.00",
		},
		{
			Json:     `0.00`,
			Expected: "$0.00",
		},

		{
			Json:     `1`,
			Expected: "$1.00",
		},
		{
			Json:     `1.00`,
			Expected: "$1.00",
		},

		{
			Json:     `123.45`,
			Expected: "$123.45",
		},

		{
			Json:     `1234.45`,
			Expected: "$1234.45",
		},
		{
			Json:     `1234.4`,
			Expected: "$1234.40",
		},
		{
			Json:     `5.5500`,
			Expected: "$5.55",
		},
		{
			Json:     `5.5503`,
			Expected: "$1234.40",
			Err:      ErrTrimmingNonZeroDecimalPoint,
		},
		{
			Json:     `5.5550`,
			Expected: "$1234.40",
			Err:      ErrTrimmingNonZeroDecimalPoint,
		},
	}

	return tests
}

func TestUnmarshalJsonCAD(t *testing.T) {

	tests := testsForUnmashalJson()

	for testNumber, test := range tests {

		var m CAD

		if err := (&m).UnmarshalJSON([]byte(test.Json)); nil != err {

			if err != test.Err {
				t.Errorf("For test #%d, did not expect an error after passing JSON data %q, but got one: %v", testNumber, test.Json, err.Error())
			}

			continue
		}

		if expected, actual := test.Expected, m.String(); expected != actual {
			t.Errorf("For test #%d, expected unmarshaled JSON (i.e., money) to be %q, but actually was %q; for original JSON: %q.", testNumber, expected, actual, test.Json)
			continue
		}
	}

}

func TestUnmarshalJsonUSD(t *testing.T) {

	tests := testsForUnmashalJson()

	for testNumber, test := range tests {

		var m USD

		if err := (&m).UnmarshalJSON([]byte(test.Json)); nil != err {
			if err != test.Err {
				t.Errorf("For test #%d, did not expect an error after passing JSON data %q, but got one: %v", testNumber, test.Json, err.Error())
			}
			continue
		}

		if expected, actual := test.Expected, m.String(); expected != actual {
			t.Errorf("For test #%d, expected unmarshaled JSON (i.e., money) to be %q, but actually was %q; for original JSON: %q.", testNumber, expected, actual, test.Json)
			continue
		}
	}
}

package money


import (
	"testing"
)


func testsForMashalJson() []struct{String string; Expected string} {

	tests := []struct{
		String   string
		Expected string
	}{
		{
			String:    "$0",
			Expected: `"$0.00"`,

		},
		{
			String:    "$0.00",
			Expected: `"$0.00"`,

		},

		{
			String:    "$1",
			Expected: `"$1.00"`,

		},
		{
			String:    "$1.00",
			Expected: `"$1.00"`,

		},

		{
			String:    "$123.45",
			Expected: `"$123.45"`,
		},

		{
			String:    "$1234.45",
			Expected: `"$1234.45"`,
		},
		{
			String:    "$1,234.45",
			Expected: `"$1234.45"`,
		},



		{
			String:    "$1234567",
			Expected: `"$1234567.00"`,
		},
		{
			String:    "$1234567.00",
			Expected: `"$1234567.00"`,
		},
		{
			String:    "$1234567,00",
			Expected: `"$1234567.00"`,
		},
		{
			String:    "$1,234,567",
			Expected: `"$1234567.00"`,
		},
		{
			String:    "$1,234,567.00",
			Expected: `"$1234567.00"`,
		},
		{
			String:    "$1.234.567",
			Expected: `"$1234567.00"`,
		},
		{
			String:    "$1.234.567,00",
			Expected: `"$1234567.00"`,
		},



		{
			String:    "-$1234567",
			Expected: `"-$1234567.00"`,
		},
		{
			String:    "-$1234567.00",
			Expected: `"-$1234567.00"`,
		},
		{
			String:    "-$1234567,00",
			Expected: `"-$1234567.00"`,
		},
		{
			String:    "-$1,234,567",
			Expected: `"-$1234567.00"`,
		},
		{
			String:    "-$1,234,567.00",
			Expected: `"-$1234567.00"`,
		},
		{
			String:    "-$1.234.567",
			Expected: `"-$1234567.00"`,
		},
		{
			String:    "-$1.234.567,00",
			Expected: `"-$1234567.00"`,
		},


		{
			String:    "$1234567.89",
			Expected: `"$1234567.89"`,
		},
		{
			String:    "$1234567,89",
			Expected: `"$1234567.89"`,
		},
		{
			String:    "$1,234,567.89",
			Expected: `"$1234567.89"`,
		},
		{
			String:    "$1.234.567,89",
			Expected: `"$1234567.89"`,
		},
	}


	return tests
}


func TestMarshalJsonCAD(t *testing.T) {

	tests := testsForMashalJson()

	for testNumber, test := range tests {

		m, err := ParseCAD(test.String)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error after pasing %q, but got one: %v", testNumber, test.String, err)
			continue
		}

		bytes, err := m.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error after marshaling %q, but got one: %v", testNumber, test.String, err)
			continue
		}

		if expected, actual := test.Expected, string(bytes); expected != actual {
			t.Errorf("For test #%d, expected marshaled JSON to be %q, but actually was %q; for original: %q.", testNumber, expected, actual, test.String)
			continue
		}
	}
}


func TestMarshalJsonUSD(t *testing.T) {

	tests := testsForMashalJson()

	for testNumber, test := range tests {

		m, err := ParseUSD(test.String)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error after pasing %q, but got one: %v", testNumber, test.String, err)
			continue
		}

		bytes, err := m.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error after marshaling %q, but got one: %v", testNumber, test.String, err)
			continue
		}

		if expected, actual := test.Expected, string(bytes); expected != actual {
			t.Errorf("For test #%d, expected marshaled JSON to be %q, but actually was %q; for original: %q.", testNumber, expected, actual, test.String)
			continue
		}
	}
}

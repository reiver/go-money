package money


import (
	"testing"
)


func TestParseCAD(t *testing.T) {

	tests := []struct{
		String   string
		Expected CAD
	}{
		{
			String:   "$0",
			Expected: 0 * CAD_DOLLAR,
		},
		{
			String:   "$1",
			Expected: 1 * CAD_DOLLAR,
		},
		{
			String:   "$2",
			Expected: 2 * CAD_DOLLAR,
		},
		{
			String:   "$3",
			Expected: 3 * CAD_DOLLAR,
		},
		{
			String:   "$4",
			Expected: 4 * CAD_DOLLAR,
		},
		{
			String:   "$5",
			Expected: 5 * CAD_DOLLAR,
		},
		{
			String:   "$10",
			Expected: 10 * CAD_DOLLAR,
		},
		{
			String:   "$12345",
			Expected: 12345 * CAD_DOLLAR,
		},
		{
			String:   "$12.34",
			Expected: (12 * CAD_DOLLAR) + (34 * CAD_CENT),
		},
		{
			String:   "$1.01",
			Expected: (1 * CAD_DOLLAR) + (1 * CAD_CENT),
		},
		{
			String:   "$0.00",
			Expected: 0 * CAD_CENT,
		},
		{
			String:   "$0.01",
			Expected: 1 * CAD_CENT,
		},
		{
			String:   "$0.02",
			Expected: 2 * CAD_CENT,
		},
		{
			String:   "$0.03",
			Expected: 3 * CAD_CENT,
		},
		{
			String:   "$0.04",
			Expected: 4 * CAD_CENT,
		},
		{
			String:   "$0.05",
			Expected: 5 * CAD_CENT,
		},
		{
			String:   "$0.10",
			Expected: 10 * CAD_CENT,
		},
		{
			String:   "$123.45",
			Expected: (123 * CAD_DOLLAR) + (45 * CAD_CENT),
		},
		{
			String:   "$1.00",
			Expected: 1 * CAD_DOLLAR,
		},
		{
			String:   "$2.00",
			Expected: 2 * CAD_DOLLAR,
		},
		{
			String:   "$3.00",
			Expected: 3 * CAD_DOLLAR,
		},
		{
			String:   "$4.00",
			Expected: 4 * CAD_DOLLAR,
		},
		{
			String:   "$5.00",
			Expected: 5 * CAD_DOLLAR,
		},
		{
			String:   "$10.00",
			Expected: 10 * CAD_DOLLAR,
		},
	}


	for testNumber, test := range tests {

		actual, err := ParseCAD(test.String)
		if nil != err {
			t.Errorf("For test #%d, received unexpected error: %v", testNumber, err)
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, parsed %q, expected %q but actually got %q.", testNumber, test.String, expected, actual)
		}
	}
}


func TestParseUSD(t *testing.T) {

	tests := []struct{
		String   string
		Expected USD
	}{
		{
			String:   "$0",
			Expected: 0 * USD_DOLLAR,
		},
		{
			String:   "$1",
			Expected: 1 * USD_DOLLAR,
		},
		{
			String:   "$2",
			Expected: 2 * USD_DOLLAR,
		},
		{
			String:   "$3",
			Expected: 3 * USD_DOLLAR,
		},
		{
			String:   "$4",
			Expected: 4 * USD_DOLLAR,
		},
		{
			String:   "$5",
			Expected: 5 * USD_DOLLAR,
		},
		{
			String:   "$10",
			Expected: 10 * USD_DOLLAR,
		},
		{
			String:   "$12345",
			Expected: 12345 * USD_DOLLAR,
		},
		{
			String:   "$12.34",
			Expected: (12 * USD_DOLLAR) + (34 * USD_CENT),
		},
		{
			String:   "$1.01",
			Expected: (1 * USD_DOLLAR) + (1 * USD_CENT),
		},
		{
			String:   "$0.00",
			Expected: 0 * USD_CENT,
		},
		{
			String:   "$0.01",
			Expected: 1 * USD_CENT,
		},
		{
			String:   "$0.02",
			Expected: 2 * USD_CENT,
		},
		{
			String:   "$0.03",
			Expected: 3 * USD_CENT,
		},
		{
			String:   "$0.04",
			Expected: 4 * USD_CENT,
		},
		{
			String:   "$0.05",
			Expected: 5 * USD_CENT,
		},
		{
			String:   "$0.10",
			Expected: 10 * USD_CENT,
		},
		{
			String:   "$123.45",
			Expected: (123 * USD_DOLLAR) + (45 * USD_CENT),
		},
		{
			String:   "$1.00",
			Expected: 1 * USD_DOLLAR,
		},
		{
			String:   "$2.00",
			Expected: 2 * USD_DOLLAR,
		},
		{
			String:   "$3.00",
			Expected: 3 * USD_DOLLAR,
		},
		{
			String:   "$4.00",
			Expected: 4 * USD_DOLLAR,
		},
		{
			String:   "$5.00",
			Expected: 5 * USD_DOLLAR,
		},
		{
			String:   "$10.00",
			Expected: 10 * USD_DOLLAR,
		},
	}


	for testNumber, test := range tests {

		actual, err := ParseUSD(test.String)
		if nil != err {
			t.Errorf("For test #%d, received unexpected error: %v", testNumber, err)
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, parsed %q, expected %q but actually got %q.", testNumber, test.String, expected, actual)
		}
	}
}

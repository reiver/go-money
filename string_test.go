package money


import (
	"testing"
)


func TestStringCAD(t *testing.T) {

	tests := []struct{
		Money    CAD
		Expected string
	}{
		{
			Money:    0 * CAD_CENT,
			Expected: "$0.00",
		},
		{
			Money:    0 * CAD_NICKEL,
			Expected: "$0.00",
		},
		{
			Money:    0 * CAD_DIME,
			Expected: "$0.00",
		},
		{
			Money:   0 * CAD_QUARTER,
			Expected: "$0.00",
		},
		{
			Money:    0 * CAD_TWO_BITS,
			Expected: "$0.00",
		},
		{
			Money:    0 * CAD_FOUR_BITS,
			Expected: "$0.00",
		},
		{
			Money:    0 * CAD_DOLLAR,
			Expected: "$0.00",
		},


		{
			Money:    CAD_CENT,
			Expected: "$0.01",
		},
		{
			Money:    CAD_NICKEL,
			Expected: "$0.05",
		},
		{
			Money:    CAD_DIME,
			Expected: "$0.10",
		},
		{
			Money:    CAD_QUARTER,
			Expected: "$0.25",
		},
		{
			Money:    CAD_TWO_BITS,
			Expected: "$0.25",
		},
		{
			Money:    CAD_FOUR_BITS,
			Expected: "$0.50",
		},
		{
			Money:    CAD_DOLLAR,
			Expected: "$1.00",
		},


		{
			Money:    1 * CAD_CENT,
			Expected: "$0.01",
		},
		{
			Money:    2 * CAD_CENT,
			Expected: "$0.02",
		},
		{
			Money:    3 * CAD_CENT,
			Expected: "$0.03",
		},
		{
			Money:    4 * CAD_CENT,
			Expected: "$0.04",
		},
		{
			Money:    5 * CAD_CENT,
			Expected: "$0.05",
		},
		{
			Money:    10 * CAD_CENT,
			Expected: "$0.10",
		},
		{
			Money:    12345 * CAD_CENT,
			Expected: "$123.45",
		},


		{
			Money:    -1 * CAD_CENT,
			Expected: "-$0.01",
		},
		{
			Money:    -2 * CAD_CENT,
			Expected: "-$0.02",
		},
		{
			Money:    -3 * CAD_CENT,
			Expected: "-$0.03",
		},
		{
			Money:    -4 * CAD_CENT,
			Expected: "-$0.04",
		},
		{
			Money:    -5 * CAD_CENT,
			Expected: "-$0.05",
		},
		{
			Money:    -10 * CAD_CENT,
			Expected: "-$0.10",
		},
		{
			Money:    -12345 * CAD_CENT,
			Expected: "-$123.45",
		},


		{
			Money:    1 * CAD_DOLLAR,
			Expected: "$1.00",
		},
		{
			Money:    2 * CAD_DOLLAR,
			Expected: "$2.00",
		},
		{
			Money:    3 * CAD_DOLLAR,
			Expected: "$3.00",
		},
		{
			Money:    4 * CAD_DOLLAR,
			Expected: "$4.00",
		},
		{
			Money:    5 * CAD_DOLLAR,
			Expected: "$5.00",
		},
		{
			Money:    10 * CAD_DOLLAR,
			Expected: "$10.00",
		},
		{
			Money:    12345 * CAD_DOLLAR,
			Expected: "$12345.00",
		},


		{
			Money:    -1 * CAD_DOLLAR,
			Expected: "-$1.00",
		},
		{
			Money:    -2 * CAD_DOLLAR,
			Expected: "-$2.00",
		},
		{
			Money:    -3 * CAD_DOLLAR,
			Expected: "-$3.00",
		},
		{
			Money:    -4 * CAD_DOLLAR,
			Expected: "-$4.00",
		},
		{
			Money:    -5 * CAD_DOLLAR,
			Expected: "-$5.00",
		},
		{
			Money:    -10 * CAD_DOLLAR,
			Expected: "-$10.00",
		},
		{
			Money:    -12345 * CAD_DOLLAR,
			Expected: "-$12345.00",
		},
	}


	for testNumber, test := range tests {

		actual := test.Money.String()

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, Expected %q but actually got %q.", testNumber, expected, actual)
		}
	}
}


func TestStringUSD(t *testing.T) {

	tests := []struct{
		Money    USD
		Expected string
	}{
		{
			Money:    0 * USD_CENT,
			Expected: "$0.00",
		},
		{
			Money:    0 * USD_NICKEL,
			Expected: "$0.00",
		},
		{
			Money:    0 * USD_DIME,
			Expected: "$0.00",
		},
		{
			Money:    0 * USD_QUARTER,
			Expected: "$0.00",
		},
		{
			Money:    0 * USD_TWO_BITS,
			Expected: "$0.00",
		},
		{
			Money:    0 * USD_FOUR_BITS,
			Expected: "$0.00",
		},
		{
			Money:    0 * USD_DOLLAR,
			Expected: "$0.00",
		},


		{
			Money:    USD_CENT,
			Expected: "$0.01",
		},
		{
			Money:    USD_NICKEL,
			Expected: "$0.05",
		},
		{
			Money:    USD_DIME,
			Expected: "$0.10",
		},
		{
			Money:    USD_QUARTER,
			Expected: "$0.25",
		},
		{
			Money:    USD_TWO_BITS,
			Expected: "$0.25",
		},
		{
			Money:    USD_FOUR_BITS,
			Expected: "$0.50",
		},
		{
			Money:    USD_DOLLAR,
			Expected: "$1.00",
		},


		{
			Money:    1 * USD_CENT,
			Expected: "$0.01",
		},
		{
			Money:    2 * USD_CENT,
			Expected: "$0.02",
		},
		{
			Money:    3 * USD_CENT,
			Expected: "$0.03",
		},
		{
			Money:    4 * USD_CENT,
			Expected: "$0.04",
		},
		{
			Money:    5 * USD_CENT,
			Expected: "$0.05",
		},
		{
			Money:    10 * USD_CENT,
			Expected: "$0.10",
		},
		{
			Money:    12345 * USD_CENT,
			Expected: "$123.45",
		},


		{
			Money:    -1 * USD_CENT,
			Expected: "-$0.01",
		},
		{
			Money:    -2 * USD_CENT,
			Expected: "-$0.02",
		},
		{
			Money:    -3 * USD_CENT,
			Expected: "-$0.03",
		},
		{
			Money:    -4 * USD_CENT,
			Expected: "-$0.04",
		},
		{
			Money:    -5 * USD_CENT,
			Expected: "-$0.05",
		},
		{
			Money:    -10 * USD_CENT,
			Expected: "-$0.10",
		},
		{
			Money:    -12345 * USD_CENT,
			Expected: "-$123.45",
		},


		{
			Money:    1 * USD_DOLLAR,
			Expected: "$1.00",
		},
		{
			Money:    2 * USD_DOLLAR,
			Expected: "$2.00",
		},
		{
			Money:    3 * USD_DOLLAR,
			Expected: "$3.00",
		},
		{
			Money:    4 * USD_DOLLAR,
			Expected: "$4.00",
		},
		{
			Money:    5 * USD_DOLLAR,
			Expected: "$5.00",
		},
		{
			Money:    10 * USD_DOLLAR,
			Expected: "$10.00",
		},
		{
			Money:    12345 * USD_DOLLAR,
			Expected: "$12345.00",
		},


		{
			Money:    -1 * USD_DOLLAR,
			Expected: "-$1.00",
		},
		{
			Money:    -2 * USD_DOLLAR,
			Expected: "-$2.00",
		},
		{
			Money:    -3 * USD_DOLLAR,
			Expected: "-$3.00",
		},
		{
			Money:    -4 * USD_DOLLAR,
			Expected: "-$4.00",
		},
		{
			Money:    -5 * USD_DOLLAR,
			Expected: "-$5.00",
		},
		{
			Money:    -10 * USD_DOLLAR,
			Expected: "-$10.00",
		},
		{
			Money:    -12345 * USD_DOLLAR,
			Expected: "-$12345.00",
		},
	}


	for testNumber, test := range tests {

		actual := test.Money.String()

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, Expected %q but actually got %q.", testNumber, expected, actual)
		}
	}
}

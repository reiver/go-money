package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestParseCAD(t *testing.T) {

	tests := []struct{
		String   string
		Expected money.CAD
	}{
		{
			String:   "$0",
			Expected: money.CADDollars(0),
		},
		{
			String:   "0",
			Expected: money.CADDollars(0),
		},
		{
			String:   "-$0",
			Expected: money.CADDollars(0),
		},
		{
			String:   "-0",
			Expected: money.CADDollars(0),
		},



		{
			String:   "$1",
			Expected: money.CADDollars(1),
		},
		{
			String:   "1",
			Expected: money.CADDollars(1),
		},
		{
			String:   "-$1",
			Expected: money.CADDollars(-1),
		},
		{
			String:   "-1",
			Expected: money.CADDollars(-1),
		},



		{
			String:   "$2",
			Expected: money.CADDollars(2),
		},
		{
			String:   "2",
			Expected: money.CADDollars(2),
		},
		{
			String:   "-$2",
			Expected: money.CADDollars(-2),
		},
		{
			String:   "-2",
			Expected: money.CADDollars(-2),
		},



		{
			String:   "$3",
			Expected: money.CADDollars(3),
		},
		{
			String:   "3",
			Expected: money.CADDollars(3),
		},
		{
			String:   "-$3",
			Expected: money.CADDollars(-3),
		},
		{
			String:   "-3",
			Expected: money.CADDollars(-3),
		},



		{
			String:   "$4",
			Expected: money.CADDollars(4),
		},
		{
			String:   "4",
			Expected: money.CADDollars(4),
		},
		{
			String:   "-$4",
			Expected: money.CADDollars(-4),
		},
		{
			String:   "-4",
			Expected: money.CADDollars(-4),
		},



		{
			String:   "$5",
			Expected: money.CADDollars(5),
		},
		{
			String:   "5",
			Expected: money.CADDollars(5),
		},
		{
			String:   "-$5",
			Expected: money.CADDollars(-5),
		},
		{
			String:   "-5",
			Expected: money.CADDollars(-5),
		},



		{
			String:   "$10",
			Expected: money.CADDollars(10),
		},
		{
			String:   "10",
			Expected: money.CADDollars(10),
		},
		{
			String:   "-$10",
			Expected: money.CADDollars(-10),
		},
		{
			String:   "-10",
			Expected: money.CADDollars(-10),
		},



		{
			String:   "$12345",
			Expected: money.CADDollars(12345),
		},
		{
			String:   "12345",
			Expected: money.CADDollars(12345),
		},
		{
			String:   "-$12345",
			Expected: money.CADDollars(-12345),
		},
		{
			String:   "-12345",
			Expected: money.CADDollars(-12345),
		},



		{
			String:   "$12.34",
			Expected: money.CADCanonicalForm(12,34),
		},
		{
			String:   "12.34",
			Expected: money.CADCanonicalForm(12,34),
		},
		{
			String:   "-$12.34",
			Expected: money.CADCanonicalForm(-12,-34),
		},
		{
			String:   "-12.34",
			Expected: money.CADCanonicalForm(-12,-34),
		},


		{
			String:   "$1.01",
			Expected: money.CADCanonicalForm(1,1),
		},
		{
			String:   "1.01",
			Expected: money.CADCanonicalForm(1,1),
		},
		{
			String:   "-$1.01",
			Expected: money.CADCanonicalForm(-1,-1),
		},
		{
			String:   "-1.01",
			Expected: money.CADCanonicalForm(-1,-1),
		},



		{
			String:   "$0.00",
			Expected: money.CADCents(0),
		},
		{
			String:   "0.00",
			Expected: money.CADCents(0),
		},
		{
			String:   "-$0.00",
			Expected: money.CADCents(0),
		},
		{
			String:   "-0.00",
			Expected: money.CADCents(0),
		},



		{
			String:   "$0.01",
			Expected: money.CADCents(1),
		},
		{
			String:   "0.01",
			Expected: money.CADCents(1),
		},
		{
			String:   "-$0.01",
			Expected: money.CADCents(-1),
		},
		{
			String:   "-0.01",
			Expected: money.CADCents(-1),
		},



		{
			String:   "$0.02",
			Expected: money.CADCents(2),
		},
		{
			String:   "0.02",
			Expected: money.CADCents(2),
		},
		{
			String:   "-$0.02",
			Expected: money.CADCents(-2),
		},
		{
			String:   "-0.02",
			Expected: money.CADCents(-2),
		},



		{
			String:   "$0.03",
			Expected: money.CADCents(3),
		},
		{
			String:   "0.03",
			Expected: money.CADCents(3),
		},
		{
			String:   "-$0.03",
			Expected: money.CADCents(-3),
		},
		{
			String:   "-0.03",
			Expected: money.CADCents(-3),
		},



		{
			String:   "$0.04",
			Expected: money.CADCents(4),
		},
		{
			String:   "0.04",
			Expected: money.CADCents(4),
		},
		{
			String:   "-$0.04",
			Expected: money.CADCents(-4),
		},
		{
			String:   "-0.04",
			Expected: money.CADCents(-4),
		},



		{
			String:   "$0.05",
			Expected: money.CADCents(5),
		},
		{
			String:   "0.05",
			Expected: money.CADCents(5),
		},
		{
			String:   "-$0.05",
			Expected: money.CADCents(-5),
		},
		{
			String:   "-0.05",
			Expected: money.CADCents(-5),
		},



		{
			String:   "$0.10",
			Expected: money.CADCents(10),
		},
		{
			String:   "0.10",
			Expected: money.CADCents(10),
		},
		{
			String:   "-$0.10",
			Expected: money.CADCents(-10),
		},
		{
			String:   "-0.10",
			Expected: money.CADCents(-10),
		},



		{
			String:   "$123.45",
			Expected: money.CADCanonicalForm(123,45),
		},
		{
			String:   "123.45",
			Expected: money.CADCanonicalForm(123,45),
		},
		{
			String:   "-$123.45",
			Expected: money.CADCanonicalForm(-123,-45),
		},
		{
			String:   "-123.45",
			Expected: money.CADCanonicalForm(-123,-45),
		},



		{
			String:   "$1.00",
			Expected: money.CADDollars(1),
		},
		{
			String:   "1.00",
			Expected: money.CADDollars(1),
		},
		{
			String:   "-$1.00",
			Expected: money.CADDollars(-1),
		},
		{
			String:   "-1.00",
			Expected: money.CADDollars(-1),
		},



		{
			String:   "$2.00",
			Expected: money.CADDollars(2),
		},
		{
			String:   "2.00",
			Expected: money.CADDollars(2),
		},
		{
			String:   "-$2.00",
			Expected: money.CADDollars(-2),
		},
		{
			String:   "-2.00",
			Expected: money.CADDollars(-2),
		},



		{
			String:   "$3.00",
			Expected: money.CADDollars(3),
		},
		{
			String:   "3.00",
			Expected: money.CADDollars(3),
		},
		{
			String:   "-$3.00",
			Expected: money.CADDollars(-3),
		},
		{
			String:   "-3.00",
			Expected: money.CADDollars(-3),
		},



		{
			String:   "$4.00",
			Expected: money.CADDollars(4),
		},
		{
			String:   "4.00",
			Expected: money.CADDollars(4),
		},
		{
			String:   "-$4.00",
			Expected: money.CADDollars(-4),
		},
		{
			String:   "-4.00",
			Expected: money.CADDollars(-4),
		},



		{
			String:   "$5.00",
			Expected: money.CADDollars(5),
		},
		{
			String:   "5.00",
			Expected: money.CADDollars(5),
		},
		{
			String:   "-$5.00",
			Expected: money.CADDollars(-5),
		},
		{
			String:   "-5.00",
			Expected: money.CADDollars(-5),
		},



		{
			String:   "$10.00",
			Expected: money.CADDollars(10),
		},
		{
			String:   "10.00",
			Expected: money.CADDollars(10),
		},
		{
			String:   "-$10.00",
			Expected: money.CADDollars(-10),
		},
		{
			String:   "-10.00",
			Expected: money.CADDollars(-10),
		},



		{
			String:   "$862.49",
			Expected: money.CADCanonicalForm(862,49),
		},
		{
			String:   "862.49",
			Expected: money.CADCanonicalForm(862,49),
		},
		{
			String:   "-$862.49",
			Expected: money.CADCanonicalForm(-862,-49),
		},
		{
			String:   "-862.49",
			Expected: money.CADCanonicalForm(-862,-49),
		},



		{
			String:   "$0.3",
			Expected: money.CADCents(30),
		},
		{
			String:   "0.3",
			Expected: money.CADCents(30),
		},
		{
			String:   "-$0.3",
			Expected: money.CADCents(-30),
		},
		{
			String:   "-0.3",
			Expected: money.CADCents(-30),
		},



		{
			String:   "$30.0",
			Expected: money.CADDollars(30),
		},
		{
			String:   "-$40.0",
			Expected: money.CADDollars(-40),
		},
		{
			String:   "2342343.4",
			Expected: money.CADCanonicalForm(2342343,40),
		},
		{
			String:   "34,3",
			Expected: money.CADCanonicalForm(34,30),
		},



		{
			String:   ".21",
			Expected: money.CADCents(21),
		},
		{
			String:   "-.21",
			Expected: money.CADCents(-21),
		},
		{
			String:   ".2",
			Expected: money.CADCents(20),
		},
		{
			String:   "-.2",
			Expected: money.CADCents(-20),
		},



		{
			String:   ",21",
			Expected: money.CADCents(21),
		},
		{
			String:   "-,21",
			Expected: money.CADCents(-21),
		},
		{
			String:   ",2",
			Expected: money.CADCents(20),
		},
		{
			String:   "-,2",
			Expected: money.CADCents(-20),
		},
	}

	for testNumber, test := range tests {

		actual, err := money.ParseCAD(test.String)
		if nil != err {
			t.Errorf("For test #%d, received unexpected error: %v", testNumber, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual value, from parsing, is not what was expected.", testNumber)
			t.Logf("STRING: %q", test.String)
			t.Logf("EXPECTED: %q (%#v)", expected, expected)
			t.Logf("ACTUAL:   %q (%#v)", actual, actual)
			continue
		}
	}
}

func TestParseUSD(t *testing.T) {

	tests := []struct{
		String   string
		Expected money.USD
	}{
		{
			String:   "$0",
			Expected: money.USDDollars(0),
		},
		{
			String:   "0",
			Expected: money.USDDollars(0),
		},
		{
			String:   "-$0",
			Expected: money.USDDollars(0),
		},
		{
			String:   "-0",
			Expected: money.USDDollars(0),
		},



		{
			String:   "$1",
			Expected: money.USDDollars(1),
		},
		{
			String:   "1",
			Expected: money.USDDollars(1),
		},
		{
			String:   "-$1",
			Expected: money.USDDollars(-1),
		},
		{
			String:   "-1",
			Expected: money.USDDollars(-1),
		},



		{
			String:   "$2",
			Expected: money.USDDollars(2),
		},
		{
			String:   "2",
			Expected: money.USDDollars(2),
		},
		{
			String:   "-$2",
			Expected: money.USDDollars(-2),
		},
		{
			String:   "-2",
			Expected: money.USDDollars(-2),
		},



		{
			String:   "$3",
			Expected: money.USDDollars(3),
		},
		{
			String:   "3",
			Expected: money.USDDollars(3),
		},
		{
			String:   "-$3",
			Expected: money.USDDollars(-3),
		},
		{
			String:   "-3",
			Expected: money.USDDollars(-3),
		},



		{
			String:   "$4",
			Expected: money.USDDollars(4),
		},
		{
			String:   "4",
			Expected: money.USDDollars(4),
		},
		{
			String:   "-$4",
			Expected: money.USDDollars(-4),
		},
		{
			String:   "-4",
			Expected: money.USDDollars(-4),
		},



		{
			String:   "$5",
			Expected: money.USDDollars(5),
		},
		{
			String:   "5",
			Expected: money.USDDollars(5),
		},
		{
			String:   "-$5",
			Expected: money.USDDollars(-5),
		},
		{
			String:   "-5",
			Expected: money.USDDollars(-5),
		},



		{
			String:   "$10",
			Expected: money.USDDollars(10),
		},
		{
			String:   "10",
			Expected: money.USDDollars(10),
		},
		{
			String:   "-$10",
			Expected: money.USDDollars(-10),
		},
		{
			String:   "-10",
			Expected: money.USDDollars(-10),
		},



		{
			String:   "$12345",
			Expected: money.USDDollars(12345),
		},
		{
			String:   "12345",
			Expected: money.USDDollars(12345),
		},
		{
			String:   "-$12345",
			Expected: money.USDDollars(-12345),
		},
		{
			String:   "-12345",
			Expected: money.USDDollars(-12345),
		},



		{
			String:   "$12.34",
			Expected: money.USDCanonicalForm(12,34),
		},
		{
			String:   "12.34",
			Expected: money.USDCanonicalForm(12,34),
		},
		{
			String:   "-$12.34",
			Expected: money.USDCanonicalForm(-12,-34),
		},
		{
			String:   "-12.34",
			Expected: money.USDCanonicalForm(-12,-34),
		},


		{
			String:   "$1.01",
			Expected: money.USDCanonicalForm(1,1),
		},
		{
			String:   "1.01",
			Expected: money.USDCanonicalForm(1,1),
		},
		{
			String:   "-$1.01",
			Expected: money.USDCanonicalForm(-1,-1),
		},
		{
			String:   "-1.01",
			Expected: money.USDCanonicalForm(-1,-1),
		},



		{
			String:   "$0.00",
			Expected: money.USDCents(0),
		},
		{
			String:   "0.00",
			Expected: money.USDCents(0),
		},
		{
			String:   "-$0.00",
			Expected: money.USDCents(0),
		},
		{
			String:   "-0.00",
			Expected: money.USDCents(0),
		},



		{
			String:   "$0.01",
			Expected: money.USDCents(1),
		},
		{
			String:   "0.01",
			Expected: money.USDCents(1),
		},
		{
			String:   "-$0.01",
			Expected: money.USDCents(-1),
		},
		{
			String:   "-0.01",
			Expected: money.USDCents(-1),
		},



		{
			String:   "$0.02",
			Expected: money.USDCents(2),
		},
		{
			String:   "0.02",
			Expected: money.USDCents(2),
		},
		{
			String:   "-$0.02",
			Expected: money.USDCents(-2),
		},
		{
			String:   "-0.02",
			Expected: money.USDCents(-2),
		},



		{
			String:   "$0.03",
			Expected: money.USDCents(3),
		},
		{
			String:   "0.03",
			Expected: money.USDCents(3),
		},
		{
			String:   "-$0.03",
			Expected: money.USDCents(-3),
		},
		{
			String:   "-0.03",
			Expected: money.USDCents(-3),
		},



		{
			String:   "$0.04",
			Expected: money.USDCents(4),
		},
		{
			String:   "0.04",
			Expected: money.USDCents(4),
		},
		{
			String:   "-$0.04",
			Expected: money.USDCents(-4),
		},
		{
			String:   "-0.04",
			Expected: money.USDCents(-4),
		},



		{
			String:   "$0.05",
			Expected: money.USDCents(5),
		},
		{
			String:   "0.05",
			Expected: money.USDCents(5),
		},
		{
			String:   "-$0.05",
			Expected: money.USDCents(-5),
		},
		{
			String:   "-0.05",
			Expected: money.USDCents(-5),
		},



		{
			String:   "$0.10",
			Expected: money.USDCents(10),
		},
		{
			String:   "0.10",
			Expected: money.USDCents(10),
		},
		{
			String:   "-$0.10",
			Expected: money.USDCents(-10),
		},
		{
			String:   "-0.10",
			Expected: money.USDCents(-10),
		},



		{
			String:   "$123.45",
			Expected: money.USDCanonicalForm(123,45),
		},
		{
			String:   "123.45",
			Expected: money.USDCanonicalForm(123,45),
		},
		{
			String:   "-$123.45",
			Expected: money.USDCanonicalForm(-123,-45),
		},
		{
			String:   "-123.45",
			Expected: money.USDCanonicalForm(-123,-45),
		},



		{
			String:   "$1.00",
			Expected: money.USDDollars(1),
		},
		{
			String:   "1.00",
			Expected: money.USDDollars(1),
		},
		{
			String:   "-$1.00",
			Expected: money.USDDollars(-1),
		},
		{
			String:   "-1.00",
			Expected: money.USDDollars(-1),
		},



		{
			String:   "$2.00",
			Expected: money.USDDollars(2),
		},
		{
			String:   "2.00",
			Expected: money.USDDollars(2),
		},
		{
			String:   "-$2.00",
			Expected: money.USDDollars(-2),
		},
		{
			String:   "-2.00",
			Expected: money.USDDollars(-2),
		},



		{
			String:   "$3.00",
			Expected: money.USDDollars(3),
		},
		{
			String:   "3.00",
			Expected: money.USDDollars(3),
		},
		{
			String:   "-$3.00",
			Expected: money.USDDollars(-3),
		},
		{
			String:   "-3.00",
			Expected: money.USDDollars(-3),
		},



		{
			String:   "$4.00",
			Expected: money.USDDollars(4),
		},
		{
			String:   "4.00",
			Expected: money.USDDollars(4),
		},
		{
			String:   "-$4.00",
			Expected: money.USDDollars(-4),
		},
		{
			String:   "-4.00",
			Expected: money.USDDollars(-4),
		},



		{
			String:   "$5.00",
			Expected: money.USDDollars(5),
		},
		{
			String:   "5.00",
			Expected: money.USDDollars(5),
		},
		{
			String:   "-$5.00",
			Expected: money.USDDollars(-5),
		},
		{
			String:   "-5.00",
			Expected: money.USDDollars(-5),
		},



		{
			String:   "$10.00",
			Expected: money.USDDollars(10),
		},
		{
			String:   "10.00",
			Expected: money.USDDollars(10),
		},
		{
			String:   "-$10.00",
			Expected: money.USDDollars(-10),
		},
		{
			String:   "-10.00",
			Expected: money.USDDollars(-10),
		},



		{
			String:   "$862.49",
			Expected: money.USDCanonicalForm(862,49),
		},
		{
			String:   "862.49",
			Expected: money.USDCanonicalForm(862,49),
		},
		{
			String:   "-$862.49",
			Expected: money.USDCanonicalForm(-862,-49),
		},
		{
			String:   "-862.49",
			Expected: money.USDCanonicalForm(-862,-49),
		},



		{
			String:   "$0.3",
			Expected: money.USDCents(30),
		},
		{
			String:   "0.3",
			Expected: money.USDCents(30),
		},
		{
			String:   "-$0.3",
			Expected: money.USDCents(-30),
		},
		{
			String:   "-0.3",
			Expected: money.USDCents(-30),
		},



		{
			String:   "$30.0",
			Expected: money.USDDollars(30),
		},
		{
			String:   "-$40.0",
			Expected: money.USDDollars(-40),
		},
		{
			String:   "2342343.4",
			Expected: money.USDCanonicalForm(2342343,40),
		},
		{
			String:   "34,3",
			Expected: money.USDCanonicalForm(34,30),
		},



		{
			String:   ".21",
			Expected: money.USDCents(21),
		},
		{
			String:   "-.21",
			Expected: money.USDCents(-21),
		},
		{
			String:   ".2",
			Expected: money.USDCents(20),
		},
		{
			String:   "-.2",
			Expected: money.USDCents(-20),
		},



		{
			String:   ",21",
			Expected: money.USDCents(21),
		},
		{
			String:   "-,21",
			Expected: money.USDCents(-21),
		},
		{
			String:   ",2",
			Expected: money.USDCents(20),
		},
		{
			String:   "-,2",
			Expected: money.USDCents(-20),
		},
	}

	for testNumber, test := range tests {

		actual, err := money.ParseUSD(test.String)
		if nil != err {
			t.Errorf("For test #%d, received unexpected error: %v", testNumber, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual value, from parsing, is not what was expected.", testNumber)
			t.Logf("STRING: %q", test.String)
			t.Logf("EXPECTED: %q (%#v)", expected, expected)
			t.Logf("ACTUAL:   %q (%#v)", actual, actual)
			continue
		}
	}
}

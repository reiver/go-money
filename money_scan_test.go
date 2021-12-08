package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func testsForScan() []struct{String string ; Expected int64} {

	tests := []struct{
		String   string
		Expected int64
	}{
		{
			String:   "$0",
			Expected:   0,
		},
		{
			String:   "$0.00",
			Expected:   0,
		},

		{
			String:   "$0.01",
			Expected:      1,
		},

		{
			String:   "$0.02",
			Expected:      2,
		},

		{
			String:   "$0.03",
			Expected:      3,
		},

		{
			String:   "$0.04",
			Expected:      4,
		},

		{
			String:   "$0.05",
			Expected:      5,
		},

		{
			String:   "$0.06",
			Expected:      6,
		},

		{
			String:   "$0.07",
			Expected:      7,
		},

		{
			String:   "$0.08",
			Expected:      8,
		},

		{
			String:   "$0.09",
			Expected:      9,
		},

		{
			String:   "$0.10",
			Expected:     10,
		},

		{
			String:   "$0.11",
			Expected:     11,
		},

		{
			String:   "$0.12",
			Expected:     12,
		},

		{
			String:   "$0.42",
			Expected:     42,
		},

		{
			String:   "$0.89",
			Expected:     89,
		},

		{
			String:   "$0.99",
			Expected:     99,
		},

		{
			String:   "$1",
			Expected:   100,
		},
		{
			String:   "$1.00",
			Expected:   100,
		},

		{
			String:   "$1.23",
			Expected:   123,
		},

		{
			String:   "$5.43",
			Expected:   543,
		},

		{
			String:   "$9.99",
			Expected:   999,
		},

		{
			String:   "$2",
			Expected:   200,
		},
		{
			String:   "$2.00",
			Expected:   200,
		},

		{
			String:   "$3",
			Expected:   300,
		},
		{
			String:   "$3.00",
			Expected:   300,
		},

		{
			String:   "$4",
			Expected:   400,
		},
		{
			String:   "$4.00",
			Expected:   400,
		},

		{
			String:   "$5",
			Expected:   500,
		},
		{
			String:   "$5.00",
			Expected:   500,
		},

		{
			String:   "$6",
			Expected:   600,
		},
		{
			String:   "$6.00",
			Expected:   600,
		},

		{
			String:   "$7",
			Expected:   700,
		},
		{
			String:   "$7.00",
			Expected:   700,
		},

		{
			String:   "$8",
			Expected:   800,
		},
		{
			String:   "$8.00",
			Expected:   800,
		},

		{
			String:   "$9",
			Expected:   900,
		},
		{
			String:   "$9.00",
			Expected:   900,
		},

		{
			String:   "$10",
			Expected:   1000,
		},
		{
			String:   "$10.00",
			Expected:   1000,
		},

		{
			String:   "$11",
			Expected:   1100,
		},
		{
			String:   "$11.00",
			Expected:   1100,
		},

		{
			String:   "$12",
			Expected:   1200,
		},
		{
			String:   "$12.00",
			Expected:   1200,
		},

		{
			String:   "$42",
			Expected:   4200,
		},
		{
			String:   "$42.00",
			Expected:   4200,
		},

		{
			String:   "$89",
			Expected:   8900,
		},
		{
			String:   "$89.00",
			Expected:   8900,
		},

		{
			String:   "$99",
			Expected:   9900,
		},
		{
			String:   "$99.00",
			Expected:   9900,
		},

		{
			String:   "$100",
			Expected:   10000,
		},
		{
			String:   "$100.00",
			Expected:   10000,
		},

		{
			String:   "$101",
			Expected:   10100,
		},
		{
			String:   "$101.00",
			Expected:   10100,
		},

		{
			String:   "$123",
			Expected:   12300,
		},
		{
			String:   "$123.00",
			Expected:   12300,
		},

		{
			String:   "$123.45",
			Expected:   12345,
		},

		{
			String:   "$543",
			Expected:   54300,
		},
		{
			String:   "$543.00",
			Expected:   54300,
		},

		{
			String:   "$543.21",
			Expected:   54321,
		},

		{
			String:   "$999",
			Expected:   99900,
		},
		{
			String:   "$999.00",
			Expected:   99900,
		},

		{
			String:   "$1000",
			Expected:   100000,
		},
		{
			String:   "$1,000",
			Expected:   100000,
		},
		{
			String:   "$1000.00",
			Expected:   100000,
		},
		{
			String:   "$1,000.00",
			Expected:   100000,
		},

		{
			String:   "$1234",
			Expected:   123400,
		},
		{
			String:   "$1,234",
			Expected:   123400,
		},
		{
			String:   "$1234.00",
			Expected:   123400,
		},
		{
			String:   "$1,234.00",
			Expected:   123400,
		},

		{
			String:   "$1234.56",
			Expected:   123456,
		},
		{
			String:   "$1,234.56",
			Expected:   123456,
		},

		{
			String:   "$9999",
			Expected:   999900,
		},
		{
			String:   "$9,999",
			Expected:   999900,
		},
		{
			String:   "$9999.00",
			Expected:   999900,
		},
		{
			String:   "$9,999.00",
			Expected:   999900,
		},

		{
			String:   "$12345",
			Expected:   1234500,
		},
		{
			String:   "$12,345",
			Expected:   1234500,
		},
		{
			String:   "$12.345",
			Expected:   1234500,
		},
		{
			String:   "$12345.00",
			Expected:   1234500,
		},
		{
			String:   "$12,345.00",
			Expected:   1234500,
		},
		{
			String:   "$12.345,00",
			Expected:   1234500,
		},

		{
			String:   "$12345.67",
			Expected:   1234567,
		},
		{
			String:   "$12,345.67",
			Expected:   1234567,
		},
		{
			String:   "$12345,67",
			Expected:   1234567,
		},
		{
			String:   "$12.345,67",
			Expected:   1234567,
		},

		{
			String:   "$99999",
			Expected:   9999900,
		},
		{
			String:   "$99,999",
			Expected:   9999900,
		},
		{
			String:   "$99.999",
			Expected:   9999900,
		},
		{
			String:   "$99999.00",
			Expected:   9999900,
		},
		{
			String:   "$99,999.00",
			Expected:   9999900,
		},
		{
			String:   "$99.999,00",
			Expected:   9999900,
		},



		{
			String:   "$1234567",
			Expected:   123456700,
		},
		{
			String:   "$1234567.00",
			Expected:   123456700,
		},
		{
			String:   "$1,234,567",
			Expected:   123456700,
		},
		{
			String:   "$1,234,567.00",
			Expected:   123456700,
		},
		{
			String:   "$1.234.567",
			Expected:   123456700,
		},
		{
			String:   "$1.234.567,00",
			Expected:   123456700,
		},



		{
			String:   "$1234567.89",
			Expected:   123456789,
		},
		{
			String:   "$1,234,567.89",
			Expected:   123456789,
		},
		{
			String:   "$1.234.567,89",
			Expected:   123456789,
		},
	}

	return tests
}

func TestScanCAD(t *testing.T) {

	tests := testsForScan()

	for testNumber, test := range tests {

		var expected money.CAD = money.CADCents(test.Expected)

		var actual money.CAD
		(&actual).Scan(test.String)

		if expected != actual {
			t.Errorf("For test #%d, expected %q but actually got %q, for scanned string %q.", testNumber, expected, actual, test.String)
			t.Logf("STRING:   %q", test.String)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
		}


	}
}

func TestScanUSD(t *testing.T) {

	tests := testsForScan()

	for testNumber, test := range tests {

		var expected money.USD = money.USDCents(test.Expected)

		var actual money.USD
		(&actual).Scan(test.String)

		if expected != actual {
			t.Errorf("For test #%d, expected %q but actually got %q, for scanned string %q.", testNumber, expected, actual, test.String)
			t.Logf("STRING:   %q", test.String)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
		}


	}
}

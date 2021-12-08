package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCADCanonicalForm(t *testing.T) {

	tests := []struct{
		Dollars int32
		Cents   int8
		Expected money.CAD
	}{
		{
			Dollars:  0,
			Cents:    0,
			Expected: money.CADCents(0),
		},
		{
			Dollars:  -0,
			Cents:    -0,
			Expected: money.CADCents(0),
		},



		{
			Dollars:  0,
			Cents:    1,
			Expected: money.CADCents(1),
		},
		{
			Dollars:  0,
			Cents:    2,
			Expected: money.CADCents(2),
		},
		{
			Dollars:  0,
			Cents:    3,
			Expected: money.CADCents(3),
		},
		{
			Dollars:  0,
			Cents:    9,
			Expected: money.CADCents(9),
		},
		{
			Dollars:  0,
			Cents:    10,
			Expected: money.CADCents(10),
		},
		{
			Dollars:  0,
			Cents:    15,
			Expected: money.CADCents(15),
		},
		{
			Dollars:  0,
			Cents:    16,
			Expected: money.CADCents(16),
		},
		{
			Dollars:  0,
			Cents:    52,
			Expected: money.CADCents(52),
		},
		{
			Dollars:  0,
			Cents:    99,
			Expected: money.CADCents(99),
		},



		{
			Dollars:  0,
			Cents:    -1,
			Expected: money.CADCents(-1),
		},
		{
			Dollars:  0,
			Cents:    -2,
			Expected: money.CADCents(-2),
		},
		{
			Dollars:  0,
			Cents:    -3,
			Expected: money.CADCents(-3),
		},
		{
			Dollars:  0,
			Cents:    -9,
			Expected: money.CADCents(-9),
		},
		{
			Dollars:  0,
			Cents:    -10,
			Expected: money.CADCents(-10),
		},
		{
			Dollars:  0,
			Cents:    -15,
			Expected: money.CADCents(-15),
		},
		{
			Dollars:  0,
			Cents:    -16,
			Expected: money.CADCents(-16),
		},
		{
			Dollars:  0,
			Cents:    -52,
			Expected: money.CADCents(-52),
		},
		{
			Dollars:  0,
			Cents:    -99,
			Expected: money.CADCents(-99),
		},



		{
			Dollars:  1,
			Cents:    0,
			Expected: money.CADCents(100),
		},
		{
			Dollars:  1,
			Cents:    1,
			Expected: money.CADCents(101),
		},
		{
			Dollars:  1,
			Cents:    2,
			Expected: money.CADCents(102),
		},
		{
			Dollars:  1,
			Cents:    3,
			Expected: money.CADCents(103),
		},
		{
			Dollars:  1,
			Cents:    9,
			Expected: money.CADCents(109),
		},
		{
			Dollars:  1,
			Cents:    10,
			Expected: money.CADCents(110),
		},
		{
			Dollars:  1,
			Cents:    15,
			Expected: money.CADCents(115),
		},
		{
			Dollars:  1,
			Cents:    16,
			Expected: money.CADCents(116),
		},
		{
			Dollars:  1,
			Cents:    52,
			Expected: money.CADCents(152),
		},
		{
			Dollars:  1,
			Cents:    99,
			Expected: money.CADCents(199),
		},



		{
			Dollars:  -1,
			Cents:    -0,
			Expected: money.CADCents(-100),
		},
		{
			Dollars:  -1,
			Cents:    -1,
			Expected: money.CADCents(-101),
		},
		{
			Dollars:  -1,
			Cents:    -2,
			Expected: money.CADCents(-102),
		},
		{
			Dollars:  -1,
			Cents:    -3,
			Expected: money.CADCents(-103),
		},
		{
			Dollars:  -1,
			Cents:    -9,
			Expected: money.CADCents(-109),
		},
		{
			Dollars:  -1,
			Cents:    -10,
			Expected: money.CADCents(-110),
		},
		{
			Dollars:  -1,
			Cents:    -15,
			Expected: money.CADCents(-115),
		},
		{
			Dollars:  -1,
			Cents:    -16,
			Expected: money.CADCents(-116),
		},
		{
			Dollars:  -1,
			Cents:    -52,
			Expected: money.CADCents(-152),
		},
		{
			Dollars:  -1,
			Cents:    -99,
			Expected: money.CADCents(-199),
		},



		{
			Dollars:  12345,
			Cents:    0,
			Expected: money.CADCents(1234500),
		},
		{
			Dollars:  12345,
			Cents:    1,
			Expected: money.CADCents(1234501),
		},
		{
			Dollars:  12345,
			Cents:    2,
			Expected: money.CADCents(1234502),
		},
		{
			Dollars:  12345,
			Cents:    3,
			Expected: money.CADCents(1234503),
		},
		{
			Dollars:  12345,
			Cents:    9,
			Expected: money.CADCents(1234509),
		},
		{
			Dollars:  12345,
			Cents:    10,
			Expected: money.CADCents(1234510),
		},
		{
			Dollars:  12345,
			Cents:    15,
			Expected: money.CADCents(1234515),
		},
		{
			Dollars:  12345,
			Cents:    16,
			Expected: money.CADCents(1234516),
		},
		{
			Dollars:  12345,
			Cents:    52,
			Expected: money.CADCents(1234552),
		},
		{
			Dollars:  12345,
			Cents:    99,
			Expected: money.CADCents(1234599),
		},



		{
			Dollars:  -12345,
			Cents:    -0,
			Expected: money.CADCents(-1234500),
		},
		{
			Dollars:  -12345,
			Cents:    -1,
			Expected: money.CADCents(-1234501),
		},
		{
			Dollars:  -12345,
			Cents:    -2,
			Expected: money.CADCents(-1234502),
		},
		{
			Dollars:  -12345,
			Cents:    -3,
			Expected: money.CADCents(-1234503),
		},
		{
			Dollars:  -12345,
			Cents:    -9,
			Expected: money.CADCents(-1234509),
		},
		{
			Dollars:  -12345,
			Cents:    -10,
			Expected: money.CADCents(-1234510),
		},
		{
			Dollars:  -12345,
			Cents:    -15,
			Expected: money.CADCents(-1234515),
		},
		{
			Dollars:  -12345,
			Cents:    -16,
			Expected: money.CADCents(-1234516),
		},
		{
			Dollars:  -12345,
			Cents:    -52,
			Expected: money.CADCents(-1234552),
		},
		{
			Dollars:  -12345,
			Cents:    -99,
			Expected: money.CADCents(-1234599),
		},
	}

	for testNumber, test := range tests {

		var expected money.CAD = test.Expected

		var actual money.CAD = money.CADCanonicalForm(test.Dollars, test.Cents)

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("DOLLARS:   %d", test.Dollars)
			t.Logf("CENTS:      %d", test.Cents)
			t.Logf("EXPECTED: %s (%#v)", expected, expected)
			t.Logf("ACTUAL:   %s (%#v)", actual, actual)
			continue
		}
	}
}

func TestUSDCanonicalForm(t *testing.T) {

	tests := []struct{
		Dollars int32
		Cents   int8
		Expected money.USD
	}{
		{
			Dollars:  0,
			Cents:    0,
			Expected: money.USDCents(0),
		},
		{
			Dollars:  -0,
			Cents:    -0,
			Expected: money.USDCents(0),
		},



		{
			Dollars:  0,
			Cents:    1,
			Expected: money.USDCents(1),
		},
		{
			Dollars:  0,
			Cents:    2,
			Expected: money.USDCents(2),
		},
		{
			Dollars:  0,
			Cents:    3,
			Expected: money.USDCents(3),
		},
		{
			Dollars:  0,
			Cents:    9,
			Expected: money.USDCents(9),
		},
		{
			Dollars:  0,
			Cents:    10,
			Expected: money.USDCents(10),
		},
		{
			Dollars:  0,
			Cents:    15,
			Expected: money.USDCents(15),
		},
		{
			Dollars:  0,
			Cents:    16,
			Expected: money.USDCents(16),
		},
		{
			Dollars:  0,
			Cents:    52,
			Expected: money.USDCents(52),
		},
		{
			Dollars:  0,
			Cents:    99,
			Expected: money.USDCents(99),
		},



		{
			Dollars:  0,
			Cents:    -1,
			Expected: money.USDCents(-1),
		},
		{
			Dollars:  0,
			Cents:    -2,
			Expected: money.USDCents(-2),
		},
		{
			Dollars:  0,
			Cents:    -3,
			Expected: money.USDCents(-3),
		},
		{
			Dollars:  0,
			Cents:    -9,
			Expected: money.USDCents(-9),
		},
		{
			Dollars:  0,
			Cents:    -10,
			Expected: money.USDCents(-10),
		},
		{
			Dollars:  0,
			Cents:    -15,
			Expected: money.USDCents(-15),
		},
		{
			Dollars:  0,
			Cents:    -16,
			Expected: money.USDCents(-16),
		},
		{
			Dollars:  0,
			Cents:    -52,
			Expected: money.USDCents(-52),
		},
		{
			Dollars:  0,
			Cents:    -99,
			Expected: money.USDCents(-99),
		},



		{
			Dollars:  1,
			Cents:    0,
			Expected: money.USDCents(100),
		},
		{
			Dollars:  1,
			Cents:    1,
			Expected: money.USDCents(101),
		},
		{
			Dollars:  1,
			Cents:    2,
			Expected: money.USDCents(102),
		},
		{
			Dollars:  1,
			Cents:    3,
			Expected: money.USDCents(103),
		},
		{
			Dollars:  1,
			Cents:    9,
			Expected: money.USDCents(109),
		},
		{
			Dollars:  1,
			Cents:    10,
			Expected: money.USDCents(110),
		},
		{
			Dollars:  1,
			Cents:    15,
			Expected: money.USDCents(115),
		},
		{
			Dollars:  1,
			Cents:    16,
			Expected: money.USDCents(116),
		},
		{
			Dollars:  1,
			Cents:    52,
			Expected: money.USDCents(152),
		},
		{
			Dollars:  1,
			Cents:    99,
			Expected: money.USDCents(199),
		},



		{
			Dollars:  -1,
			Cents:    -0,
			Expected: money.USDCents(-100),
		},
		{
			Dollars:  -1,
			Cents:    -1,
			Expected: money.USDCents(-101),
		},
		{
			Dollars:  -1,
			Cents:    -2,
			Expected: money.USDCents(-102),
		},
		{
			Dollars:  -1,
			Cents:    -3,
			Expected: money.USDCents(-103),
		},
		{
			Dollars:  -1,
			Cents:    -9,
			Expected: money.USDCents(-109),
		},
		{
			Dollars:  -1,
			Cents:    -10,
			Expected: money.USDCents(-110),
		},
		{
			Dollars:  -1,
			Cents:    -15,
			Expected: money.USDCents(-115),
		},
		{
			Dollars:  -1,
			Cents:    -16,
			Expected: money.USDCents(-116),
		},
		{
			Dollars:  -1,
			Cents:    -52,
			Expected: money.USDCents(-152),
		},
		{
			Dollars:  -1,
			Cents:    -99,
			Expected: money.USDCents(-199),
		},



		{
			Dollars:  12345,
			Cents:    0,
			Expected: money.USDCents(1234500),
		},
		{
			Dollars:  12345,
			Cents:    1,
			Expected: money.USDCents(1234501),
		},
		{
			Dollars:  12345,
			Cents:    2,
			Expected: money.USDCents(1234502),
		},
		{
			Dollars:  12345,
			Cents:    3,
			Expected: money.USDCents(1234503),
		},
		{
			Dollars:  12345,
			Cents:    9,
			Expected: money.USDCents(1234509),
		},
		{
			Dollars:  12345,
			Cents:    10,
			Expected: money.USDCents(1234510),
		},
		{
			Dollars:  12345,
			Cents:    15,
			Expected: money.USDCents(1234515),
		},
		{
			Dollars:  12345,
			Cents:    16,
			Expected: money.USDCents(1234516),
		},
		{
			Dollars:  12345,
			Cents:    52,
			Expected: money.USDCents(1234552),
		},
		{
			Dollars:  12345,
			Cents:    99,
			Expected: money.USDCents(1234599),
		},



		{
			Dollars:  -12345,
			Cents:    -0,
			Expected: money.USDCents(-1234500),
		},
		{
			Dollars:  -12345,
			Cents:    -1,
			Expected: money.USDCents(-1234501),
		},
		{
			Dollars:  -12345,
			Cents:    -2,
			Expected: money.USDCents(-1234502),
		},
		{
			Dollars:  -12345,
			Cents:    -3,
			Expected: money.USDCents(-1234503),
		},
		{
			Dollars:  -12345,
			Cents:    -9,
			Expected: money.USDCents(-1234509),
		},
		{
			Dollars:  -12345,
			Cents:    -10,
			Expected: money.USDCents(-1234510),
		},
		{
			Dollars:  -12345,
			Cents:    -15,
			Expected: money.USDCents(-1234515),
		},
		{
			Dollars:  -12345,
			Cents:    -16,
			Expected: money.USDCents(-1234516),
		},
		{
			Dollars:  -12345,
			Cents:    -52,
			Expected: money.USDCents(-1234552),
		},
		{
			Dollars:  -12345,
			Cents:    -99,
			Expected: money.USDCents(-1234599),
		},
	}

	for testNumber, test := range tests {

		var expected money.USD = test.Expected

		var actual money.USD = money.USDCanonicalForm(test.Dollars, test.Cents)

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("DOLLARS:   %d", test.Dollars)
			t.Logf("CENTS:      %d", test.Cents)
			t.Logf("EXPECTED: %s (%#v)", expected, expected)
			t.Logf("ACTUAL:   %s (%#v)", actual, actual)
			continue
		}
	}
}

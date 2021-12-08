package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCAD_SubCAD(t *testing.T) {

	tests := []struct {
		money    money.CAD
		Expected money.CAD
		money2   money.CAD
	}{
		{
			money2:   money.CADCents(100),
			money:    money.CADCents(300),
			Expected: money.CADCents(200),
		},

		{
			money2:   money.CADCents(-100),
			money:    money.CADCents(5400),
			Expected: money.CADCents(5500),
		},
		{
			money2:   money.CADCents(-100),
			money:    money.CADCents(-5400),
			Expected: money.CADCents(-5300),
		},
		{
			money2:   money.CADCents(0),
			money:    money.CADCents(-5400),
			Expected: money.CADCents(-5400),
		},
	}

	for _, test := range tests {

		if result := test.money.Sub(test.money2); result != test.Expected {
			t.Logf("failed to subtract function Expected %v, Got %v", test.Expected, result)
			t.Fail()
		}
	}
}

func TestCAD_SubUSD(t *testing.T) {

	tests := []struct {
		money    money.USD
		Expected money.USD
		money2   money.USD
	}{
		{
			money2:   money.USDCents(100),
			money:    money.USDCents(300),
			Expected: money.USDCents(200),
		},

		{
			money2:   money.USDCents(-100),
			money:    money.USDCents(5400),
			Expected: money.USDCents(5500),
		},
		{
			money2:   money.USDCents(-100),
			money:    money.USDCents(-5400),
			Expected: money.USDCents(-5300),
		},
		{
			money2:   money.USDCents(0),
			money:    money.USDCents(-5400),
			Expected: money.USDCents(-5400),
		},
	}

	for _, test := range tests {

		if result := test.money.Sub(test.money2); result != test.Expected {
			t.Logf("failed to subtract function Expected %v, Got %v", test.Expected, result)
			t.Fail()
		}
	}
}

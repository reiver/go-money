package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCAD_Abs(t *testing.T) {
	tests := []struct {
		money    money.CAD
		Expected money.CAD
	}{
		{
			money:    money.CADCents(300),
			Expected: money.CADCents(300),
		},

		{
			money:    money.CADCents(5400),
			Expected: money.CADCents(5400),
		},
		{
			money:    money.CADCents(-5400),
			Expected: money.CADCents(5400),
		},
		{
			money:    money.CADCents(0),
			Expected: money.CADCents(0),
		},
	}

	for _, test := range tests {

		if result := test.money.Abs(); result != test.Expected {
			t.Logf("failed with CAD.Abs() method; Expected %v, Got %v", test.Expected, result)
			t.Fail()
		}
	}
}

func TestUSD_Abs(t *testing.T) {
	tests := []struct {
		money    money.USD
		Expected money.USD
	}{
		{
			money:    money.USDCents(300),
			Expected: money.USDCents(300),
		},

		{
			money:    money.USDCents(5400),
			Expected: money.USDCents(5400),
		},
		{
			money:    money.USDCents(-5400),
			Expected: money.USDCents(5400),
		},
		{
			money:    money.USDCents(0),
			Expected: money.USDCents(0),
		},
	}

	for _, test := range tests {

		if result := test.money.Abs(); result != test.Expected {
			t.Logf("failed with USD.Abs() method; Expected %v, Got %v", test.Expected, result)
			t.Fail()
		}
	}
}

package money

import (
	"testing"
)

func TestCAD_Abs(t *testing.T) {
	tests := []struct {
		money    CAD
		Expected CAD
	}{
		{
			money:    CAD(300),
			Expected: CAD(300),
		},

		{
			money:    CAD(5400),
			Expected: CAD(5400),
		},
		{
			money:    CAD(-5400),
			Expected: CAD(5400),
		},
		{
			money:    CAD(0),
			Expected: CAD(0),
		},
	}

	for _, test := range tests {

		if result := test.money.Abs(); result != test.Expected {
			t.Logf("failed to divide function Expected %v, Got %v", test.Expected, result)
			t.Fail()
		}
	}
}

func TestUSD_Abs(t *testing.T) {
	tests := []struct {
		money    USD
		Expected USD
	}{
		{
			money:    USD(300),
			Expected: USD(300),
		},

		{
			money:    USD(5400),
			Expected: USD(5400),
		},
		{
			money:    USD(-5400),
			Expected: USD(5400),
		},
		{
			money:    USD(0),
			Expected: USD(0),
		},
	}

	for _, test := range tests {

		if result := test.money.Abs(); result != test.Expected {
			t.Logf("failed to divide function Expected %v, Got %v", test.Expected, result)
			t.Fail()
		}
	}
}

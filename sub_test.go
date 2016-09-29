package money

import (
	"testing"
)

func TestCAD_SubCAD(t *testing.T) {

	tests := []struct {
		money    CAD
		Expected CAD
		money2   CAD
	}{
		{
			money2:   CAD(100),
			money:    CAD(300),
			Expected: CAD(200),
		},

		{
			money2:   CAD(-100),
			money:    CAD(5400),
			Expected: CAD(5500),
		},
		{
			money2:   CAD(-100),
			money:    CAD(-5400),
			Expected: CAD(-5300),
		},
		{
			money2:   CAD(0),
			money:    CAD(-5400),
			Expected: CAD(-5400),
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
		money    USD
		Expected USD
		money2   USD
	}{
		{
			money2:   USD(100),
			money:    USD(300),
			Expected: USD(200),
		},

		{
			money2:   USD(-100),
			money:    USD(5400),
			Expected: USD(5500),
		},
		{
			money2:   USD(-100),
			money:    USD(-5400),
			Expected: USD(-5300),
		},
		{
			money2:   USD(0),
			money:    USD(-5400),
			Expected: USD(-5400),
		},
	}

	for _, test := range tests {

		if result := test.money.Sub(test.money2); result != test.Expected {
			t.Logf("failed to subtract function Expected %v, Got %v", test.Expected, result)
			t.Fail()
		}
	}
}

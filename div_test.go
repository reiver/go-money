package money

import (
	"math"
	"testing"
)

func TestCAD_FloatDiv(t *testing.T) {

	tests := []struct {
		money    CAD
		Expected float64
		Divider  int64
	}{
		{
			Divider:  2,
			money:    CAD(300),
			Expected: 1.5,
		},

		{
			Divider:  3,
			money:    CAD(5400),
			Expected: 18,
		},
		{
			Divider:  3,
			money:    CAD(-5400),
			Expected: -18,
		},
		{
			Divider:  0,
			money:    CAD(-5400),
			Expected: math.Inf(-1),
		},
		{
			Divider:  0,
			money:    CAD(5400),
			Expected: math.Inf(1),
		},
		{
			Divider:  math.MaxInt64,
			money:    CAD(-5400),
			Expected: -5.854691731421724e-18,
		},
	}

	for _, test := range tests {

		if result := test.money.FloatDiv(test.Divider); result != test.Expected {
			t.Logf("failed to divide function Expected %v, Got %v", test.Expected, result)
			t.Fail()
		}
	}
}

func TestUSD_FloatDiv(t *testing.T) {

	tests := []struct {
		money    USD
		Expected float64
		Divider  int64
	}{
		{
			Divider:  2,
			money:    USD(300),
			Expected: 1.5,
		},

		{
			Divider:  3,
			money:    USD(5400),
			Expected: 18,
		},
		{
			Divider:  3,
			money:    USD(-5400),
			Expected: -18,
		},
		{
			Divider:  0,
			money:    USD(-5400),
			Expected: math.Inf(-1),
		},
		{
			Divider:  0,
			money:    USD(5400),
			Expected: math.Inf(1),
		},
		{
			Divider:  math.MaxInt64,
			money:    USD(-5400),
			Expected: -5.854691731421724e-18,
		},
	}

	for _, test := range tests {

		if result := test.money.FloatDiv(test.Divider); result != test.Expected {
			t.Logf("failed to divide function Expected %v, Got %v", test.Expected, result)
			t.Fail()
		}
	}
}

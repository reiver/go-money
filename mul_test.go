package money

import (
	"math"
	"testing"
)

func TestCAD_MulCAD(t *testing.T) {

	tests := []struct {
		money      CAD
		Expected   CAD
		Multiplier int64
	}{
		{
			Multiplier: 2,
			money:      CAD(300),
			Expected:   CAD(600),
		},

		{
			Multiplier: 3,
			money:      CAD(5400),
			Expected:   CAD(16200),
		},
		{
			Multiplier: 3,
			money:      CAD(-5400),
			Expected:   CAD(-16200),
		},
		{
			Multiplier: 0,
			money:      CAD(-5400),
			Expected:   CAD(0),
		},
		{
			Multiplier: 0,
			money:      CAD(5400),
			Expected:   CAD(0),
		},
		{
			Multiplier: math.MaxInt64,
			money:      CAD(-5400),
			Expected:   CAD(5400),
		},
	}

	for index, test := range tests {

		if result := test.money.Mul(test.Multiplier); result != test.Expected {
			t.Logf("failed test %v to multiplication function Expected %v, Got %v", index, test.Expected, result)
			t.Fail()
		}
	}
}

func TestCAD_MulUSD(t *testing.T) {

	tests := []struct {
		money      USD
		Expected   USD
		Multiplier int64
	}{
		{
			Multiplier: 2,
			money:      USD(300),
			Expected:   USD(600),
		},

		{
			Multiplier: 3,
			money:      USD(5400),
			Expected:   USD(16200),
		},
		{
			Multiplier: 3,
			money:      USD(-5400),
			Expected:   USD(-16200),
		},
		{
			Multiplier: 0,
			money:      USD(-5400),
			Expected:   USD(0),
		},
		{
			Multiplier: 0,
			money:      USD(5400),
			Expected:   USD(0),
		},
		{
			Multiplier: math.MaxInt64,
			money:      USD(-5400),
			Expected:   USD(5400),
		},
	}

	for index, test := range tests {

		if result := test.money.Mul(test.Multiplier); result != test.Expected {
			t.Logf("failed test %v to multiplication function Expected %v, Got %v", index, test.Expected, result)
			t.Fail()
		}
	}
}

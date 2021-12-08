package money_test

import (
	"github.com/reiver/go-money"

	"math"
	"testing"
)

func TestCAD_MulCAD(t *testing.T) {

	tests := []struct {
		money      money.CAD
		Expected   money.CAD
		Multiplier int64
	}{
		{
			Multiplier: 2,
			money:      money.CADCents(300),
			Expected:   money.CADCents(600),
		},

		{
			Multiplier: 3,
			money:      money.CADCents(5400),
			Expected:   money.CADCents(16200),
		},
		{
			Multiplier: 3,
			money:      money.CADCents(-5400),
			Expected:   money.CADCents(-16200),
		},
		{
			Multiplier: 0,
			money:      money.CADCents(-5400),
			Expected:   money.CADCents(0),
		},
		{
			Multiplier: 0,
			money:      money.CADCents(5400),
			Expected:   money.CADCents(0),
		},
		{
			Multiplier: math.MaxInt64,
			money:      money.CADCents(-5400),
			Expected:   money.CADCents(5400),
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
		money      money.USD
		Expected   money.USD
		Multiplier int64
	}{
		{
			Multiplier: 2,
			money:      money.USDCents(300),
			Expected:   money.USDCents(600),
		},

		{
			Multiplier: 3,
			money:      money.USDCents(5400),
			Expected:   money.USDCents(16200),
		},
		{
			Multiplier: 3,
			money:      money.USDCents(-5400),
			Expected:   money.USDCents(-16200),
		},
		{
			Multiplier: 0,
			money:      money.USDCents(-5400),
			Expected:   money.USDCents(0),
		},
		{
			Multiplier: 0,
			money:      money.USDCents(5400),
			Expected:   money.USDCents(0),
		},
		{
			Multiplier: math.MaxInt64,
			money:      money.USDCents(-5400),
			Expected:   money.USDCents(5400),
		},
	}

	for index, test := range tests {

		if result := test.money.Mul(test.Multiplier); result != test.Expected {
			t.Logf("failed test %v to multiplication function Expected %v, Got %v", index, test.Expected, result)
			t.Fail()
		}
	}
}

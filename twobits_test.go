package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCADTwoBits(t *testing.T) {

	var expected money.CAD = money.CADCents(25)
	var actual   money.CAD = money.CADTwoBits()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

func TestUSDTwoBits(t *testing.T) {

	var expected money.USD = money.USDCents(25)
	var actual   money.USD = money.USDTwoBits()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCADEightBits(t *testing.T) {

	var expected money.CAD = money.CADCents(100)
	var actual   money.CAD = money.CADEightBits()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

func TestUSDEightBits(t *testing.T) {

	var expected money.USD = money.USDCents(100)
	var actual   money.USD = money.USDEightBits()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCADFourBits(t *testing.T) {

	var expected money.CAD = money.CADCents(50)
	var actual   money.CAD = money.CADFourBits()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

func TestUSDFourBits(t *testing.T) {

	var expected money.USD = money.USDCents(50)
	var actual   money.USD = money.USDFourBits()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

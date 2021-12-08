package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCADSixBits(t *testing.T) {

	var expected money.CAD = money.CADCents(75)
	var actual   money.CAD = money.CADSixBits()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

func TestUSDSixBits(t *testing.T) {

	var expected money.USD = money.USDCents(75)
	var actual   money.USD = money.USDSixBits()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

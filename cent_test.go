package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCADCent(t *testing.T) {

	var expected money.CAD = money.CADCents(1)
	var actual   money.CAD = money.CADCent()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

func TestUSDCent(t *testing.T) {

	var expected money.USD = money.USDCents(1)
	var actual   money.USD = money.USDCent()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

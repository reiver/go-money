package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCADQuarter(t *testing.T) {

	var expected money.CAD = money.CADCents(25)
	var actual   money.CAD = money.CADQuarter()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

func TestUSDQuarter(t *testing.T) {

	var expected money.USD = money.USDCents(25)
	var actual   money.USD = money.USDQuarter()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

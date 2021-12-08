package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCADNickel(t *testing.T) {

	var expected money.CAD = money.CADCents(5)
	var actual   money.CAD = money.CADNickel()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

func TestUSDNickel(t *testing.T) {

	var expected money.USD = money.USDCents(5)
	var actual   money.USD = money.USDNickel()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

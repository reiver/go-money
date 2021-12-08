package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestCADDime(t *testing.T) {

	var expected money.CAD = money.CADCents(10)
	var actual   money.CAD = money.CADDime()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

func TestUSDDime(t *testing.T) {

	var expected money.USD = money.USDCents(10)
	var actual   money.USD = money.USDDime()

	if expected != actual {
		t.Error("The actual value is not equal to the expected value.")
		t.Logf("EXPECTED: %q (%#v)", expected, expected)
		t.Logf("ACTUAL:   %q (%#v)", actual, actual)
		return
	}
}

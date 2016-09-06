package money

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCents(t *testing.T) {
    tests := []struct {
        money string
        expect int64
    }{
        { "100.00", 10000, },
        { "0.00", 0, },
        { "0.01", 1, },
        { "0.55", 55, },
        { "1.11", 111, },
        { "-25.52", -2552, },
    }

    for _, test := range tests {
        cad, _ := ParseCAD(test.money)
        assert.Equal(t, test.expect, cad.Cents())

        usd, _ := ParseUSD(test.money)
        assert.Equal(t, test.expect, usd.Cents())
    }
}

func TestDollars(t *testing.T) {
    tests := []struct {
        money string
        expect int64
    }{
        { "100.01", 100, },
        { "0.25", 0, },
        { "11.10", 11, },
        { "11.49", 11, },
        { "11.50", 12, },
        { "11.99", 12, },
        { "-25.50", -26, },
        { "-25.49", -25, },
        { "25.50", 26, },
        { "25.49", 25, },
    }

    for _, test := range tests {
        cad, _ := ParseCAD(test.money)
        assert.Equal(t, test.expect, cad.Dollars(), "Expected %s to be %d Dollars", test.money, test.expect)

        usd, _ := ParseUSD(test.money)
        assert.Equal(t, test.expect, usd.Dollars(), "Expected %s to be %d Dollars", test.money, test.expect)
    }
}
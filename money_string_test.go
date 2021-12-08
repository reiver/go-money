package money_test

import (
	"github.com/reiver/go-money"

	"testing"
)

func TestStringCAD(t *testing.T) {

	tests := []struct{
		Money    money.CAD
		Expected string
	}{
		{
			Money:    money.CADCents(0),
			Expected: "$0.00",
		},
		{
			Money:    money.CADDollars(0),
			Expected: "$0.00",
		},
		{
			Money:    money.CADCent().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.CADNickel().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.CADDime().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.CADQuarter().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.CADTwoBits().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.CADFourBits().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.CADDollar().Mul(0),
			Expected: "$0.00",
		},



		{
			Money:    money.CADCent(),
			Expected: "$0.01",
		},
		{
			Money:    money.CADNickel(),
			Expected: "$0.05",
		},
		{
			Money:    money.CADDime(),
			Expected: "$0.10",
		},
		{
			Money:    money.CADQuarter(),
			Expected: "$0.25",
		},
		{
			Money:    money.CADTwoBits(),
			Expected: "$0.25",
		},
		{
			Money:    money.CADFourBits(),
			Expected: "$0.50",
		},
		{
			Money:    money.CADDollar(),
			Expected: "$1.00",
		},



		{
			Money: money.CADCents(1),
			Expected: "$0.01",
		},
		{
			Money: money.CADCents(2),
			Expected: "$0.02",
		},
		{
			Money: money.CADCents(3),
			Expected: "$0.03",
		},
		{
			Money: money.CADCents(4),
			Expected: "$0.04",
		},
		{
			Money: money.CADCents(5),
			Expected: "$0.05",
		},
		{
			Money: money.CADCents(10),
			Expected: "$0.10",
		},
		{
			Money: money.CADCents(12345),
			Expected: "$123.45",
		},



		{
			Money: money.CADCents(-1),
			Expected: "-$0.01",
		},
		{
			Money: money.CADCents(-2),
			Expected: "-$0.02",
		},
		{
			Money: money.CADCents(-3),
			Expected: "-$0.03",
		},
		{
			Money: money.CADCents(-4),
			Expected: "-$0.04",
		},
		{
			Money: money.CADCents(-5),
			Expected: "-$0.05",
		},
		{
			Money: money.CADCents(-10),
			Expected: "-$0.10",
		},
		{
			Money: money.CADCents(-12345),
			Expected: "-$123.45",
		},



		{
			Money: money.CADDollars(1),
			Expected: "$1.00",
		},
		{
			Money: money.CADDollars(2),
			Expected: "$2.00",
		},
		{
			Money: money.CADDollars(3),
			Expected: "$3.00",
		},
		{
			Money: money.CADDollars(4),
			Expected: "$4.00",
		},
		{
			Money: money.CADDollars(5),
			Expected: "$5.00",
		},
		{
			Money: money.CADDollars(10),
			Expected: "$10.00",
		},
		{
			Money: money.CADDollars(12345),
			Expected: "$12345.00",
		},



		{
			Money: money.CADDollars(-1),
			Expected: "-$1.00",
		},
		{
			Money: money.CADDollars(-2),
			Expected: "-$2.00",
		},
		{
			Money: money.CADDollars(-3),
			Expected: "-$3.00",
		},
		{
			Money: money.CADDollars(-4),
			Expected: "-$4.00",
		},
		{
			Money: money.CADDollars(-5),
			Expected: "-$5.00",
		},
		{
			Money: money.CADDollars(-10),
			Expected: "-$10.00",
		},
		{
			Money: money.CADDollars(-12345),
			Expected: "-$12345.00",
		},
	}


	for testNumber, test := range tests {

		actual := test.Money.String()

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, Expected %q but actually got %q.", testNumber, expected, actual)
		}
	}
}

func TestStringUSD(t *testing.T) {

	tests := []struct{
		Money    money.USD
		Expected string
	}{
		{
			Money:    money.USDCents(0),
			Expected: "$0.00",
		},
		{
			Money:    money.USDDollars(0),
			Expected: "$0.00",
		},
		{
			Money:    money.USDCent().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.USDNickel().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.USDDime().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.USDQuarter().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.USDTwoBits().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.USDFourBits().Mul(0),
			Expected: "$0.00",
		},
		{
			Money:    money.USDDollar().Mul(0),
			Expected: "$0.00",
		},



		{
			Money:    money.USDCent(),
			Expected: "$0.01",
		},
		{
			Money:    money.USDNickel(),
			Expected: "$0.05",
		},
		{
			Money:    money.USDDime(),
			Expected: "$0.10",
		},
		{
			Money:    money.USDQuarter(),
			Expected: "$0.25",
		},
		{
			Money:    money.USDTwoBits(),
			Expected: "$0.25",
		},
		{
			Money:    money.USDFourBits(),
			Expected: "$0.50",
		},
		{
			Money:    money.USDDollar(),
			Expected: "$1.00",
		},



		{
			Money: money.USDCents(1),
			Expected: "$0.01",
		},
		{
			Money: money.USDCents(2),
			Expected: "$0.02",
		},
		{
			Money: money.USDCents(3),
			Expected: "$0.03",
		},
		{
			Money: money.USDCents(4),
			Expected: "$0.04",
		},
		{
			Money: money.USDCents(5),
			Expected: "$0.05",
		},
		{
			Money: money.USDCents(10),
			Expected: "$0.10",
		},
		{
			Money: money.USDCents(12345),
			Expected: "$123.45",
		},



		{
			Money: money.USDCents(-1),
			Expected: "-$0.01",
		},
		{
			Money: money.USDCents(-2),
			Expected: "-$0.02",
		},
		{
			Money: money.USDCents(-3),
			Expected: "-$0.03",
		},
		{
			Money: money.USDCents(-4),
			Expected: "-$0.04",
		},
		{
			Money: money.USDCents(-5),
			Expected: "-$0.05",
		},
		{
			Money: money.USDCents(-10),
			Expected: "-$0.10",
		},
		{
			Money: money.USDCents(-12345),
			Expected: "-$123.45",
		},



		{
			Money: money.USDDollars(1),
			Expected: "$1.00",
		},
		{
			Money: money.USDDollars(2),
			Expected: "$2.00",
		},
		{
			Money: money.USDDollars(3),
			Expected: "$3.00",
		},
		{
			Money: money.USDDollars(4),
			Expected: "$4.00",
		},
		{
			Money: money.USDDollars(5),
			Expected: "$5.00",
		},
		{
			Money: money.USDDollars(10),
			Expected: "$10.00",
		},
		{
			Money: money.USDDollars(12345),
			Expected: "$12345.00",
		},



		{
			Money: money.USDDollars(-1),
			Expected: "-$1.00",
		},
		{
			Money: money.USDDollars(-2),
			Expected: "-$2.00",
		},
		{
			Money: money.USDDollars(-3),
			Expected: "-$3.00",
		},
		{
			Money: money.USDDollars(-4),
			Expected: "-$4.00",
		},
		{
			Money: money.USDDollars(-5),
			Expected: "-$5.00",
		},
		{
			Money: money.USDDollars(-10),
			Expected: "-$10.00",
		},
		{
			Money: money.USDDollars(-12345),
			Expected: "-$12345.00",
		},
	}


	for testNumber, test := range tests {

		actual := test.Money.String()

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, Expected %q but actually got %q.", testNumber, expected, actual)
		}
	}
}

# go-money

Package money provides a way of dealing with money in type safe way, including parsing from strings.

So, for example, the type of U.S. Dollars is `money.USD` and the type for Canadian Dollars is `money.CAD`.
These two types are completely different, and cannot be assigned to each other.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-money

[![GoDoc](https://godoc.org/github.com/reiver/go-money?status.svg)](https://godoc.org/github.com/reiver/go-money)

## Parsing

To turn a money amount represented as a string into a money type, do something like the follow.

	canadianMoney, err := money.ParseCAD("$123.45")

And:
	usMoney, err := money.ParseUSD("$123.45")

## Initializing

There are a number of functions available to initalize the value of a `money.USD` or a `money.CAD`:

	twentyThreeCanadianCents := money.CADCents(23)
	twentyThreeUSCents       := money.USDCents(23)

	fiftyThreeCanadianDollars := money.CADCents(53)
	fiftyThreeUSDollars       := money.USDCents(53)

## Common Money Amounts

There are some convenience functions for common money amounts.

For Canadian Dollars there is:

	oneCentCAD            := money.CADCent()

	fiveCentsCAD          := money.CADNickel()

	tenCentsCAD           := money.CADDime()

	twentyFiveCentsCAD    := money.CADQuarter()

	twentyFiveCentsCADToo := money.CADTwoBits()

	fiftyCentsCAD         := money.CADFourBits()

	seventyFiveCentsCAD   := money.CADSixBits()

	oneHundredCentsCAD    := money.CADEightBits()

	oneHundredCentsCADToo := money.CADDollar()

And for U.S. Dollars there is:

	oneCentUSD            := money.USDCent()

	fiveCentsUSD          := money.USDNickel()

	tenCentsUSD           := money.USDDime()

	twentyFiveCentsUSD    := money.USDQuarter()

	twentyFiveCentsUSDToo := money.USDTwoBits()

	fiftyCentsUSD         := money.USDFourBits()

	seventyFiveCentsUSD   := money.USDSixBits()

	oneHundredCentsUSD    := money.USDEightBits()

	oneHundredCentsUSDToo := money.USDDollar()

## Formatting

To turn a money amount into a specific string format, do something like the following.

	strCAD := canadianMoney.Format("")           // $123.45

	strCAD := canadianMoney.Format("$")          // $123.45

	strCAD := canadianMoney.Format("\uFE69")     // ﹩123.45
	strCAD := canadianMoney.Format("﹩")          // ﹩123.45

	strCAD := canadianMoney.Format("\uFF04")     // ＄123.45
	strCAD := canadianMoney.Format("＄")         // ＄123.45

	strCAD := canadianMoney.Format("\U0001F4B2") // 💲123.45
	strCAD := canadianMoney.Format("💲")         // 💲123.45

	strCAD := canadianMoney.Format("\u00A2")     // 12345¢
	strCAD := canadianMoney.Format("¢")          // 12345¢

	strCAD := canadianMoney.Format("\uFFE0")     // 12345￠
	strCAD := canadianMoney.Format("￠")         // 12345￠


	strUSD := usMoney.Format("")           // $123.45

	strUSD := usMoney.Format("$")          // $123.45

	strUSD := usMoney.Format("\uFE69")     // ﹩123.45
	strUSD := usMoney.Format("﹩")          // ﹩123.45

	strUSD := usMoney.Format("\uFF04")     // ＄123.45
	strUSD := usMoney.Format("＄")         // ＄123.45

	strUSD := usMoney.Format("\U0001F4B2") // 💲123.45
	strUSD := usMoney.Format("💲")         // 💲123.45

	strUSD := usMoney.Format("\u00A2")     // 12345¢
	strUSD := usMoney.Format("¢")          // 12345¢

	strUSD := usMoney.Format("\uFFE0")     // 12345￠
	strUSD := usMoney.Format("￠")         // 12345￠

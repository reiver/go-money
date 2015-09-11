/*
Package money provides a way of dealing with money in type safe way, including parsing from strings.

So, for example, the type of U.S. Dollars is money.USD and the type for Canadian Dollars is money.CAD.
These  two types are completely different, and cannot be assigned to each other.

Parsing

To turn a money amount represented as a string into a money type, do something like the follow.

	canadianMoney, err := money.ParseCAD("$123.45")

	usMoney, err := money.ParseUSD("$123.45")


Formatting

To turn a money amount into a specific string format, do something like the following.

	strCAD := canadianMoney.Format("")           // $123.45

	strCAD := canadianMoney.Format("$")          // $123.45

	strCAD := canadianMoney.Format("\uFE69")     // ﹩123.45
	strCAD := canadianMoney.Format("﹩")          // ﹩123.45

	strCAD := canadianMoney.Format("\uFF04")     // ＄123.45
	strCAD := canadianMoney.Format("＄")         // ＄123.45

	strCAD := canadianMoney.Format("\U0001F4B2") // 💲123.45
	strCAD := canadianMoney.Format("💲")         // 💲123.45

	strCAD := canadianMoney.Format("\u00A2")     // 123.45¢
	strCAD := canadianMoney.Format("¢")          // 123.45¢

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


Constants

Constants exist that can be used a money "literals". For example:

	fiveCanadianDollars := 5 * money.CAD_DOLLAR

	fiveUsDollars := 5 * money.USD_DOLLAR

	twentyCanadianCents := 20 * money.CAD_CENT

	twentyUsCents := 20 * money.USD_CENT
*/
package money

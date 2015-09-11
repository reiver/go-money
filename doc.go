/*
Package money provides a way of dealing with money in type safe way, including parsing from strings.

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

	fiveCanadianDollars := 5 * CAD_DOLLAR

	fiveUsDollars := 5 * USD_DOLLAR

	twentyCanadianCents := 20 * CAD_CENT

	twentyUsCents := 20 * USD_CENT
*/
package money

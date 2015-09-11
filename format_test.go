package money


import (
	"testing"
)


func TestFormatCAD(t *testing.T) {

	tests := []struct{
		Money    CAD
		Format   string
		Expecteds map[string]string // The key is the layout and the value is what is expected.
	}{
		{
			Money:                           1 * CAD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$1.00",
				"$":                   "$1.00",
				"\uFE69":         "\uFE691.00",
				"\uFF04":         "\uFF041.00",
				"\U0001F4B2": "\U0001F4B21.00",
				"\u00A2":               "100\u00A2",
				"\uFFE0":               "100\uFFE0",
			},
		},
		{
			Money:                           2 * CAD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$2.00",
				"$":                   "$2.00",
				"\uFE69":         "\uFE692.00",
				"\uFF04":         "\uFF042.00",
				"\U0001F4B2": "\U0001F4B22.00",
				"\u00A2":               "200\u00A2",
				"\uFFE0":               "200\uFFE0",
			},
		},
		{
			Money:                           3 * CAD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$3.00",
				"$":                   "$3.00",
				"\uFE69":         "\uFE693.00",
				"\uFF04":         "\uFF043.00",
				"\U0001F4B2": "\U0001F4B23.00",
				"\u00A2":               "300\u00A2",
				"\uFFE0":               "300\uFFE0",
			},
		},
		{
			Money:                           4 * CAD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$4.00",
				"$":                   "$4.00",
				"\uFE69":         "\uFE694.00",
				"\uFF04":         "\uFF044.00",
				"\U0001F4B2": "\U0001F4B24.00",
				"\u00A2":               "400\u00A2",
				"\uFFE0":               "400\uFFE0",
			},
		},
		{
			Money:                           5 * CAD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$5.00",
				"$":                   "$5.00",
				"\uFE69":         "\uFE695.00",
				"\uFF04":         "\uFF045.00",
				"\U0001F4B2": "\U0001F4B25.00",
				"\u00A2":               "500\u00A2",
				"\uFFE0":               "500\uFFE0",
			},
		},
		{
			Money:                           10 * CAD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$10.00",
				"$":                   "$10.00",
				"\uFE69":         "\uFE6910.00",
				"\uFF04":         "\uFF0410.00",
				"\U0001F4B2": "\U0001F4B210.00",
				"\u00A2":               "1000\u00A2",
				"\uFFE0":               "1000\uFFE0",
			},
		},
		{
			Money:                           12345 * CAD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$12345.00",
				"$":                   "$12345.00",
				"\uFE69":         "\uFE6912345.00",
				"\uFF04":         "\uFF0412345.00",
				"\U0001F4B2": "\U0001F4B212345.00",
				"\u00A2":               "1234500\u00A2",
				"\uFFE0":               "1234500\uFFE0",
			},
		},
		{
			Money:                           1234 * CAD_CENT,
			Expecteds: map[string]string{
				"":                    "$12.34",
				"$":                   "$12.34",
				"\uFE69":         "\uFE6912.34",
				"\uFF04":         "\uFF0412.34",
				"\U0001F4B2": "\U0001F4B212.34",
				"\u00A2":               "1234\u00A2",
				"\uFFE0":               "1234\uFFE0",
			},
		},
		{
			Money:                           101 * CAD_CENT,
			Expecteds: map[string]string{
				"":                    "$1.01",
				"$":                   "$1.01",
				"\uFE69":         "\uFE691.01",
				"\uFF04":         "\uFF041.01",
				"\U0001F4B2": "\U0001F4B21.01",
				"\u00A2":               "101\u00A2",
				"\uFFE0":               "101\uFFE0",
			},
		},
		{
			Money:                              1 * CAD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.01",
				"$":                   "$0.01",
				"\uFE69":         "\uFE690.01",
				"\uFF04":         "\uFF040.01",
				"\U0001F4B2": "\U0001F4B20.01",
				"\u00A2":                  "1\u00A2",
				"\uFFE0":                  "1\uFFE0",
			},
		},
		{
			Money:                              2 * CAD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.02",
				"$":                   "$0.02",
				"\uFE69":         "\uFE690.02",
				"\uFF04":         "\uFF040.02",
				"\U0001F4B2": "\U0001F4B20.02",
				"\u00A2":                  "2\u00A2",
				"\uFFE0":                  "2\uFFE0",
			},
		},
		{
			Money:                              3 * CAD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.03",
				"$":                   "$0.03",
				"\uFE69":         "\uFE690.03",
				"\uFF04":         "\uFF040.03",
				"\U0001F4B2": "\U0001F4B20.03",
				"\u00A2":                  "3\u00A2",
				"\uFFE0":                  "3\uFFE0",
			},
		},
		{
			Money:                              4 * CAD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.04",
				"$":                   "$0.04",
				"\uFE69":         "\uFE690.04",
				"\uFF04":         "\uFF040.04",
				"\U0001F4B2": "\U0001F4B20.04",
				"\u00A2":                  "4\u00A2",
				"\uFFE0":                  "4\uFFE0",
			},
		},
		{
			Money:                              5 * CAD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.05",
				"$":                   "$0.05",
				"\uFE69":         "\uFE690.05",
				"\uFF04":         "\uFF040.05",
				"\U0001F4B2": "\U0001F4B20.05",
				"\u00A2":                  "5\u00A2",
				"\uFFE0":                  "5\uFFE0",
			},
		},
		{
			Money:                             10 * CAD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.10",
				"$":                   "$0.10",
				"\uFE69":         "\uFE690.10",
				"\uFF04":         "\uFF040.10",
				"\U0001F4B2": "\U0001F4B20.10",
				"\u00A2":                 "10\u00A2",
				"\uFFE0":                 "10\uFFE0",
			},
		},
		{
			Money:                           12345 * CAD_CENT,
			Expecteds: map[string]string{
				"":                    "$123.45",
				"$":                   "$123.45",
				"\uFE69":         "\uFE69123.45",
				"\uFF04":         "\uFF04123.45",
				"\U0001F4B2": "\U0001F4B2123.45",
				"\u00A2":               "12345\u00A2",
				"\uFFE0":               "12345\uFFE0",
			},
		},
		{
			Money:                           1010 * CAD_CENT,
			Expecteds: map[string]string{
				"":                    "$10.10",
				"$":                   "$10.10",
				"\uFE69":         "\uFE6910.10",
				"\uFF04":         "\uFF0410.10",
				"\U0001F4B2": "\U0001F4B210.10",
				"\u00A2":               "1010\u00A2",
				"\uFFE0":               "1010\uFFE0",
			},
		},
	}


	for testNumber, test := range tests {

		for layout, expected := range test.Expecteds {
			actual := test.Money.Format(layout)

			if expected != actual {
				t.Errorf("For test #%d, for layout %q expected %q but instead got %q.", testNumber, layout, expected, actual)
			}
		}
	}
}


func TestFormatUSD(t *testing.T) {

	tests := []struct{
		Money    USD
		Format   string
		Expecteds map[string]string // The key is the layout and the value is what is expected.
	}{
		{
			Money:                           1 * USD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$1.00",
				"$":                   "$1.00",
				"\uFE69":         "\uFE691.00",
				"\uFF04":         "\uFF041.00",
				"\U0001F4B2": "\U0001F4B21.00",
				"\u00A2":               "100\u00A2",
				"\uFFE0":               "100\uFFE0",
			},
		},
		{
			Money:                           2 * USD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$2.00",
				"$":                   "$2.00",
				"\uFE69":         "\uFE692.00",
				"\uFF04":         "\uFF042.00",
				"\U0001F4B2": "\U0001F4B22.00",
				"\u00A2":               "200\u00A2",
				"\uFFE0":               "200\uFFE0",
			},
		},
		{
			Money:                           3 * USD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$3.00",
				"$":                   "$3.00",
				"\uFE69":         "\uFE693.00",
				"\uFF04":         "\uFF043.00",
				"\U0001F4B2": "\U0001F4B23.00",
				"\u00A2":               "300\u00A2",
				"\uFFE0":               "300\uFFE0",
			},
		},
		{
			Money:                           4 * USD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$4.00",
				"$":                   "$4.00",
				"\uFE69":         "\uFE694.00",
				"\uFF04":         "\uFF044.00",
				"\U0001F4B2": "\U0001F4B24.00",
				"\u00A2":               "400\u00A2",
				"\uFFE0":               "400\uFFE0",
			},
		},
		{
			Money:                           5 * USD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$5.00",
				"$":                   "$5.00",
				"\uFE69":         "\uFE695.00",
				"\uFF04":         "\uFF045.00",
				"\U0001F4B2": "\U0001F4B25.00",
				"\u00A2":               "500\u00A2",
				"\uFFE0":               "500\uFFE0",
			},
		},
		{
			Money:                           10 * USD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$10.00",
				"$":                   "$10.00",
				"\uFE69":         "\uFE6910.00",
				"\uFF04":         "\uFF0410.00",
				"\U0001F4B2": "\U0001F4B210.00",
				"\u00A2":               "1000\u00A2",
				"\uFFE0":               "1000\uFFE0",
			},
		},
		{
			Money:                           12345 * USD_DOLLAR,
			Expecteds: map[string]string{
				"":                    "$12345.00",
				"$":                   "$12345.00",
				"\uFE69":         "\uFE6912345.00",
				"\uFF04":         "\uFF0412345.00",
				"\U0001F4B2": "\U0001F4B212345.00",
				"\u00A2":               "1234500\u00A2",
				"\uFFE0":               "1234500\uFFE0",
			},
		},
		{
			Money:                           1234 * USD_CENT,
			Expecteds: map[string]string{
				"":                    "$12.34",
				"$":                   "$12.34",
				"\uFE69":         "\uFE6912.34",
				"\uFF04":         "\uFF0412.34",
				"\U0001F4B2": "\U0001F4B212.34",
				"\u00A2":               "1234\u00A2",
				"\uFFE0":               "1234\uFFE0",
			},
		},
		{
			Money:                           101 * USD_CENT,
			Expecteds: map[string]string{
				"":                    "$1.01",
				"$":                   "$1.01",
				"\uFE69":         "\uFE691.01",
				"\uFF04":         "\uFF041.01",
				"\U0001F4B2": "\U0001F4B21.01",
				"\u00A2":               "101\u00A2",
				"\uFFE0":               "101\uFFE0",
			},
		},
		{
			Money:                              1 * USD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.01",
				"$":                   "$0.01",
				"\uFE69":         "\uFE690.01",
				"\uFF04":         "\uFF040.01",
				"\U0001F4B2": "\U0001F4B20.01",
				"\u00A2":                  "1\u00A2",
				"\uFFE0":                  "1\uFFE0",
			},
		},
		{
			Money:                              2 * USD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.02",
				"$":                   "$0.02",
				"\uFE69":         "\uFE690.02",
				"\uFF04":         "\uFF040.02",
				"\U0001F4B2": "\U0001F4B20.02",
				"\u00A2":                  "2\u00A2",
				"\uFFE0":                  "2\uFFE0",
			},
		},
		{
			Money:                              3 * USD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.03",
				"$":                   "$0.03",
				"\uFE69":         "\uFE690.03",
				"\uFF04":         "\uFF040.03",
				"\U0001F4B2": "\U0001F4B20.03",
				"\u00A2":                  "3\u00A2",
				"\uFFE0":                  "3\uFFE0",
			},
		},
		{
			Money:                              4 * USD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.04",
				"$":                   "$0.04",
				"\uFE69":         "\uFE690.04",
				"\uFF04":         "\uFF040.04",
				"\U0001F4B2": "\U0001F4B20.04",
				"\u00A2":                  "4\u00A2",
				"\uFFE0":                  "4\uFFE0",
			},
		},
		{
			Money:                              5 * USD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.05",
				"$":                   "$0.05",
				"\uFE69":         "\uFE690.05",
				"\uFF04":         "\uFF040.05",
				"\U0001F4B2": "\U0001F4B20.05",
				"\u00A2":                  "5\u00A2",
				"\uFFE0":                  "5\uFFE0",
			},
		},
		{
			Money:                             10 * USD_CENT,
			Expecteds: map[string]string{
				"":                    "$0.10",
				"$":                   "$0.10",
				"\uFE69":         "\uFE690.10",
				"\uFF04":         "\uFF040.10",
				"\U0001F4B2": "\U0001F4B20.10",
				"\u00A2":                 "10\u00A2",
				"\uFFE0":                 "10\uFFE0",
			},
		},
		{
			Money:                           12345 * USD_CENT,
			Expecteds: map[string]string{
				"":                    "$123.45",
				"$":                   "$123.45",
				"\uFE69":         "\uFE69123.45",
				"\uFF04":         "\uFF04123.45",
				"\U0001F4B2": "\U0001F4B2123.45",
				"\u00A2":               "12345\u00A2",
				"\uFFE0":               "12345\uFFE0",
			},
		},
		{
			Money:                           1010 * USD_CENT,
			Expecteds: map[string]string{
				"":                    "$10.10",
				"$":                   "$10.10",
				"\uFE69":         "\uFE6910.10",
				"\uFF04":         "\uFF0410.10",
				"\U0001F4B2": "\U0001F4B210.10",
				"\u00A2":               "1010\u00A2",
				"\uFFE0":               "1010\uFFE0",
			},
		},
	}


	for testNumber, test := range tests {

		for layout, expected := range test.Expecteds {
			actual := test.Money.Format(layout)

			if expected != actual {
				t.Errorf("For test #%d, for layout %q expected %q but instead got %q.", testNumber, layout, expected, actual)
			}
		}
	}
}

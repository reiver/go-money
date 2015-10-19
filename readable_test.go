package money


import (
	"testing"
)


func TestReadable(t *testing.T) {

	tests := []struct{
		Number   int64
		Expecteds map[rune]string
	}{
		{
			Number:    0,
			Expecteds: map[rune]string{
				',':"0",
				'.':"0",
				' ':"0",
				'_':"0",
			},
		},
		{
			Number:      1,
			Expecteds: map[rune]string{
				',':"1",
				'.':"1",
				' ':"1",
				'_':"1",
			},
		},
		{
			Number:      2,
			Expecteds: map[rune]string{
				',':"2",
				'.':"2",
				' ':"2",
				'_':"2",
			},
		},
		{
			Number:      3,
			Expecteds: map[rune]string{
				',':"3",
				'.':"3",
				' ':"3",
				'_':"3",
			},
		},
		{
			Number:      4,
			Expecteds: map[rune]string{
				',':"4",
				'.':"4",
				' ':"4",
				'_':"4",
			},
		},
		{
			Number:      5,
			Expecteds: map[rune]string{
				',':"5",
				'.':"5",
				' ':"5",
				'_':"5",
			},
		},
		{
			Number:      6,
			Expecteds: map[rune]string{
				',':"6",
				'.':"6",
				' ':"6",
				'_':"6",
			},
		},
		{
			Number:      7,
			Expecteds: map[rune]string{
				',':"7",
				'.':"7",
				' ':"7",
				'_':"7",
			},
		},
		{
			Number:      8,
			Expecteds: map[rune]string{
				',':"8",
				'.':"8",
				' ':"8",
				'_':"8",
			},
		},
		{
			Number:      9,
			Expecteds: map[rune]string{
				',':"9",
				'.':"9",
				' ':"9",
				'_':"9",
			},
		},
		{
			Number:      10,
			Expecteds: map[rune]string{
				',':"10",
				'.':"10",
				' ':"10",
				'_':"10",
			},
		},

		{
			Number:      42,
			Expecteds: map[rune]string{
				',':"42",
				'.':"42",
				' ':"42",
				'_':"42",
			},
		},

		{
			Number:      89,
			Expecteds: map[rune]string{
				',':"89",
				'.':"89",
				' ':"89",
				'_':"89",
			},
		},

		{
			Number:      99,
			Expecteds: map[rune]string{
				',':"99",
				'.':"99",
				' ':"99",
				'_':"99",
			},
		},

		{
			Number:      100,
			Expecteds: map[rune]string{
				',':"100",
				'.':"100",
				' ':"100",
				'_':"100",
			},
		},

		{
			Number:      123,
			Expecteds: map[rune]string{
				',':"123",
				'.':"123",
				' ':"123",
				'_':"123",
			},
		},

		{
			Number:      543,
			Expecteds: map[rune]string{
				',':"543",
				'.':"543",
				' ':"543",
				'_':"543",
			},
		},

		{
			Number:      999,
			Expecteds: map[rune]string{
				',':"999",
				'.':"999",
				' ':"999",
				'_':"999",
			},
		},


		{
			Number:      1000,
			Expecteds: map[rune]string{
				',':"1,000",
				'.':"1.000",
				' ':"1 000",
				'_':"1_000",
			},
		},

		{
			Number:      1234,
			Expecteds: map[rune]string{
				',':"1,234",
				'.':"1.234",
				' ':"1 234",
				'_':"1_234",
			},
		},

		{
			Number:      12345,
			Expecteds: map[rune]string{
				',':"12,345",
				'.':"12.345",
				' ':"12 345",
				'_':"12_345",
			},
		},

		{
			Number:      123456,
			Expecteds: map[rune]string{
				',':"123,456",
				'.':"123.456",
				' ':"123 456",
				'_':"123_456",
			},
		},

		{
			Number:      1234567,
			Expecteds: map[rune]string{
				',':"1,234,567",
				'.':"1.234.567",
				' ':"1 234 567",
				'_':"1_234_567",
			},
		},

		{
			Number:      12345678,
			Expecteds: map[rune]string{
				',':"12,345,678",
				'.':"12.345.678",
				' ':"12 345 678",
				'_':"12_345_678",
			},
		},

		{
			Number:      123456789,
			Expecteds: map[rune]string{
				',':"123,456,789",
				'.':"123.456.789",
				' ':"123 456 789",
				'_':"123_456_789",
			},
		},

		{
			Number:      1234567890,
			Expecteds: map[rune]string{
				',':"1,234,567,890",
				'.':"1.234.567.890",
				' ':"1 234 567 890",
				'_':"1_234_567_890",
			},
		},

		{
			Number:      12345678901,
			Expecteds: map[rune]string{
				',':"12,345,678,901",
				'.':"12.345.678.901",
				' ':"12 345 678 901",
				'_':"12_345_678_901",
			},
		},

		{
			Number:      123456789012,
			Expecteds: map[rune]string{
				',':"123,456,789,012",
				'.':"123.456.789.012",
				' ':"123 456 789 012",
				'_':"123_456_789_012",
			},
		},

		{
			Number:      1234567890123,
			Expecteds: map[rune]string{
				',':"1,234,567,890,123",
				'.':"1.234.567.890.123",
				' ':"1 234 567 890 123",
				'_':"1_234_567_890_123",
			},
		},
	}


	for testNumber, test := range tests {

		for sep, expected := range test.Expecteds {
			actual := readable(test.Number, sep)

			if expected != actual {
				t.Errorf("For test #%d with separator %q, expected %q, but actually got %q for number %d.", testNumber, string(sep), expected, actual, test.Number)
			}
		}

	}
}

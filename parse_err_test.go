package money


import (
	"testing"

	"math/rand"
	"time"
)


func generateTests() []struct{String string} {
	// Initialize.
	randomness := rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )

	// We want to make sure we test the empty string!
	tests := []struct{
		String   string
	}{
		{
			String:   "",
		},
	}

	// Create some tests with just white space.
	//
	// Note we are only using the "traditional" (i.e., ASCII)
	// whitespace characters here.
	//
	//	\t             horizontal tab (== \011)
	//	\n             newline (== \012)
	//	\v             vertical tab character (== \013)
	//	\f             form feed (== \014)
	//	\r             carriage return (== \015)
	//
	whitespaceCharacters := []rune("\t\n\v\f\r")

	for _,character := range whitespaceCharacters {
		datum := struct{String string}{String: string(character)}

		tests = append(tests, datum)
	}

	for _,character := range whitespaceCharacters {
		datum := struct{String string}{String: string([]rune{character, character})}

		tests = append(tests, datum)
	}

	for _,character := range whitespaceCharacters {
		datum := struct{String string}{String: string([]rune{character, character, character})}

		tests = append(tests, datum)
	}

	const numMoreRandomWhiteSpace = 30
	for i:=0; i<numMoreRandomWhiteSpace; i++ {

		// Randomly choose a length for the string.
		strLen := 1 + randomness.Intn(50)

		// Create a buffer to construct the string.
		buffer := make([]rune, strLen)

		// Construct the string in the buffer.
		for ii := range buffer {
			buffer[ii] = whitespaceCharacters[randomness.Intn(len(whitespaceCharacters))]
		}

		// Add it to the tests.
		datum := struct{String string}{String: string(buffer)}

		tests = append(tests, datum)
	}

	// Create some tests with all sorts of characters,
	// that won't result in something that looks like money.
	variousCharacters := []rune("!\"#%&'()*,/:;<=>?@{|}~")
	variousCharacters = append(variousCharacters, []rune("\U0001f600\U0001f601\U0001f602\U0001f603\U0001f604\U0001f605\U0001f606\U0001f607\U0001f608\U0001f609\U0001f60a\U0001f60b\U0001f60c\U0001f60d\U0001f60e\U0001f60f\U0001f610\U0001f611\U0001f612\U0001f613\U0001f614\U0001f615\U0001f616\U0001f617\U0001f618\U0001f619\U0001f61a\U0001f61b\U0001f61c\U0001f61d\U0001f61e\U0001f61f\U0001f620\U0001f621\U0001f622\U0001f623\U0001f624\U0001f625\U0001f626\U0001f627\U0001f628\U0001f629\U0001f62a\U0001f62b\U0001f62c\U0001f62d\U0001f62e\U0001f62f\U0001f630\U0001f631\U0001f632\U0001f633\U0001f634\U0001f635\U0001f636\U0001f637")...)

	for _,character := range variousCharacters {
		datum := struct{String string}{String: string(character)}

		tests = append(tests, datum)
	}

	for _,character := range variousCharacters {
		datum := struct{String string}{String: string([]rune{character, character})}

		tests = append(tests, datum)
	}

	for _,character := range variousCharacters {
		datum := struct{String string}{String: string([]rune{character, character, character})}

		tests = append(tests, datum)
	}

	const numMoreRandomVarious = 30
	for i:=0; i<numMoreRandomVarious; i++ {

		// Randomly choose a length for the string.
		strLen := 1 + randomness.Intn(50)

		// Create a buffer to construct the string.
		buffer := make([]rune, strLen)

		// Construct the string in the buffer.
		for ii := range buffer {
			buffer[ii] = variousCharacters[randomness.Intn(len(variousCharacters))]
		}

		// Add it to the tests.
		datum := struct{String string}{String: string(buffer)}

		tests = append(tests, datum)
	}

	// Return
	return tests
}


func TestParseCADErr(t *testing.T) {

	tests := generateTests()

	for testNumber, test := range tests {

		actual, err := ParseCAD(test.String)
		if nil == err {
			t.Errorf("For test #%d, did not receive an error but expected one: err = %v and actual = %v", testNumber, err, actual)
		}
	}
}


func TestParseUSDErr(t *testing.T) {

	tests := generateTests()

	for testNumber, test := range tests {

		actual, err := ParseUSD(test.String)
		if nil == err {
			t.Errorf("For test #%d, did not receive an error but expected one: err = %v and actual = %v", testNumber, err, actual)
		}
	}
}

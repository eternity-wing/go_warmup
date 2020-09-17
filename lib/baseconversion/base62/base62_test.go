package base62

import (
	"reflect"
	"strings"
	"testing"
)

type symbolTestCase struct {
	character rune
	value     int
	err       error
}

type stringTestCase struct {
	str   string
	value int
	err   error
}

var validSymbols = []symbolTestCase{
	{character: '0', value: 0, err: nil},
	{character: '5', value: 5, err: nil},
	{character: 'A', value: 10, err: nil},
	{character: 'Y', value: 34, err: nil},
	{character: 'a', value: 36, err: nil},
	{character: 'z', value: 61, err: nil},
}

var validStringCases = []stringTestCase{
	{str: "100", value: 3844, err: nil},
	{str: "ABC", value: 39134, err: nil},
	{str: "00Z", value: 35, err: nil},
}

func TestBase62Convertor_Decode(t *testing.T) {
	convertor := Base62Convertor{}

	stringCases := []stringTestCase{
		{str: "00-", value: -1, err: &UnexpectedCharacterError{}},
	}

	copy(stringCases, validStringCases)

	for _, stringTestCase := range stringCases {
		val, err := convertor.Decode(stringTestCase.str)

		if stringTestCase.err == nil && val != stringTestCase.value {
			t.Errorf("Expect return %d for %s, but return %d", stringTestCase.value, stringTestCase.str, val)
		} else if reflect.TypeOf(err) != reflect.TypeOf(stringTestCase.err) {
			t.Errorf("Expected error of type %T but %T", stringTestCase.err, err)
		}

	}

}

func TestBase62Convertor_Encode(t *testing.T) {
	convertor := Base62Convertor{}

	for _, stringTestCase := range validStringCases {
		val, err := convertor.Encode(stringTestCase.value)
		expectedStr := strings.TrimLeft(stringTestCase.str, "0")
		if stringTestCase.err == nil && val != expectedStr {
			t.Errorf("Expect return %s for %d, but return %s", expectedStr, stringTestCase.value, val)
		} else if reflect.TypeOf(err) != reflect.TypeOf(stringTestCase.err) {
			t.Errorf("Expected error of type %T but %T", stringTestCase.err, err)
		}

	}

}

func TestBase62Convertor_getSymbolValue(t *testing.T) {
	convertor := Base62Convertor{}
	symbols := []symbolTestCase{
		{character: '*', value: -1, err: &UnexpectedCharacterError{}},
	}
	copy(symbols, validSymbols)
	for _, symbolTestCase := range symbols {
		val, err := convertor.getSymbolValue(symbolTestCase.character)

		if symbolTestCase.err != nil && val != symbolTestCase.value {
			t.Errorf("Expect return %d for %c, but return %d", symbolTestCase.value, symbolTestCase.character, val)
		} else if reflect.TypeOf(err) != reflect.TypeOf(symbolTestCase.err) {
			t.Errorf("Expected error of type %T but %T", symbolTestCase.err, err)
		}
	}

}

func TestBase62Convertor_getSymbolOfNumber(t *testing.T) {
	convertor := Base62Convertor{}
	symbols := []symbolTestCase{
		{character: '*', value: -1, err: &UnexpectedNumberError{}},
	}
	for _, symbolTestCase := range symbols {
		char, err := convertor.getSymbolOfNumber(symbolTestCase.value)

		if symbolTestCase.err == nil && char != symbolTestCase.character {
			t.Errorf("Expect return %c for %d, but return %c", symbolTestCase.character, symbolTestCase.value, char)
		} else if reflect.TypeOf(err) != reflect.TypeOf(symbolTestCase.err) {
			t.Errorf("Expected error of type %T but %T", symbolTestCase.err, err)
		}
	}

}

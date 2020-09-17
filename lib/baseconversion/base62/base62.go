package base62

import (
	"math"
	"strings"
)

type Base62Convertor struct {
}

func (r Base62Convertor) GetBaseNumber() int {
	return 62
}

func (r Base62Convertor) Encode(number int) (string, error) {
	encodedString := ""
	for true {
		remainder := number % r.GetBaseNumber()
		number = number / r.GetBaseNumber()
		symbol, err := r.getSymbolOfNumber(remainder)
		if err != nil {
			return "", err
		}
		encodedString = string(symbol) + encodedString
		if number == 0 {
			break
		}
	}
	return encodedString, nil
}
func (r Base62Convertor) Decode(str string) (int, error) {
	decodedValue := 0
	str = strings.TrimLeft(str, "0")
	strLen := len(str)
	for i, character := range str {
		symbolValue, err := r.getSymbolValue(character)
		if err != nil {
			return -1, err
		}
		positionPower := strLen - (i + 1)
		decodedValue += symbolValue * int(math.Pow(float64(r.GetBaseNumber()), float64(positionPower)))
	}
	return decodedValue, nil
}

const (
	ZeroAsciCode         rune = '0'
	NumberOfAlphabets    int  = 26
	CapitalLettersOffset      = int('A'-'9') - 1
	SmallLettersOffset        = int('a'-'9') - NumberOfAlphabets - 1
)

func (r Base62Convertor) getSymbolValue(character rune) (int, error) {
	var offsetFromZero = int(character - ZeroAsciCode)
	switch {
	case character >= '0' && character <= '9':
		return offsetFromZero, nil
	case character >= 'A' && character <= 'Z':
		return offsetFromZero - CapitalLettersOffset, nil
	case character >= 'a' && character <= 'z':
		return offsetFromZero - SmallLettersOffset, nil
	default:
		return -1, &UnexpectedCharacterError{}
	}
}

func (r Base62Convertor) getSymbolOfNumber(number int) (rune, error) {
	var offsetFromZero = rune(number) + ZeroAsciCode
	switch {
	case number >= 0 && number <= 9:
		return offsetFromZero, nil
	case number > 9 && number <= 9+NumberOfAlphabets:
		return offsetFromZero + rune(CapitalLettersOffset), nil
	case number > NumberOfAlphabets && number <= (9+2*NumberOfAlphabets):
		return offsetFromZero + rune(SmallLettersOffset), nil
	default:
		return -1, &UnexpectedNumberError{}
	}
}

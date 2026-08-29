package romannumbers

import (
	"strings"
)

type Roman struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []Roman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return arabic
}

func HasMoreThan3ConsecutivesSymbols(roman string) bool {
	var (
		counter int
		letter  string
	)
	for _, c := range roman {
		if string(c) != letter {
			letter = string(c)
			counter = 1
		} else {
			counter++
		}

		if counter == 4 {
			return true
		}
	}

	return false
}

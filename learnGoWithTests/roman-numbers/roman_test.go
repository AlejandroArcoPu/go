package romannumbers

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {

	t.Run("converting in both sides should return the original number", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				return true
			}
			t.Log("testing", arabic)
			roman := ConvertToRoman(arabic)
			fromRoman := ConvertToArabic(roman)
			return arabic == fromRoman
		}

		if err := quick.Check(assertion, nil); err != nil {
			t.Error("failed checks", err)
		}
	})

	t.Run("can't have more than 3 consecutives symbols", func(t *testing.T) {

		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				return true
			}
			roman := ConvertToRoman(arabic)
			t.Log("testing", arabic, roman)
			return HasMoreThan3ConsecutivesSymbols(roman) == false
		}

		if err := quick.Check(assertion, nil); err != nil {
			t.Error("failed checks", err)
		}
	})

}

func TestHasMoreThan3ConsecutivesSymbols(t *testing.T) {
	cases := []struct {
		Roman string
		Want  bool
	}{
		{"III", false},
		{"XCIX", false},
		{"IIII", true},  // not valid roman
		{"VVVVX", true}, // not valid roman
		{"IX", false},
		{"MCDLXXVIII", false},
		{"MCCCXCII", false},
	}

	for _, test := range cases {
		description := test.Roman + "check if more than 3 consecutives"
		t.Run(description, func(t *testing.T) {
			got := HasMoreThan3ConsecutivesSymbols(test.Roman)
			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}

// func TestHasWrongSubstractors(t *testing.T) {
// 	cases := []struct {
// 		Roman string
// 		Want  bool
// 	}{
// 		{"X", false},
// 		{"VX", true},
// 		{"IX", false},
// 		{"CXI", false},
// 		{"XDI", true}, // not valid roman
// 	}

// 	for _, test := range cases {
// 		description := test.Roman + "check if wrong subsctractors"
// 		t.Run(description, func(t *testing.T) {
// 			got := HasWrongSubstractors(test.Roman)
// 			if got != test.Want {
// 				t.Errorf("got %v, want %v", got, test.Want)
// 			}
// 		})
// 	}
// }

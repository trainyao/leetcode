package main

import (
	"log"
)

func main() {
	log.Printf("%d", romanToInt("III"))
	log.Printf("%d", romanToInt("IV"))
	log.Printf("%d", romanToInt("IX"))
	log.Printf("%d", romanToInt("LVIII"))
	log.Printf("%d", romanToInt("CXXIII"))
	log.Printf("%d", romanToInt("MCMXCIV"))
}

const ONE = "I"
const FOUR = "IV"
const FIVE = "V"
const NINE = "IX"
const TEN = "X"
const FORTY = "XL"
const FIFTY = "L"
const NINETY = "XC"
const ONE_HUNDRED = "C"
const FOUR_HUNDRED = "CD"
const FIVE_HUNDRED = "D"
const NINE_HUNDRED = "CM"
const THOUSAND = "M"

type Pair struct {
	Key    int
	Value  string
	Length int
}

func romanToInt(s string) int {
	result := 0
	charMap := []Pair{
		{900, NINE_HUNDRED, 2},
		{400, FOUR_HUNDRED, 2},
		{90, NINETY, 2},
		{40, FORTY, 2},
		{9, NINE, 2},
		{4, FOUR, 2},

		{1000, THOUSAND, 1},
		{500, FIVE_HUNDRED, 1},
		{100, ONE_HUNDRED, 1},
		{50, FIFTY, 1},
		{10, TEN, 1},
		{5, FIVE, 1},
		{1, ONE, 1},
	}

	for len(s) > 0 {
		for _, pair := range charMap {
			if len(s) >= pair.Length && s[0:pair.Length] == pair.Value {
				result += pair.Key
				s = s[pair.Length:]
				break;
			}
		}
	}

	return result
}

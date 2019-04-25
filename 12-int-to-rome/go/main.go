package main

import "log"

func main() {
	log.Printf(intToRoman(3))
	log.Printf(intToRoman(4))
	log.Printf(intToRoman(9))
	log.Printf(intToRoman(58))
	log.Printf(intToRoman(123))
	log.Printf(intToRoman(1994))
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
	Key int
	Value string
}

func intToRoman(num int) string {
	result := ""
	charMap := []Pair{
		{1000,THOUSAND},
		{1000, THOUSAND},
		{900, NINE_HUNDRED},
		{500, FIVE_HUNDRED},
		{400, FOUR_HUNDRED},
		{100, ONE_HUNDRED},
		{90, NINETY},
		{50, FIFTY},
		{40, FORTY},
		{10, TEN},
		{9, NINE},
		{5, FIVE},
		{4, FOUR},
		{1, ONE},
	}

	for _, pair := range charMap {
		count := num / pair.Key
		if count >= 1 {
			for i := 0; i < count; i++ {
				result += pair.Value
			}
			num = num - (count * pair.Key)
		}
	}

	return result
}

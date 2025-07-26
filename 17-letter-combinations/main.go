package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%+v", letterCombinations("234"))
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	m := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	output := append([]string{}, m[string(digits[0])]...)
	for i := 1; i < len(digits); i++ {
		choose := m[string(digits[i])]

		l := len(output)
		for j := 1; j < len(choose); j++ {
			for k := 0; k < l; k++ {
				// fmt.Printf("append %+v\n", output[k]+choose[j])
				output = append(output, output[k]+choose[j])
			}
		}
		for m := 0; m < l; m++ {
			// fmt.Printf("edit %+v\n", output[m]+choose[0])
			output[m] += choose[0]
		}

		// fmt.Printf("output %+v\n", output)
	}

	return output
}

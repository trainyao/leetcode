package main

import (
	"errors"
	"log"
)

const point = '.'
const star = '*'
const specialChar = "*."

func main() {
	s := "mississippi"
	p := "mis*is*p*."

	result, err := isMatch(s, p)
	if err != nil {
		result = false
	}
	log.Printf("result shoule be false , %b", result)
}

func isMatch(s string, p string) (bool, error) {
	// 是否已经处理完毕
	if len(s) == 0 && len(p) == 0 {
		log.Printf("matching finished, and match")
		return true, nil
	} else if len(s) == 0 {
		log.Printf("matching finished, and not match")
		return false, nil
	} else if len(p) == 0 {
		log.Printf("matching finished, and not match")
		return false, nil
	}

	if len(p) >= 2 {
		if (p[0] == star || p[0] == point) && (p[1] == star || p[1] == point) {
			log.Printf("invalid regex format")
			return false, errors.New("invalid regex format " + p)
		}
	}

	log.Printf("matching %s with %s", s, p)
	var result bool = false
	defer log.Printf("", result)

	for _, regexChar := range p {
		// 如果是点,吃一个
		if regexChar == point {
			result, err := isMatch(s[1:len(s)], p[1:len(p)])
			if err != nil {
				return false, err
			}
			return result, nil
		}

		// 如果是星 循环吃
		if regexChar == star {
			restRegex := p[1:len(p)]

			// 尝试不跳,直接吃掉一个regex的char来匹配
			result, err := isMatch(s, restRegex)
			if err != nil {
				return false, err
			}
			if result {
				result = true
				return true, nil
			}

			for stringIndex, _ := range s {
				result, err := isMatch(s[stringIndex+1:len(s)], p[1:len(p)])
				if err != nil {
					return false, err
				}
				if result {
					return true, nil
				}
			}

			result = false
			return false, nil
		}

		// 剩下的做字符的匹配
		asciiSubStringIndex := findNextAsciiSubString(p)
		if asciiSubStringIndex < 0 {
			log.Printf("impossible, first index is not ascii and pass first two if")
			result = false
			return false, nil
		}

		// 匹配ascii字符串, 继续匹配后续的字符串
		if asciiSubStringIndex > len(s) {
			return false, nil
		}
		if p[:asciiSubStringIndex] == s[:asciiSubStringIndex] {
			result, err := isMatch(s[asciiSubStringIndex:], p[asciiSubStringIndex:])
			if err != nil {
				return false, err
			}

			return result, nil
		}

		result = false
		return false, nil
	}

	result = false
	return false, nil
}

func findNextAsciiSubString(str string) int {
	if len(str) == 0 {
		return -1
	}

	for index, char := range str {
		if char == point || char == star {
			return index
		}
	}

	return len(str)
}

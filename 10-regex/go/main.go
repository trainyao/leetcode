package main

import (
	"log"
	"strings"
)

const point = '.'
const star = '*'

func main() {
	//s := "aa"
	//p := "a*"

	//s := "ab"
	//p := ".*"

	s := "aaa"
	p := "ab*a*c*a"

	//s := "a"
	//p := "ab*"

	//s := "aab"
	//p := "d*c*a*b"
	//
	//s := "mississippi"
	//p := "mis*is*p*."

	result := isMatch(s, p)
	log.Printf("result shoule be false , %b", result)
}

func isMatch(s string, p string) bool {
	// 是否已经处理完毕
	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(s) == 0 {
		// 后尾的字符是*,而且已经将字符串匹配完了
		starIndex := strings.Index(p, string(star))
		if starIndex == len(p)-1 {
			return true
		}
		return false
	} else if len(p) == 0 {
		return false
	}

	log.Printf("matching %s with %s", s, p)

	for index, regexChar := range p {
		// 如果是点,吃一个
		if regexChar == point {
			// 如果存在下一个字符,并且下一个字符是*,直接吃掉一个匹配的字符
			if index+1 <= len(p)-1 && p[index+1] == star {
				return isMatch(s[1:], p)
			}
			return isMatch(s[1:], p[1:])
		}

		if regexChar == star {
			return false
		}

		// 查找星
		starIndex := strings.IndexAny(p, string(star))
		if starIndex >= 0 {
			// 星前面的字符长度大于可以匹配的长度
			if len(s) < starIndex {
				return false
			}

			if s[:starIndex] == p[:starIndex] {
				return isMatch(s[starIndex:], p)
			} else {
				// 星前面的文字不匹配,继续匹配星之后的正则表达式
				return isMatch(s, p[starIndex+1:])
			}
		}

		return s == p
	}

	return false
}

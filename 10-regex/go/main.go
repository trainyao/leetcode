package main

import (
	"log"
)

func main() {
	//s := "aa"
	//p := "a*"

	//s := "ab"
	//p := ".*"

	//s := "aa"
	//p := "ab*a*"

	s := ""
	p := ".*"

	//s := "a"
	//p := "ab*"

	//s := "aab"
	//p := "c*a*b"
	//
	//s := "mississippi"
	//p := "mis*is*p*."

	result := isMatch(s, p)
	log.Printf("result shoule be false , result is %b", result)
}

func isMatch(s string, p string) bool {
	var dp [][]bool

	for i := 0; i <= len(s); i++ {
		var tmp []bool
		for j := 0; j <= len(p); j++ {
			tmp = append(tmp, false)
		}
		dp = append(dp, tmp)
	}

	dp[0][0] = true

	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if j > 1 && p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j])
			} else {
				dp[i][j] = i > 0 && dp[i-1][j-1] && ((s[i-1] == p[j-1]) || (p[j-1] == '.'))
			}
		}
	}

	return dp[len(s)][len(p)]
}

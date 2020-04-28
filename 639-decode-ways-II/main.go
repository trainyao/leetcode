package main

func main() {
	//r := numDecodings("**")
	//print(r)
	//
	//r = numDecodings("1*")
	//print(r)
	//
	//r = numDecodings("2*")
	//print(r)
	//
	//r = numDecodings("9*")
	//print(r)

	r := numDecodings("*********")
	/*
	"*********"
	Output:
	1291868919
	Expected:
	291868912

	*/
	print(r)
}

const M = 1000000007

func numDecodings(s string) int {
	l := make([]int, len(s)+1)
	l[0] = 1
	otn := map[int]struct{}{
		'1': {},
		'2': {},
		'3': {},
		'4': {},
		'5': {},
		'6': {},
		'7': {},
		'8': {},
		'9': {},
	}

	ots := map[int]struct{}{
		'1': {},
		'2': {},
		'3': {},
		'4': {},
		'5': {},
		'6': {},
	}

	if s[0] == '0' {
		l[1] = 0
	}
	if s[0] == '*' {
		l[1] = 9
	}
	if _, ok := otn[int(s[0])]; ok {
		l[1] = 1
	}

	counter := 1
	n := counter + 1
	for counter < len(s) {
		c := s[counter]
		cb := s[counter-1]
		if c == '0' {
			l[n] = 0

			if cb == '1' || cb == '2' {
				l[n] += l[n-2]
			}
			if cb == '*' {
				l[n] += (2 * l[n-2]) % M
			}
		}

		if _, isOneToNine := otn[int(c)]; isOneToNine {
			l[n] = l[n-1]

			// n[n-2] (s[n-1] = 1 or s[n] =[1:6] and s[n-1] = 2)
			// n[n-2] (s[n-1] = * and s[n] = [7-9])
			_, isOneToSix := ots[int(c)]
			if cb == '1' || (isOneToSix && cb == '2') || (!isOneToSix && cb == '*') {
				l[n] += l[n-2]
			}

			// 2n[n-2] (s[n-1] = * and s[n] = [1-6])
			if isOneToSix && cb == '*' {
				l[n] += (2 * l[n-2]) % M
			}
		}

		// 9n[n-1] +
		if c == '*' {
			l[n] = (9 * l[n-1]) % M

			// 9n[n-2](s[n-1] = 1) 6n[n-2](s[n-1] = 2) 15n[n-2](s[n-1] = *)
			if cb == '1' {
				l[n] += (9 * l[n-2]) % M
			}
			if cb == '2' {
				l[n] += (6 * l[n-2]) % M
			}
			if cb == '*' {
				l[n] += (15 * l[n-2]) % M
			}
		}

		l[n] = l[n] % M

		counter++
		n++
	}
	return l[len(l)-1]
}

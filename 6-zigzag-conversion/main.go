package main

import "fmt"

func main() {
	s := `PAYPALISHIRING`
	n := 3
	fmt.Println(convert(s, n))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	// 将输入字符串分为n组，每组n+n-2个字母
	// 分组后，输出的顺序恰好是一首一尾：
	// A   G
	// B  FH
	// C E I
	// D   J
	// 分组后是这样的⬇️
	// G1:      A B C D E F
	// 输出顺序: 1 3       4
	// G2:      G H I J x x
	// 输出顺序: 2 5       6(如果这个位置有字符的话)
	// 所以可以使用首位指针法，循环每组n次，每组使用首位指针法输出字母

	output := ""

	n := len(s)

	// 每组字符的数目，等于行数+组间字符，比如
	// A   G
	// B  FH
	// C E I
	// D   J
	// 就是A～D+EF，等于4+（4-2） = 6
	groupMembers := numRows + (numRows - 2)
	nGroups := n / groupMembers

	// i 计数器，i是行数，用来计算输出到第几行，i循环到numRows的一半，算法就结束了
	i := 0
	for {
		// j计数器，循环到哪一组了
		j := 0

		for {
			groupStart := 0 + (j * groupMembers)
			groupEnd := groupStart + groupMembers
			// x y 是首尾指针
			x := groupStart + i
			y := groupEnd - i

			// fmt.Printf("x %d y %d\n", x, y)

			// print s[x], when:
			// 1. x < n（越界判断）
			if x < n {
				output += string(s[x])
				// fmt.Println(string(s[x]))
			}

			// print s[y], when:
			// 1. x != y(同一个字符不用输出2次)
			// 2. y != groupEnd（尾指针算到每组+1，用来保持平衡）
			// 3. y < n（越界判断）
			if x != y && y != groupEnd && y < n {
				output += string(s[y])
				// fmt.Println(string(s[y]))
			}

			j++
			if j > nGroups {
				break
			}
		}
		i++   //行数++
		j = 0 // 重置组计数器，从第一组开始再循环

		// if x > groupMembers/2, break all, return output
		if i > groupMembers/2 {
			return output
		}
	}
}

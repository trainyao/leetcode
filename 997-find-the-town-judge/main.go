package main

import (
	"bytes"
)

func main() {
}

// findJudge 用trusted 和 trusting 数组记录被信任和信任别人的端点数
// 如 N=4 1 相信 3
// trusted[3-1] = []byte{[0000 0001]}
// trusting[1-1] = []byte{[0000 0100]}
// 这里做复杂了, 原题说trust 对不会重复, 可以直接用一个uint16计数, 这里做成了trust 对可重复
func findJudge(N int, trust [][]int) int {
	if N <= 1 {
		return 1
	}

	trusted := make([][]byte, N)
	trusting := make([][]byte, N)
	c := N/8+1
	// allTrusted 是 N 位为1的数组
	// 先 push N/8 个 8位都为1的byte, 最后一个高位byte 计算剩下有多少位是1, 赋值
	// 如 N = 50
	// N/8 = 6, N%8 = 2
	// 于是 len(allTrusted) = 7, 前6个byte是 8位1, 第7个是低2位是1的byte
	// 分布:  allTrusted = []byte{ [00000011](第七个byte), [1*8],[].....[],[1*8](第1-6个byte) }
	allTrusted := append(bytes.Repeat([]byte{255}, c-1), []byte{0 | (1<<(N%8) - 1)}...)

	buffer := make([]byte, 2*N*c)
	bufferCounter := 0

	check := []int{}

	for _, p := range trust {
		beingTrustedIndex := p[1] - 1
		trustingIndex := p[0] - 1
		if trusted[beingTrustedIndex] == nil {
			trusted[beingTrustedIndex] = buffer[bufferCounter : bufferCounter+c]
			bufferCounter += c
		}

		// find certain bit, set to 1
		trusted[beingTrustedIndex][trustingIndex/8] |= 1 << (trustingIndex % 8)

		if trusting[trustingIndex] == nil {
			trusting[trustingIndex] = buffer[bufferCounter : bufferCounter+c]
			bufferCounter += c
		}
		// find certain bit, set to 1
		trusting[trustingIndex][beingTrustedIndex/8] |= 1 << (beingTrustedIndex % 8)

		allTrusted[beingTrustedIndex/8] ^= 1 << (beingTrustedIndex % 8)
		if bytes.Equal(allTrusted, trusted[beingTrustedIndex]) {
			// TODO if p[1] trusted anyone, don't append it prevent checking below
			check = append(check, p[1])
		}
		// edit back, reuse allTrusted byte array
		allTrusted[beingTrustedIndex/8] |= 1 << (beingTrustedIndex % 8)
	}
	if len(check) == 0 {
		return -1
	}

	// check node in 'check' array, if node trusted no one, return it as judge
	for _, c := range check {
		// check if trusting anyone, if nil, `c` trusted no one
		if trusting[c-1] == nil {
			return c
		}
	}

	return -1
}

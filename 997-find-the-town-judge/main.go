package main

import (
	"bytes"
)

func main() {
}

func findJudge(N int, trust [][]int) int {
	if N <= 1 {
		return 1
	}

	trusted := make([][]byte, N)
	trusting := make([][]byte, N)
	c := N/8+1
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
		// 0*8 ..... 0*8   n % 8  trusted[beingT][beingT/8] |= 1<<(beingT%8 - 1)
		trusted[beingTrustedIndex][trustingIndex/8] |= 1 << (trustingIndex % 8)

		if trusting[trustingIndex] == nil {
			trusting[trustingIndex] = buffer[bufferCounter : bufferCounter+c]
			bufferCounter += c
		}
		// find certain bit, set to 1
		// 0*8 ..... 0*8   n % 8  trusted[beingT][beingT/8] |= 1<<(beingT%8 - 1)
		trusting[trustingIndex][beingTrustedIndex/8] |= 1 << (beingTrustedIndex % 8)

		//trusted[p[1]-1] |= 1 << (p[0] - 1)
		//trusting[p[0]-1] |= 1 << (p[1] - 1)
		// all trusted : current == alltrusted ^ 1<<(p-1)
		// no trusted : 1<<(p[1]-1) ^ any == 0
		allTrusted[beingTrustedIndex/8] ^= 1 << (beingTrustedIndex % 8)
		if bytes.Equal(allTrusted, trusted[beingTrustedIndex]) {
			check = append(check, p[1])
		}
		// edit back
		allTrusted[beingTrustedIndex/8] |= 1 << (beingTrustedIndex % 8)

		//trusted[beingTrustedIndex][] |= 1 << ((beingTrustedIndex % 8) - 1)
		//if trusted[p[1]-1] == allTrusted^1<<(p[1]-1) { //  && trusted[p[1]-1] == 0 {
		//	check = append(check, p[1])
		//}
	}
	if len(check) == 0 {
		return -1
	}

	//trustingNoOne := bytes.Repeat([]byte{0}, N)
	for _, c := range check {
		// check if trusting anyone
		if trusting[c-1] == nil {
			return c
		}
	}

	return -1
}

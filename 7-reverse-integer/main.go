package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", reverse(200))
	fmt.Printf("%d\n", reverse(0))
	fmt.Printf("%d\n", reverse(-200))
	fmt.Printf("%d\n", reverse(123))
	fmt.Printf("%d\n", reverse(-123))
	fmt.Printf("%d\n", reverse(-9999999999))
	fmt.Printf("%d\n", reverse(9999999999))
}

func reverse(x int) int {
	pmax := 214748364
	mmax := -214748364

	output := 0
	for {
		if x == 0 {
			return output
		}

		t := x % 10
		if output > pmax || (output == pmax && t > 7) {
			return 0
		}
		if output < mmax || (output == mmax && t < -8) {
			return 0
		}
		x = x / 10
		output = output*10 + t
	}
}

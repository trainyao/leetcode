package main

import (
	"log"
)

func main() {
	log.Print(scoreOfParentheses("(()(()))"))
}


func scoreOfParentheses(S string) int {
	stack := []int{}
	for index, token := range S {
		//log.Printf("token %s index %d", string(token), index)
		if string(token) == "(" {
			stack = append(stack, 1)
			//log.Printf("pushing %v", stack)
		}
		if string(token) == ")" {
			//log.Printf("%v", stack)
			//log.Printf("%v %v %v", index, string(token), S)
			stack = stack[:len(stack)-1]
			//log.Printf("poping %v", stack)
		}

		if len(stack) == 0 {
			if index == len(S)-1 {
				if len(S) == 2 {
					return 1
				}

				// end x ing
				//log.Printf("end x ing %s", S[1:len(S)-1])

				return 2 * scoreOfParentheses(S[1:len(S)-1])
			}
			// not end,+ing

			//log.Printf("not end, + ing %s %s", S[:index+1],S[index+1:])
			return scoreOfParentheses(S[:index+1]) + scoreOfParentheses(S[index+1:])
		}
	}

	return 0
}

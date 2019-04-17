package main

import "log"

func main() {
	slice := [0]int{}
	log.Printf("%a", slice[:0])
}

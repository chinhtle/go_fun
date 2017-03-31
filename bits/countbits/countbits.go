package main

import "fmt"

func countBits(num int) (n int) {
	for ; num > 0; num &= num - 1 {
		n++
	}

	return n
}

func main() {
	fmt.Println(countBits(5))
}

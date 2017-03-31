package main

import "fmt"

func swapBits(num int, i, j uint) (int) {
	if (num & (1 << i)) != (num & (1 << j)) {
		num ^= 1 << i | 1 << j
	}

	return num
}

func main() {
	fmt.Println(swapBits(0, 0, 1))
	fmt.Println(swapBits(4, 0, 2))
	fmt.Println(swapBits(4, 1, 2))
	fmt.Println(swapBits(4, 2, 2))
}

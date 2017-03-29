package main

import (
	"fmt"
	"os"
	"strconv"
)

func hanoi(disks, from, helper, to int) {
	if disks == 1 {
		fmt.Println(from, "->", to)
		return
	}

	// Move n-1 disks from source to helper
	hanoi(disks-1, from, to, helper)

	// Move 1 disk from source to target
	fmt.Println(from, "->", to)

	// Move the original moved disks from helper to target
	hanoi(disks-1, helper, from, to)
}

func main() {
	numDisks := 3

	if len(os.Args) > 1 {
		if val, err := strconv.Atoi(os.Args[1]); err == nil {
			numDisks = val
		}
	}

	from, helper, to := 1, 2, 3
	hanoi(numDisks, from, helper, to)
}

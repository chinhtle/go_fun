package main

import "fmt"

func binarySearch(arr []int, num int) int {
	var start int
	end := len(arr) - 1

	for start <= end {
		mid := start + ((end - start) / 2)

		if num > arr[mid] {
			start = mid + 1
		} else if num < arr[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}

	// Return -1 if not found
	return -1
}

// RecursiveBinarySearch performs binary search, recursively, to search for a
// number. Returns the index if found; otherwise, returns -1
func RecursiveBinarySearch(arr []int, num int) int {
	return recursiveBinarySearch(arr, num, 0, len(arr)-1)
}

func recursiveBinarySearch(arr []int, num, start, end int) int {
	if start > end {
		return -1
	}

	mid := start + ((end - start) / 2)

	if num > arr[mid] {
		return recursiveBinarySearch(arr, num, mid+1, end)
	} else if num < arr[mid] {
		return recursiveBinarySearch(arr, num, start, mid-1)
	}

	return mid
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 61, 70, 100, 2000}
	numSearch := arr[5]

	fmt.Println(numSearch, "index:", binarySearch(arr, numSearch))
	fmt.Println(numSearch, "index:", RecursiveBinarySearch(arr, numSearch))
}

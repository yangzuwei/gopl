// binary_search project binary_search.go
package main

import "fmt"

func findNumber(arr []int, number int) int {
	low := 0
	high := len(arr)

	for high >= low {
		mid := low + (high-low)/2
		if arr[mid] == number {
			return mid
		}

		if arr[mid] > number {
			high = mid - 1
		}

		if arr[mid] < number {
			low = mid + 1
		}
	}
	return low
}

func findByRecursion(arr []int, number int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2

	if arr[mid] == number {
		return mid
	}

	if arr[mid] > number {
		return findByRecursion(arr, number, low, mid-1)
	} else {
		return findByRecursion(arr, number, mid+1, high)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("position is ", findNumber(arr, 3))
	fmt.Println("position is ", findByRecursion(arr, 5, 0, len(arr)))
}

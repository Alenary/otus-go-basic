package main

import (
	"fmt"
)

// binarySearch возвращает индекс найденного элемента или -1, если элемент не найден.

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // элемент не найден.
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	result := binarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Элемент найден на индексе %d\n", result)
	} else {
		fmt.Println("Элемент не найден")
	}
}

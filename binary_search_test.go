package algorithm_practice

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}
	fmt.Println("ele position: ", BinarySearch(arr, 19))
}

// BinarySearch 二分查找, 要求数据有序
func BinarySearch(arr []int, guess int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Printf("low:%d high:%d mid:%d\n", low, high, mid)
		if guess == arr[mid] {
			return mid
		}
		if guess < arr[mid] {
			high = mid - 1
		}
		if guess > arr[mid] {
			low = mid + 1
		}
	}

	return -1
}

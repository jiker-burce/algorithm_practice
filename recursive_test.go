package algorithm_practice

import (
	"fmt"
	"testing"
)

func TestRecursive(t *testing.T) {
	arr := []int{4, 2, 19, 7, 3, 14}
	var max int = -1
	MaxNum(arr, 0, &max)
	fmt.Println("max:", max)
}

func MaxNum(arr []int, i int, max *int) {
	if arr[i] > *max {
		*max = arr[i]
	}
	if i >= len(arr)-1 {
		return
	}
	MaxNum(arr, i+1, max)
}

// 求和: 缩小问题规模
func r2() {
	var r func(arr []int) int
	arr := []int{2, 4, 6}

	r = func(arr []int) int {
		if len(arr) <= 1 {
			return arr[0]
		}
		lenArr := len(arr)
		return r(arr[:lenArr-1]) + arr[lenArr-1]
	}
	fmt.Println("result: ", r(arr))
}

// 求和: 常规递归
func r1() {
	var r func(arr []int, i int) int
	arr := []int{2, 4, 6}

	r = func(arr []int, i int) int {
		if i == len(arr)-1 {
			return arr[i]
		}
		return r(arr, i+1) + arr[i]
	}
	fmt.Println("result: ", r(arr, 0))
}

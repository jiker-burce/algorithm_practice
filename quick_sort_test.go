package algorithm_practice

import (
	"fmt"
	"testing"
)

// 快速排序
// 算法图解第4章
func TestQuickSort(t *testing.T) {
	arr := []int{3, 9, 1, 4, 7, 61, 31}
	fmt.Println("sorted arr", QuickSort(arr))
}

func QuickSort(arr []int) []int {
	pivot := arr[0]
	less := make([]int, 0)
	greater := make([]int, 0)

	if len(arr) <= 1 {
		return arr
	}

	// 每层O(n)
	for _, v := range arr {
		if v < pivot {
			less = append(less, v)
		}
		if v > pivot {
			greater = append(greater, v)
		}
	}

	tmp := make([]int, 0)
	if len(less) > 0 {
		tmp = append(QuickSort(less), pivot) // 总共迭代O(log n)层
	}
	if len(greater) > 0 {
		tmp = append(tmp, QuickSort(greater)...)
	}
	return tmp
}

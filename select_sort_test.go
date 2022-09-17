package algorithm_practice

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{5, 1, 3, 9, 7}
	SelectSort(arr)

	fmt.Println("newArr:", arr)
}

func SelectSort(arr []int) {
	var newArr []int
	lenArr := len(arr)
	tmpArr := make([]int, len(arr))
	copy(tmpArr, arr)

	for i := 0; i < lenArr; i++ {
		min := 0
		for j := 0; j < len(tmpArr); j++ {
			if tmpArr[j] < tmpArr[min] {
				min = j
			}
		}
		newArr = append(newArr, tmpArr[min])
		tmpArr = append(tmpArr[:min], tmpArr[min+1:]...)
	}

	copy(arr, newArr)
}

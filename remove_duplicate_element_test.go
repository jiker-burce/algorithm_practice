package algorithm_practice

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicateElement(t *testing.T) {
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}

// 双指针: left指针指向和下一个元素不一样的位置;right指针如果和left不一样则赋值给left+1位置;否则一直往前跑
// 时间复杂度为:O(n);
func removeDuplicates(nums []int) int {
	left := 0
	right := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
		right++
	}

	return left + 1
}

// 暴力解法: 每次都会整体移动后面所有元素, 效率很低
//func removeDuplicates(nums []int) int {
//	cur := 0
//	lNum := len(nums) - 1
//	for i := 0; i < lNum; i++ {
//		if nums[cur] == nums[cur+1] {
//			nums = append(nums[0:cur+1], nums[cur+1+1:]...)
//		} else {
//			cur += 1
//		}
//	}
//
//	return len(nums)
//}

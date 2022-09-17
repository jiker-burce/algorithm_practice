package algorithm_practice

import (
	"fmt"
	"testing"
)

// 存在重复元素 https://leetcode.cn/leetbook/read/top-interview-questions-easy/x248f5/
// 关键是用set集合(map)判断元素是否存在;和两数之和差不多的解法
func containsDuplicate(nums []int) bool {
	var existM = make(map[int]bool)
	for _, v := range nums {
		if _, ok := existM[v]; ok {
			return true
		}
		existM[v] = true
	}
	return false
}

func TestDuplicateCheck(t *testing.T) {
	nums := []int{1, 2, 3, 12}

	fmt.Println("result:", containsDuplicate(nums))
}

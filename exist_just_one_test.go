package algorithm_practice

import (
	"fmt"
	"testing"
)

// 只出现一次的数字 https://leetcode.cn/leetbook/read/top-interview-questions-easy/x21ib6/
// 异或运算: 如果出现2个相同的元素,相当于操作的值就被清除了,只有出现一次的值正好保留
func singleNumber(nums []int) int {
	result := nums[0]
	for _, v := range nums[1:] {
		result ^= v
	}
	return result
}

func TestSingleNumber(t *testing.T) {
	nums := []int{1, 2, 2, 12, 12, 1, 9}

	fmt.Println("result:", singleNumber(nums))
}

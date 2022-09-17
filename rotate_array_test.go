package algorithm_practice

import "testing"

/*
*
k %= len(nums)
[1,2,3]  =>

	1 [3,1,2]
	2 [2,3,1]
	3 [1,2,3]
	4 [3,1,2]
*/
func rotate1(nums []int, k int) {
	newNum := make([]int, 0)
	k %= len(nums)
	newNum = append(newNum, nums[len(nums)-k:]...)
	newNum = append(newNum, nums[:len(nums)-k]...)
	copy(nums, newNum)
}

func reverse(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-1-i] = nums[l-1-i], nums[i]
	}
}
func rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func rotate3(nums []int, k int) {
	newNums := make([]int, len(nums))
	l := len(nums)
	for i, v := range nums {
		newNums[(i+k)%l] = v // i+k % l 比 l 小的就都移到后面去了
	}
	copy(nums, newNums)
}

func TestRotate1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate1(nums, 3)
	t.Logf("rotate1: %v", nums)
}

func TestRotate2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(nums, 3)
	t.Logf("rotate2: %v", nums)
}

func TestRotate3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(nums, 3)
	t.Logf("rotate3: %v", nums)
}

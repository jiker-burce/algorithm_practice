package algorithm_practice

import (
	"fmt"
	"sort"
	"testing"
)

// 两个数组的交集 II https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2y0c2/

// 先对2个数组进行排序,
// 使用2个指针分别指向2个不同的切片
// 4 ms
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	newNums := make([]int, 0)

	i, j := 0, 0

	n1Len := len(nums1)
	n2Len := len(nums2)
	for i < n1Len && j < n2Len {
		switch {
		case nums1[i] > nums2[j]:
			j++
		case nums1[i] < nums2[j]:
			i++
		default:
			newNums = append(newNums, nums1[i])
			i++
			j++
		}
	}
	return newNums
}

func TestIntersect(t *testing.T) {
	//nums1 := []int{1, 2, 2, 1}
	//nums2 := []int{2, 2}
	nums1 := []int{4, 9, 5}
	nums2 := []int{8, 4, 9, 8, 4}

	fmt.Println("intersect result: ", intersectMap(nums1, nums2))
}

// map方法
// 分别将2个数组的数据存入不同的map中
//
// 0ms
func intersectMap(nums1 []int, nums2 []int) []int {
	countMap1 := make(map[int]int, len(nums1))
	countMap2 := make(map[int]int, len(nums2))

	for _, v := range nums1 {
		countMap1[v]++
	}

	for _, v := range nums2 {
		countMap2[v]++
	}

	result := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		num, ok := countMap1[v]
		if !ok {
			continue
		}

		n2 := countMap2[v]
		if num > n2 {
			num = n2
		}

		for i := 0; i < num; i++ {
			result = append(result, v)
		}
		delete(countMap1, v)
	}

	return result
}

//英语和缅甸语者
//教我中文,韩语和日语
//一只社交蝴蝶试图在这里营造一种友好的快乐氛围
//不要浪漫的接手我的友好

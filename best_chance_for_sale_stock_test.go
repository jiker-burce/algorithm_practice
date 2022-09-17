package algorithm_practice

import (
	"fmt"
	"testing"
)

// 买卖股票的最佳时机 II: https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2zsx1/
// -- 贪心算法
// -- 动态规划
// Go 1.18 泛型全面讲解：一篇讲清泛型的全部: https://segmentfault.com/a/1190000041634906
// -- Ordered 代表所有可比大小排序的类型
// 泛型会让你的 Go 代码运行变慢: https://www.infoq.cn/article/xprmcl5qbf6yvdroajyn

func TestBestChance(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println("profit: ", maxProfit(prices))
	fmt.Println("profit2: ", maxProfit2(prices))
	fmt.Println("profit3: ", maxProfit3(prices))
}

// Ordered 代表所有可比大小排序的类型
type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

// 通过泛型获取整型相关的最大值
// 如果T为comparable, 编译器会告知没有实现>运算操作
func myMaxInt[T Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}

	return v2
}

/*
*
初步看题目时,会被题目绕晕
将题目变为直角坐标系后,发现只需要计算上升区间的值之和即可
*/
func maxProfit(prices []int) int {
	totalProfit := 0

	// 边界值, 只有一个元素, 或者一个元素都没有, 就说明没有利润产生
	if len(prices) <= 1 {
		return 0
	}

	for i := 0; i < len(prices)-1; i++ {
		//if prices[i] < prices[i+1] {
		//	totalProfit += prices[i+1] - prices[i]
		//}
		totalProfit += myMaxInt(prices[i+1]-prices[i], 0)
	}

	return totalProfit
}

// 高抛低吸
func maxProfit2(prices []int) int {
	var ans, pre = 0, prices[0]

	for _, v := range prices[1:] {
		if v > pre {
			ans += v - pre
		}
		pre = v
	}

	return ans
}

// 双指针，这种算法如果涉及买卖费用则还需要优化
func maxProfit3(prices []int) int {
	left := 0
	right := 1
	//总利润
	result := 0
	for right < len(prices) {
		if prices[left] < prices[right] {
			result += prices[right] - prices[left]
		}
		left++
		right++
	}

	return result
}

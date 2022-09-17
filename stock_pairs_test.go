package algorithm_practice

import (
	"fmt"
	"testing"
)

// https://linyencheng.github.io/2022/01/07/js-hacker-rank-profit-targets-a-financial-analyst/

func stockPairs(stocksProfit []int32, target int64) int32 {
	// Write your code here
	m := map[int32]bool{}
	count := int32(0)
	for _, v := range stocksProfit {
		if _, ok := m[int32(target-int64(v))]; ok {
			count++
		} else {
			m[v] = true
		}
	}
	return count
}
func TestStackPairs(t *testing.T) {
	stocksProfit := []int32{1, 3, 9, 8, 2, 7}
	fmt.Println("out: ", stockPairs(stocksProfit, 10))
}

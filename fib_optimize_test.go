package algorithm_practice

import (
	"fmt"
	"testing"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}
func TestFibOptimize(t *testing.T) {
	fmt.Println(fact(1000))

	i := 0
	// 闭包和任何go结构体,变量,函数都一样: 注册,定义,运行
	//var fib1 func(n int) int
	//fib1 = func(n int) int {
	//	i++
	//	if n < 2 {
	//		return n
	//	}
	//
	//	return fib1(n-1) + fib1(n-2)
	//}
	//fmt.Println(i, fib1(50))

	var fib func(n int) int
	midM := make(map[int]int)
	i = 0

	fib = func(n int) int {
		i++
		fmt.Println(" ", n)
		if n < 2 {
			return n
		}

		if _, flag := midM[n-1]; flag == false {
			midM[n-1] = fib(n - 1)
		}

		if _, flag := midM[n-2]; flag == false {
			midM[n-2] = fib(n - 2)
		}
		return midM[n-1] + midM[n-2]
	}

	fmt.Println(i, fib(4))
}

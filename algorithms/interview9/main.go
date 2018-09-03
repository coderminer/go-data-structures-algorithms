package main

import "fmt"

/*
	输入n,求斐波那契数列的第n项,递归的方式
	f(n) {
		0：n=0
		1：n=1
		f(n-1)+f(n-2):n>1
	}
*/

func fibonacci(n int) int64 {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

/*
	改进，使用已经计算出来的结果，非递归的方式
*/
func fibonacci2(n int) int64 {
	result := [2]int64{0, 1}
	if n < 2 {
		return result[n]
	}

	var fibNMinusOne int64 = 1
	var fibNMinusTwo int64 = 0
	var fibN int64 = 0
	for i := 2; i <= n; i++ {
		fibN = fibNMinusOne + fibNMinusTwo
		
		fibNMinusTwo = fibNMinusOne
		fibNMinusOne = fibN
		
	}
	return fibN
}

func main() {
	// 如果n=100，递归方式时，非常耗时
	// 非递归方式会节省一些时间
	fmt.Println(fibonacci2(100))
}

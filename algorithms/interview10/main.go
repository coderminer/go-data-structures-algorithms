package main

import (
	"fmt"
)

/*
	输入一个整数，输出该数的二进制表示中1的个数，如9的二进制是 1001，输出 2
	先判断整数的二进制中最右边一位是不是1，接着把整数右移一位，这样循环
*/

func numOf1(n int) int {
	count := 0
	for n != 0 {
		if n&1 != 0 {
			count += 1
		}
		n = n >> 1
	}
	return count
}

/*
	为了避免死循环，可以不右移数字，首先吧数字和1做与运算，判断最低位是不是1，接着把1左移一位得到2
	再和数字做与运算，依次反复左移
*/
func numOf1_2(n int) int {
	count := 0
	flag := 1
	for flag != 0 {
		if n&flag != 0 {
			count += 1
		}
		flag = flag << 1
	}
	return count
}

/*
	把整数减去1，再和原整数做与运算，把该整数最右边的1变为0
*/
func numOf1_3(n int) int {
	count := 0
	for n != 0 {
		count += 1
		n = (n - 1) & n
	}
	return count
}

func main() {
	//fmt.Println(numOf1(-2))
	fmt.Println(numOf1_2(-2))
	fmt.Println(numOf1_3(-2))
}

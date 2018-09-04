package main

import (
	"fmt"
)

/*
	调整整数数组中的数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后部分
*/
func reOrderOddEven(arrs []int) {
	if arrs == nil {
		return
	}
	if len(arrs) == 0 {
		return
	}
	low := 0
	high := len(arrs) - 1
	for low < high {
		for low < high && (arrs[low]&0x01) != 0 {
			low += 1
		}
		for low < high && (arrs[high]&0x01) == 0 {
			high -= 1
		}

		if low < high {
			arrs[low], arrs[high] = arrs[high], arrs[low]
		}
	}
}

func reOrder(arrs []int, compare func(int) bool) {
	if arrs == nil || len(arrs) == 0 {
		return
	}
	low := 0
	high := len(arrs) - 1
	for low < high {
		for low < high && !compare(arrs[low]) {
			low += 1
		}
		for low < high && compare(arrs[high]) {
			high -= 1
		}
		if low < high {
			arrs[low], arrs[high] = arrs[high], arrs[low]
		}
	}
}

func isEven(n int) bool {
	return n&1 == 0
}

func reOrderOddEven1(arrs []int) {
	reOrder(arrs, isEven)
}

func main() {
	arrs := []int{1, 2, 3, 4, 5, 6, 7, 9}
	reOrderOddEven1(arrs)
	fmt.Println(arrs)
}

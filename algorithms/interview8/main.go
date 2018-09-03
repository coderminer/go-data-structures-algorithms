package main

import (
	"fmt"
)

/*
	把一个数组最开始的若干元素搬至数组的末尾，称之为数组的旋转。
	输入一个递增排序的数组的一个旋转，输出该旋转数组的最小值
*/
func min(arrs []int) int {
	if arrs == nil || len(arrs) <= 0 {
		return -1
	}

	low := 0
	high := len(arrs) - 1
	mid := low
	for arrs[low] >= arrs[high] {
		if high-low == 1 {
			mid = high
			break
		}
		mid = (low + high) / 2
		if arrs[low] == arrs[high] && arrs[mid] == arrs[low] {
			return minOrder(arrs, low, high)
		}
		if arrs[mid] >= arrs[low] {
			low = mid
		} else if arrs[mid] <= arrs[high] {
			high = mid
		}
	}
	return arrs[mid]
}

func minOrder(arrs []int, index1 int, index2 int) int {
	result := arrs[index1]
	for i := index1 + 1; i <= index2; i++ {
		if result > arrs[i] {
			result = arrs[i]
		}
	}
	return result
}

func main() {
	arr := []int{4, 5, 6, 2, 2, 3, 4}
	fmt.Println(min(arr))
}

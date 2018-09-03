package main

import "fmt"

/*
	在一个二维数组中，每一行都按照从左到右递增的顺序排序，
	每一列都按照从上到下递增的顺序排序，输入一个这样的二位数组
	和一个整数，判断数组中是否含有这整数
*/
func find(arrs [][]int, number int) bool {
	rows := len(arrs)
	columns := len(arrs[0])

	row := 0
	column := columns - 1
	for row < rows && column >= 0 {
		if arrs[row][column] == number {
			return true
		} else if arrs[row][column] > number {
			column -= 1
		} else {
			row += 1
		}
	}
	return false
}

func main() {
	arrs := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	r := find(arrs, 15)
	fmt.Println(r)
}

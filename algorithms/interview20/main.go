package main

import (
	"fmt"
)

/*
	从外到里顺时针打印二维数组
*/
func PrintMatrixClockWisely(arrs [][]int) {
	rows := len(arrs)
	columns := len(arrs[0])
	if arrs == nil || columns <= 0 || rows <= 0 {
		return
	}
	start := 0
	for columns > start*2 && rows > start*2 {
		PrintMatrixInCircle(arrs, columns, rows, start)
		start += 1
	}

}

func PrintMatrixInCircle(arrs [][]int, columns, rows, start int) {
	endX := columns - 1 - start
	endY := rows - 1 - start

	//从做到右打印一行
	for i := start; i <= endX; i++ {
		fmt.Print(arrs[start][i], " ")
	}

	//从上到下打印一列
	if start < endY {
		for i := start + 1; i <= endY; i++ {
			fmt.Print(arrs[i][endY], " ")
		}
	}

	//从右到左打印一行
	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			fmt.Print(arrs[endY][i], " ")
		}
	}

	//从下到上打印一列
	if start < endX && start < endY-1 {
		for i := endY - 1; i >= start+1; i-- {
			fmt.Print(arrs[i][start], " ")
		}
	}
}

func main() {
	arrs := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	PrintMatrixClockWisely(arrs)
}

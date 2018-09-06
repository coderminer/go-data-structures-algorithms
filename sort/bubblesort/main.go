package main

import (
	"fmt"
)

func bubbleSort(arrs []int)  {
	for i := 0; i < len(arrs); i++ {
		for j := 0; j < len(arrs); j++ {
			if arrs[i] < arrs[j] {
				arrs[i], arrs[j] = arrs[j], arrs[i]
			}
		}
	}
}

func main() {
	arrs := []int{10,3,4,8,20,32,6,-43,5,-2}
	bubbleSort(arrs)
	fmt.Println(arrs)
}

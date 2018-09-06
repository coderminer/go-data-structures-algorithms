package main

import (
	"fmt"
)

/*
	数值的整数次方
*/
func power(base float64, exponent int) float64 {
	result := 1.0
	for i := 1;i <= exponent;i++{
		result *= base
	}
	return result
}

func powerWithUnsignedExponent(base float64,exponent int) float64{
	if exponent == 0{
		return 1
	}
	if exponent == 1{
		return base
	}
	result := powerWithUnsignedExponent(base,exponent >> 1)
	result *= result
	if exponent & 0x1 == 1{
		result *= base
	}
	return result
}

func main() {
	fmt.Println(power(0,2))

	fmt.Println(powerWithUnsignedExponent(4,2))
}

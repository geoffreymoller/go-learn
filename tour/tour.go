package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128, z complex128) complex128 {
	for {
		var temp = z - ((cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2)))
		if temp == z {
			return z
		} else {
			z = temp
		}
	}
}

func main() {
	fmt.Println(Cbrt(2, 1.33))
}


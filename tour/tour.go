package main

import (
	"fmt"
	"math"
)

func Cbrt(x float64, z float64) float64 {
	for {
		var temp = z - ((math.Pow(z, 3) - x) / (3 * math.Pow(z, 2)))
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


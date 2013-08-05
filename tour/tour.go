package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64, z float64) float64 {
  for {
    var temp = z - ((math.Pow(z, 2) - x) / (2 * z))
    if temp == z {
      return z
    } else {
      z = temp
    }
  }
}

func main(){
  fmt.Println(Sqrt(2000, 1))
}


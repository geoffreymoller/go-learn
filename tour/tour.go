package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %d", e)
}

func Sqrt(x float64, z float64) (float64, error) {

  if(x < 0){
    error := ErrNegativeSqrt(x)
    return x, error
  }

  for {
    var temp = z - ((math.Pow(z, 2) - x) / (2 * z))
    if temp == z {
      return z, nil
    } else {
      z = temp
    }
  }

}

func main() {
  fmt.Println(Sqrt(4, 2))
  root, error := fmt.Println(Sqrt(-4, 2))
  if(error != nil){
    fmt.Println(error)
    fmt.Println(root)
  }
}


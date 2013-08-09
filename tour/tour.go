package main

import (
  "fmt"
)

func fibonacci() func() int {
  var count, sum, first, second int
  return func() int {
    if count == 0 {
      count++
      return 0
    } else if count == 1 {
      second = count
      count++
      return 1
    }
    sum = first + second
    first = second
    second = sum
    return sum
  }
}

func main(){
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f());
  }
}


package main


import (
  "code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
  ary := make([][]uint8, dy)
  for i := 0; i < dy; i++ {
    ary2 := make([]uint8, dx)
    for j := 0; j < dx; j++ {
      var pixel = uint8(j * i)
      ary2[j] = pixel
    }
    ary[i] = ary2
  }
  return ary
}

func main(){
  pic.Show(Pic)
}


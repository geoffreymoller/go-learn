package main

import (
  "code.google.com/p/go-tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {

  var m = make(map[string]int)
  var fields = strings.Fields(s)

  for _, v:= range fields {
    _, ok := m[v]
    if ok {
      m[v]++
    } else {
      m[v] = 1
    }
  }

  return m

}

func main(){
  wc.Test(WordCount)
}


package main

import (
  "fmt"
  "runtime"
)

func main(){
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
    case "darwin":
      fmt.Println("OS X.")
    case "linuk":
      fmt.Println("Linux.")
    default:
      //freebsd, openbsd,
      //plan9, windows...
      fmt.Printf("%s.", os)
  }
}


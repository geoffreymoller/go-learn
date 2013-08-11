package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter,  r *http.Request) {
  fmt.Fprint(w, "I'm a frayed knot.")
}

type Struct struct {
  Greeting string
  Punct string
  Who string
}

func (f *Struct) ServeHTTP(w http.ResponseWriter,  r *http.Request) {
  fmt.Fprint(w, f.Greeting + " " + f.Punct + " " + f.Who)
}

func main() {
  http.Handle("/string", String("I'm a frayed knot."))
  http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
  http.ListenAndServe("localhost:4000", nil)
}


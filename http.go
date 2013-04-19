package main

import (
    "fmt"
    "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "<div>Hello!</div>")
}

func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	_ = http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
}

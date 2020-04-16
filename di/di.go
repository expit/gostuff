package main

import (
	"fmt"
	"io"
	"net/http"
)

//Greet ...
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler ...
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "Elodie")
	err := http.ListenAndServe(":5555", http.HandlerFunc(MyGreeterHandler))
	fmt.Printf(err.Error())
}

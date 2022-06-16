package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// fmt.Printf("Hello, %s", name) - prints to default stdout
	fmt.Fprintf(writer, "Hello, %s", name) // takes a writer to send the string to
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

// Dependency injection:
// 1. Testing
// 2. Separate our concerns
// 3. Allow our code to be re-used in different contexts
// Mocking - Replacing the real thing with the pretend version

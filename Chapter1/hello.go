package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjur, "

// This is a public method, as it is starting
// with uppercase letter
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	// return greetingPrefix(language) + name
	return greetingPrefix_2(language) + name
}

// This is a private method, as it is starting
// with lowercase letter
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// This is a private method, as it is starting
// with lowercase letter
func greetingPrefix_2(language string) string {
	prefix := englishHelloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}

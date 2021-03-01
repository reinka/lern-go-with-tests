package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Holla, "

func Hello(name, language string) string {
	if name == "" {
		return "Hello, world"
	}

	prefix := getPrefix(language)
	return prefix + name
}

func getPrefix(language string) (prefix string) {
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

func main() {
	fmt.Println(Hello("world", ""))
}

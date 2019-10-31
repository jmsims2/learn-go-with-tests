package main

import "fmt"

const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "
const frenchGreeting = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return getGreeting(language) + name + "!"
}

func getGreeting(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchGreeting
	case spanish:
		prefix = spanishGreeting
	default:
		prefix = englishGreeting
	}

	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}

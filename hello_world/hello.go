package main

import "fmt"

const hungarian = "Hungarian"
const french = "French"
const englishHelloPrefix = "Hello, "
const hungarianHelloPrefix = "Szevasz, "
const frenchHelloPrefix = "Bonjour, "

func hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return languagePrefix((lang)) + name
}

func languagePrefix(lang string) (prefix string) {
	switch lang {
	case hungarian:
		prefix = hungarianHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(hello("world", "English"))
}

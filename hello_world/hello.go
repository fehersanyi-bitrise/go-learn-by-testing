package main

import "fmt"

const englishHelloPrefix = "Hello, "

func hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(hello("world"))
}

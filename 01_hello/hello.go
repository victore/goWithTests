package main

import "fmt"

const englishHelloPrefix = "Hello, "

// This function returns a string
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

// main function
func main() {
	fmt.Println(Hello("Victor"))
}

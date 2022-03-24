package main

import "fmt"

// This function returns a string
// func Hello() string {
func Hello(name string) string {
	// return "Hello, world"
	return "Hello, " + name
}

// main function
func main() {
	fmt.Println(Hello("Victor"))
}

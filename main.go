package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	whatToSay := "Hello, World!"
	fmt.Println(whatToSay)

	i := 0

	fmt.Println("i is set to", i)
	a, b := saySomething()

	fmt.Println(a, b)
}

func saySomething() (string, int) {
	return "SomeThing", 2
}

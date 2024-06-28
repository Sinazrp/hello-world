package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	whatToSay := "Hello, World!"
	fmt.Println(whatToSay)

	i := 0

	fmt.Println("i is set to", i)
	fmt.Println(saySomething())
}

func saySomething() string {
	return "SomeThing"
}

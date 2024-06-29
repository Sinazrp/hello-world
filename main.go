package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, World!")

	whatToSay := "Hello, World!"
	fmt.Println(whatToSay)

	i := 0

	fmt.Println("i is set to", i)
	a, b := saySomething()

	fmt.Println(a, b)

	myString := "green"
	log.Println("my Satring is ", myString)

	changeUsingpointer(&myString)
	log.Println("my Satring is ", myString)

}

func saySomething() (string, int) {
	return "SomeThing", 2
}

func changeUsingpointer(s *string) {
	log.Print("my Satring  set to ", s)

	*s = "changed"
}

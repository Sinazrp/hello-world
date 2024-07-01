package main

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	phoneNumber string
	age         int
	BirthDate   time.Time
}

func (m *User) printUserName() {
	log.Println("this is the call back of  func : ", m.FirstName)

}

func main() {

	user1 := User{age: 15, LastName: "zp", FirstName: "sina", BirthDate: time.Now()}
	user1.printUserName()
	log.Println(user1)
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

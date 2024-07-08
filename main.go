package main

import (
	"errors"
	"fmt"
)

//type Animal interface {
//	says() string
//	NumberOfLegs() int
//}
//type Dog struct {
//	name  string
//	breed string
//}
//type Cat struct {
//	name          string
//	breed         string
//	numberOfTeeth int
//}

//type User struct {
//	FirstName   string
//	LastName    string
//	phoneNumber string
//	age         int
//	BirthDate   time.Time
//}
//
//func (m *User) printUserName() {
//	log.Println("this is the call back of  func : ", m.FirstName)
//
//}

//func CalculateValue(intChan chan int) {
//	randomNumber := helpers.RandomInt(10)
//	intChan <- randomNumber
//}

func main() {
	//var myVar helpers.SomeType
	//myVar.TypeName = "Sina"
	//log.Println(myVar)

	//myMap := make(map[string]string)
	//myMap["wife"] = "Elnaz"
	//log.Println(myMap["wife"])
	//var mySlice []string
	//mySlice = append(mySlice, "sina")
	//mySlice = append(mySlice, "Elnaz", "ali")
	//mySlice = append(mySlice, "habib")
	//log.Println(mySlice)
	//log.Println(myMap["wife"])
	//numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//log.Println(numbers)
	//log.Println(numbers[:3])
	//names := [4]string{"one", "two", "three", "four"}
	//log.Println(names)

	//user1 := User{age: 15, LastName: "zp", FirstName: "sina", BirthDate: time.Now()}
	//user1.printUserName()
	//log.Println(user1)
	//fmt.Println("Hello, World!")
	//
	//whatToSay := "Hello, World!"
	//fmt.Println(whatToSay)
	//
	//i := 0
	//
	//fmt.Println("i is set to", i)
	//a, b := saySomething()
	//
	//fmt.Println(a, b)
	//
	//myString := "green"
	//log.Println("my Satring is ", myString)
	//
	//changeUsingpointer(&myString)
	//log.Println("my Satring is ", myString)

	//animals := []string{"dog", "cat", "cow"}
	//for _, animal := range animals {
	//	fmt.Println(animal)
	//}
	//animals := map[string]string{}
	//animals["cat"] = "mewo"
	//animals["dog"] = "bark"
	//animals["goat"] = "BeEee"
	//for k, v := range animals {
	//	fmt.Println(k, v)
	//}

	//dog := Dog{
	//	name:  "Dog",
	//	breed: "german Shephered",
	//}
	//PrintInfo(dog)

	//intChan := make(chan int)
	//defer close(intChan)
	//go CalculateValue(intChan)
	//num := <-intChan
	//fmt.Println(num)

	result, err := divide(100.0, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}
	result = x / y
	return result, nil

}

//func PrintInfo(a Animal) {
//	fmt.Println(a.says(), a.NumberOfLegs())
//}
//func (d Dog) says() string {
//	return "woof"
//}
//func (d Dog) NumberOfLegs() int {
//	return 4
//}

//func saySomething() (string, int) {
//	return "SomeThing", 2
//}
//
//func changeUsingpointer(s *string) {
//	log.Print("my Satring  set to ", s)
//
//	*s = "changed"
//}

package db

import (
	"fmt"
	"math/rand"
	"time"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says: Woof!"
}

func (c Cat) Speak() string {
	return c.Name + " says: Meow!"
}

func GenerateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func Multiply(a, b int) int {
	return a * b
}

func GreetUser(name string) string {
	if name == "" {
		return "Hello, stranger!"
	}
	return "Hello, " + name + "!"
}

func CheckEvenOdd(number int) string {
	if number%2 == 0 {
		return fmt.Sprintf("%d is even", number)
	}
	return fmt.Sprintf("%d is odd", number)
}

func Dummy() {
	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Dog{Name: "Rex"},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	number := GenerateRandomNumber(1, 100)
	fmt.Println("Random Number:", number)
	fmt.Println(CheckEvenOdd(number))

	product := Multiply(GenerateRandomNumber(1, 10), 5)
	fmt.Printf("Random product: %d\n", product)

	fmt.Println(GreetUser("John"))
	fmt.Println(GreetUser(""))

	// Loop 5 times and print a message
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration #%d\n", i)
	}
}

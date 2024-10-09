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

func CalculateFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * CalculateFactorial(n-1)
}

func GetFibonacciSeries(n int) []int {
	fib := make([]int, n)
	for i := 0; i < n; i++ {
		if i <= 1 {
			fib[i] = i
		} else {
			fib[i] = fib[i-1] + fib[i-2]
		}
	}
	return fib
}

func PrintSlice(slice []int) {
	for _, value := range slice {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func Dummy() {
	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Dog{Name: "Rex"},
		Cat{Name: "Mittens"},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	number := GenerateRandomNumber(1, 100)
	fmt.Println("Random Number:", number)
	fmt.Println(CheckEvenOdd(number))

	product := Multiply(GenerateRandomNumber(1, 10), 5)
	fmt.Printf("Random product: %d\n", product)

	factorial := CalculateFactorial(5)
	fmt.Printf("Factorial of 5: %d\n", factorial)

	fibSeries := GetFibonacciSeries(10)
	fmt.Print("Fibonacci Series (10 terms): ")
	PrintSlice(fibSeries)

	fmt.Println(GreetUser("John"))
	fmt.Println(GreetUser(""))

	// Loop 5 times and print a message
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration #%d\n", i)
	}

	// Calculate the sum of first 10 natural numbers
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("Sum of first 10 natural numbers: %d\n", sum)

	// Generate and print 5 random numbers
	fmt.Print("Five Random Numbers: ")
	for i := 0; i < 5; i++ {
		fmt.Print(GenerateRandomNumber(1, 50), " ")
	}
	fmt.Println()

	// Check if a number is prime
	for i := 1; i <= 20; i++ {
		if IsPrime(i) {
			fmt.Printf("%d is a prime number\n", i)
		} else {
			fmt.Printf("%d is not a prime number\n", i)
		}
	}
}

// Function to check if a number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

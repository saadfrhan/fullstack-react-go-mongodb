package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var myName string = "John Doe"

	const mySecondName string = "Jane Doe"
	// mySecondName = "Jane Doe" // This will throw an error because mySecondName is a constant

	myThirdName := "John Doe" // This is a shorthand way of declaring a variable i.e. var myThirdName string = "John Doe"
	fmt.Println(myName)
	fmt.Println(mySecondName)
	fmt.Println(myThirdName)
}

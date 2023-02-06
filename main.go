package main

import "fmt"

// main function
func main() {

    // Println function is used to
    // display output in the next line
    fmt.Println("Enter the first number: ")

    // var then variable name then variable type
    var first int

    // Taking input from user
    fmt.Scanln(&first)
    fmt.Println("Enter the second number: ")
    var second int
    fmt.Scanln(&second)

    // Print function is used to
    // display output in the same line
    fmt.Print("The result is: 444444")

    // Addition of two string
    fmt.Println(first + second)
}
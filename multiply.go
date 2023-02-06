package main

import "fmt"

func main(){
 fmt.Println("Enter the first number : ")
 var num1 int
 fmt.Scanln(&num1)
 fmt.Println("Enter the second number : ")
 var num2 int
 fmt.Scanln(&num2)
 fmt.Println("The result of multiplication is : ",num1*num2)

}
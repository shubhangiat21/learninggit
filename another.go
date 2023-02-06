package main

import "fmt"

func main(){
var name string
fmt.Println("What's your name?")
fmt.Scanln(&name)
fmt.Print("Hi "+name+", nice to meet you!")
}
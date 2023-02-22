package main

import "fmt"

func main() {
	var name string
	fmt.Println("please enter your name")
	fmt.Scanln(&name)
	fmt.Println("i love", name)
}

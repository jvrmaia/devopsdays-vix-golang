package main

import "fmt"

func oi() {
	defer func() { fmt.Println("fim oi") }()

	fmt.Println("oi")
}

func main() {
	defer func() { fmt.Println("fim main") }()

	oi()

	fmt.Println("olá")
}

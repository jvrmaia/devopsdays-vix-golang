package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func imprime(any interface{}) {
	fmt.Println(any)
}

func main() {
	fmt.Println(soma(1, 1))

	imprime(1)
	imprime("teste")
	imprime(byte(3))
}

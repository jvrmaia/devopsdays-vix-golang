package main

import "fmt"

func main() {
	hello := "Olá mundo"
	fmt.Println(hello)

	longText := `linha 1
linha 2
linha 3
linha 4
linha 5
linha 6`
	fmt.Println(longText)

	numeros := "9876543210"
	for i, v := range numeros {
		fmt.Println("posicao:", i, "| valor:", string(v))
	}

	fmt.Println([]byte(numeros))

	var b byte
	b = 47
	fmt.Println(string(b))

	b = 0
	fmt.Print(string(b))
	b = 1
	fmt.Print(string(b))
	b = 2
	fmt.Println(string(b))

	b = 253
	fmt.Print(string(b))
	b = 254
	fmt.Print(string(b))
	b = 255
	fmt.Println(string(b))
}

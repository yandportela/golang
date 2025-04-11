package main

import (
	"fmt"
)

func sayHello(nome string) {
	fmt.Println("Olá", nome)
}

func addNumbers(a, b int) int {
	return a + b
}
func main() {
	sayHello("Yan")
	resultado := addNumbers(1, 2)
	fmt.Println("Resultado: ", resultado)
}

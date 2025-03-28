package main

import "fmt"

func main() {
	a, b := 10, 3
	a++
	fmt.Println("A soma é: ", a + b)
	fmt.Println("A subtração é: ", a - b)
	fmt.Println("A multiplicação é: ", a * b)
	fmt.Println("A divisão é: ", a / b)
	fmt.Println("O resto da divisão é: ", a % b)
}

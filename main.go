package main

import (
	"fmt"
)

func main() {
	inteiros := []int{1, 2, 3, 4, 5}
	fmt.Println("Por favor, digite o primeiro número:")
	fmt.Scan(&inteiros[0])
	fmt.Println("Por favor, digite o segundo número:")
	fmt.Scan(&inteiros[1])
	fmt.Println("Por favor, digite o terceiro número:")
	fmt.Scan(&inteiros[2])
	fmt.Println("Por favor, digite o quarto número:")
	fmt.Scan(&inteiros[3])
	fmt.Println("Por favor, digite o quinto número:")
	fmt.Scan(&inteiros[4])
	fmt.Println("A soma dos números digitados é:", inteiros[0]+inteiros[1]+inteiros[2]+inteiros[3]+inteiros[4])
}
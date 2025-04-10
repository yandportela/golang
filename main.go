package main

import (
	"fmt"
)

func main() {
	var idade int
	fmt.Println("Por favor digite sua idade:")
	fmt.Scan(&idade)
	if idade < 0 || idade > 120 {
	for {
		fmt.Println("Idade inválida.")
		fmt.Println("Por favor digite uma idade válida:")
		fmt.Scan(&idade)
		if (idade > 0 && idade < 120) {
			break
		}
	}
	}
	if idade < 18 {
		fmt.Println("Você é menor de idade.")
	}	else if idade >= 18 && idade < 60 {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é idoso.")
	}
}
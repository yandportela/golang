package main

import (
	"fmt"
)

func main() {
	//age := 45
	/*fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)
	if age < 30 {
		fmt.Println("Menor que 30")
	} else if age < 40 {
		fmt.Println("Menor que 40")
	} else {
		fmt.Println("Não é menor que 40")
	}
*/
	nomes := []string{"Yan", "Maria", "Pedro", "Vinícius", "Eduardo"}
	for index, valor := range nomes {
		if index == 1 {
			fmt.Println("Continue após a posição: ", index, "e valor ", valor)
			continue
		}
		if index > 2 {
			fmt.Println("Sair após", index)
			break 
		}
		fmt.Println("Valor: ", valor)
	}
}
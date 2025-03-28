package main

import "fmt"

func main() {
	var numa, numb int
	fmt.Println("Escreva o primeiro número: ")
	fmt.Scan(&numa)
	fmt.Println("Escreva o segundo número: ")
	fmt.Scan(&numb)
	fmt.Println("A soma dos números é: ", numa + numb)
	fmt.Println("A subtração dos números é: ", numa - numb)
	fmt.Println("A multiplicação dos números é: ", numa * numb)
	fmt.Println("A divisão dos números é: ", numa / numb)
	fmt.Println("O resto da divisão dos números é: ", numa % numb)
	fmt.Print("Se a subtração for ", numa, " - ", numb," Será ", numb - numa, "\n")
	fmt.Print("Se a divisão for ", numa, " / ", numb," Será ", numb / numa)
	
}

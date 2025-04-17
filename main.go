package main

import ("fmt")

var nome string
var idade int

func dadosPessoa(nome string, idade int) (int, string) {
	if idade <= 0 || idade > 120 {
		return idade, "\n"+nome+" idade inválida"
	} else if idade <= 17 {
		return idade, "\n"+nome+" é menor de idade"
	} else if idade <= 59 {
		return idade, "\n"+nome + " é maior de idade"
	} else {
		return idade, "\n"+nome + " é idoso"
	}
}

func main() {
	fmt.Println("Por favor digite seu nome:")
	fmt.Scan(&nome)
	fmt.Println("Por favor digite sua idade:")
	fmt.Scan(&idade)
	fmt.Print("\n\n")
	fmt.Println(dadosPessoa(nome, idade))
}
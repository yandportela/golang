package main

import "fmt"

func main() {
	var usuario string
	var senha int
	fmt.Println("Digite seu usuário")
	fmt.Scanln(&usuario)
	if usuario == "admin" {
		fmt.Println("Digite sua senha")
		fmt.Scanln(&senha)
		if senha == 1234 {
			fmt.Println("Bem-vindo")
		} else {
			fmt.Println("Senha incorreta")
		}
	} else {
		fmt.Println("Usuário incorreto")
	}
}

package main

import (
	"fmt"
)

func main() {
	var saldo float64
	var continuar string
	var opcao string
	fmt.Print("Digite seu saldo inicial: ")
	fmt.Scan(&saldo)
	for {
		if saldo < 0 {
			fmt.Println("Saldo negativo. Por favor, insira um saldo válido.")
			fmt.Print("Digite seu saldo inicial: ")
			fmt.Scan(&saldo)
		} else {
			break
		}
	}

	for {
		for {
			fmt.Print("Você deseja sacar ou depositar? (s/d): ")
			fmt.Scan(&opcao)
			if opcao == "s" || opcao == "d" {
				break
			} else {
				fmt.Println("Opção inválida. Por favor, escolha 's' para sacar ou 'd' para depositar.")
			}
		}

		if opcao == "s" {
			var valorSaque float64
			fmt.Print("Digite o valor do saque: ")
			fmt.Scan(&valorSaque)
			if valorSaque > saldo {
				fmt.Println("Saldo insuficiente.")
			} else {
				saldo -= valorSaque
				fmt.Printf("Saque de %.2f realizado com sucesso. Saldo atual: %.2f\n", valorSaque, saldo)
			}
			fmt.Print("Você deseja realizar outra operação? (s/n): ")
			fmt.Scan(&continuar)
			if continuar == "s" {
				main()
			} else {
				fmt.Println("Obrigado por usar nosso sistema!")
				break
			}
		}

		if opcao == "d" {
			var valorDeposito float64
			for {
				fmt.Print("Digite o valor do depósito: ")
				fmt.Scan(&valorDeposito)
				if valorDeposito <= 0 {
					fmt.Println("Valor inválido. O depósito deve ser maior que zero.")
				} else {
					saldo += valorDeposito
					fmt.Printf("Depósito de %.2f realizado com sucesso. \nSaldo atual: %.2f\n", valorDeposito, saldo)
					break
				}
			}

			fmt.Print("Você deseja realizar outra operação? (s/n): ")
			fmt.Scan(&continuar)
			if continuar == "s" {
			} else {
				fmt.Println("Obrigado por usar nosso sistema!")
				break
			}
		}
	}
}

package main

import ("fmt")

var saldo, valorSaque, valorDeposito int
var opcao int

func sacar() {
	if saldo == 0 {
		fmt.Println("\nVocê não tem saldo suficiente.")
	} else {
	fmt.Println("\nQuanto de", saldo, "real(is) você quer sacar?")
	fmt.Scanln(&valorSaque)
	if valorSaque > saldo || valorSaque < 0 {
		fmt.Println("\nVocê não tem saldo suficiente.")
	} else {
		saldo -= valorSaque
		fmt.Println("\nVocê sacou", valorSaque, "real(is) e agora tem", saldo, "real(is).")
	}
}
}

func depositar() {
	fmt.Println("Quanto você quer depositar?")
	for {
		fmt.Scanln(&valorDeposito)
		if valorSaque < 0 {
			fmt.Println("Valor inválido. Tente novamente.")
			continue
		}
		break
	}
	saldo += valorDeposito
	fmt.Println("\nVocê depositou", valorDeposito, "real(is) e agora tem", saldo, "real(is).")
}

func mostrarSaldo() {
	fmt.Println("\nSeu saldo é:", saldo, "real(is).")
}
func main() {
	fmt.Println("Bem-vindo ao banco!")
	for {
		fmt.Println("\n\n\nEscolha uma opção: \n1. Sacar \n2. Depositar \n3. Mostrar saldo\n4. Sair")
		fmt.Scanln(&opcao)
		switch opcao {
		case 1:
			sacar()
		case 2:
			depositar()
		case 3:
			mostrarSaldo()
		case 4:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
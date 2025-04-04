package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	cumprimento := "Olá mundo!"
	palavras := []string{}
	fmt.Println(strings.Contains(cumprimento, "mundo"))
	fmt.Println(cumprimento)
	fmt.Println(strings.ReplaceAll(cumprimento, "mundo", "internet"))
	fmt.Println(strings.ToUpper(cumprimento))
	fmt.Println(strings.ToLower(cumprimento))
	fmt.Println(strings.Index(cumprimento, "mundo"))
	palavras = strings.Split(cumprimento, " ")
	fmt.Println(palavras)
	//
	idades := []int{50, 80, 10, 14, 17, 3, 47, 58, 90, 100}
	sort.Ints(idades)
	fmt.Println(idades)
	indice := sort.SearchInts(idades, 80)
	fmt.Println(indice)
	//
	nomes := []string{"Yan", "Pedro", "Vinicius", "Bruno", "Eduardo"}
	sort.Strings(nomes)
	fmt.Println(nomes)
	//
	x := 0

	for x <= 5 {
		fmt.Println(x)
		x++
	}
	for i := 0; i <= 5; i++ {
		fmt.Println("For 2: ", i)
	}
	for i := 0; i < len(nomes); i++ {
		fmt.Println("For 3: ", nomes[i])
	}
	for index, value := range nomes {
		fmt.Println("For 4: ", index, value)
	}
	for indice, valor := range idades {
		fmt.Println("O", indice, "º é ", valor)
	}
	superherois := []string{"Deadpool", "Homem Aranha", "Batman"}

	for index, value := range superherois {
		fmt.Println("O", index, "º superheroi é ", value)
	}
	for i := 0; i < len(superherois); i++ {
		fmt.Println(superherois[i])
	}
}

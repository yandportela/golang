package main

import "fmt"

func main() {
	var idades = [4]int{15, 16, 16, 16}
	nomes := [4]string{"Bruno", "Eduardo", "Pedro", "Vinicius"}
	fmt.Println(idades)
	fmt.Println(nomes)
	nomes[3] = "Morales"
	fmt.Println(nomes)
	// Slice
	var score = [4]int{100, 200, 300, 400}
	fmt.Println(score)
	score[1] = 2
	fmt.Println(score, len(score), cap(score))
	rangeOne := score[1:3]
	fmt.Println(rangeOne)
	rangeTwo := score[2:]
	fmt.Println(rangeTwo)
	rangeThree := score[:3]
	fmt.Println(rangeThree)

	var superherois = []string{"Deadpool", "Homem-Aranha", "Motoqueiro Fantasma"}
	fmt.Println(superherois)
	superherois = append(superherois, "Ben 10", "Ciborgue")
	fmt.Println(superherois, len(superherois), cap(superherois))
}
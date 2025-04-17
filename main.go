package main

import (
	"fmt"
)

func main() {
	// Mapa de idades
	AlunoIdade := make(map[string]int)
	AlunoIdade["Laura"] = 16
	AlunoIdade["Fabiano"] = 40
	AlunoIdade["Rauanny"] = 17
	AlunoIdade["Livia"] = 16
	AlunoIdade["Isabela"] = 15

	fmt.Println("Idade da Laura:", AlunoIdade["Laura"])

	// Mapa de notas
	notasAlunos := map[string]float64{
		"Laura":   8.5,
		"Fabiano": 10,
		"Rauanny": 7.4,
		"Isabela": 6.8,
		"Livia":   9.3,
	}

	for nome, nota := range notasAlunos {
		fmt.Printf("%s tirou a nota %.2f\n", nome, nota)
	}
}
package main

import (
	"fmt"
)

func main() {
	var numeros [5]int 

	for i := 0; i < len(numeros); i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&numeros[i])
	}

	soma :=0
   
	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}

	fmt.Printf("A soma dos números é: %d\n", soma)
}

 

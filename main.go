package main

import (
	"fmt"
)

func main() {
	var idade int
	fmt.Println("Informe sua idade")
	fmt.Scan(&idade)

	if idade < 18 {
		fmt.Println("você é menor de idade")
	} else if idade >=18 && idade <60 {
		fmt.Println("você é adulto")
	} else if idade >=60 {
		fmt.Println("voce é idoso")
}
}
package main

import (
	"fmt"
)
func dadosPessoa(idade int) (int, string) {
	valor := "menor de idade"
	if idade >= 18 {
		valor = "maior de idade"
	}
	return idade, valor
}
func main(){
	idade, condicao:= dadosPessoa(20)
	fmt.Println(idade, condicao)
}

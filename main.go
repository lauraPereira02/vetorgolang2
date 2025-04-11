package main

import (
	"fmt"
)

var saldo float32
var valor float32

func main() {
	var opcao int

	fmt.Print("Informe seu saldo inicial: ")
	fmt.Scan(&saldo)
	fmt.Println("Escolha uma opção: 1 para sacar, 2 para depositar")
	fmt.Print("Opção: ")
	fmt.Scan(&opcao)

	fmt.Print("Informe o valor:")
	fmt.Scan(&valor)

	if opcao == 1 {
		sacar()
	} else if opcao == 2 {
		depositar()
	} else {
		fmt.Println("Opção inválida.")
	} 
} 

func sacar() {
	if valor > saldo {
		fmt.Println("Saldo insuficiente.")
	} else {
		saldo -= valor
		fmt.Println("Saque realizado com sucesso! Saldo atual:", saldo)
	}
}

func depositar() {
	saldo += valor
	fmt.Println("Depósito realizado com sucesso! Saldo atual:", saldo)
}

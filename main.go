package main

import (
	"fmt"
)

func main() {
	var saldo float32
	var saque float32
	var deposito float32
	var decisao int
	var saldofinal float32
	var depositofinal float32
	fmt.Println("Informe seu saldo")
	fmt.Scan(&saldo)
	fmt.Println("Para sacar digite 1 ou digite 2 para depositar")
	fmt.Scan(&decisao)

	if decisao == 1 {
		fmt.Println("Valor desejado para sue saque:")
		fmt.Scan(&saque)
		saldofinal = saldo - saque
		fmt.Println("seu saldo final é ",saldofinal)

		} else if decisao == 2 {
		fmt.Println("Valor desejado para deposito:")
		fmt.Scan(&deposito)
		depositofinal = saldo + deposito
		fmt.Println("Seu saldo final é ", depositofinal)
		}
	}

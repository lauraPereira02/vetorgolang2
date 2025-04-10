package main

import (
	"fmt"
)

func main() {
	var numero int
	var numero02 int
   var numero03 int
   var numero04 int
   var numero05 int

	fmt.Print("Digite um numero: ")
	fmt.Scan(&numero)
	fmt.Print("Digite outro numero: ")
	fmt.Scan(&numero02)
   fmt.Print("Digite um numero: ")
	fmt.Scan(&numero03)
   fmt.Print("Digite um numero: ")
	fmt.Scan(&numero04)
   fmt.Print("Digite um numero: ")
	fmt.Scan(&numero05)
	resultado := numero + numero02 + numero03 + numero04 + numero05
	fmt.Printf("A soma dos numeros é: %d\n", resultado)
}






 

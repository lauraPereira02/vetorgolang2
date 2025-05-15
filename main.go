package main

import "fmt"

type carro struct{
	marca string
	ano int
	velocidade int
}
 
func mostrarDados(c carro){
	fmt.Println("marca do carro: ", c.marca)
	fmt.Println("ano do carro: ", c.ano)
	fmt.Println("velocidade do carro: ", c.velocidade)
}

func main() {
	c1 := carro {"ford", 2020, 190}
	mostrarDados(c1)
}



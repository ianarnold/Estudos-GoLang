package main

import "fmt"

type hotdog int

var cachorroquente hotdog

var qualquer string

func main() {

	texto := "oi"
	x := 10
	y := 20.5
	sim := true
	cachorroquente = 15

	fmt.Print("Digite qualquer coisa: ")
	fmt.Scanln(&qualquer)
	

	fmt.Printf("Valor: %v \t Tipo: %T \n", texto, texto)
	fmt.Printf("Valor: %v \t Tipo: %T \n", x, x)
	fmt.Printf("Valor: %v \t Tipo: %T \n", y, y)
	fmt.Printf("Valor: %v \t Tipo: %T \n", sim, sim)
	fmt.Printf("Valor: %v \t Tipo: %T \n", cachorroquente, cachorroquente)
	fmt.Printf("Valor: %v \t Tipo: %T \n", qualquer, qualquer)




}
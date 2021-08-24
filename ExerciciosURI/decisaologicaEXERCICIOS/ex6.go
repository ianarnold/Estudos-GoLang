/*

Entrar via teclado com três valores distintos. Exibir o
maior deles.

ta cheio de bobeira to usando esse aq p experimentar coisa doidao

*/
package main

import (
	"fmt"
	//"os"
)

var (
	n1 int
	n2 int
	n3 int
)

func main() {

	receber()
	comparar()
	tchau()

}

func receber() {
	
	fmt.Print("Digite o 1º valor: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite o 2º valor: ")
	fmt.Scanln(&n2)
	fmt.Print("Digite o 3º valor: ")
	fmt.Scanln(&n3)

}

func comparar() {

	if n1 > n2 && n1 > n3 {
		fmt.Println(n1, "é o maior valor digitado.")
	} else if n2 > n1 && n2 > n3 {
		fmt.Println(n2, "é o maior valor digitado.")
	} else if n3 > n1 && n3 > n2 {
		fmt.Println(n3, "é o maior valor digitado.")
	} else if n1 == n2 && n2 == n3 {
		fmt.Println("Valores iguais")
	}

}

func tchau() {
	
	for i:= 1; i <= 10; i++ {
		fmt.Printf("%vº Tchau\n", i)
	}
}
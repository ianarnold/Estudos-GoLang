/*

Calcular e exibir a média aritmética de quatro
valores quaisquer que serão digitados

*/

package main

import (
	"fmt"
)

var (
	n1 float64
	n2 float64
	n3 float64
	n4 float64
	media float64
)

func main() {

	fmt.Print("Digite a 1º nota: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite a 2º nota: ")
	fmt.Scanln(&n2)
	fmt.Print("Digite a 3º nota: ")
	fmt.Scanln(&n3)
	fmt.Print("Digite a 4º nota: ")
	fmt.Scanln(&n4)
	media = (n1 + n2 + n3 + n4) / 4
	fmt.Println("Média das notas: ", media)

}
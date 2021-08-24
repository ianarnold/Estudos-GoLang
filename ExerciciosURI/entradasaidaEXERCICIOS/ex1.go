/*

Entrar via teclado com a base e a altura de um
retângulo, calcular e exibir sua área.

*/

package main

import "fmt"

var (
	area float64
	base float64
	altura float64
)

func main() {

	fmt.Print("Digite a base do retângulo: ")
	fmt.Scanln(&base)
	fmt.Print("Digite a altura do retângulo: ")
	fmt.Scanln(&altura)
	area = (base * altura)
	fmt.Println("Área do retângulo:", area, "cm")

}
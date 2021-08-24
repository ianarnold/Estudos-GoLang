/*

A partir dos valores da base e altura de um
triângulo, calcular e exibir sua área. 

*/

package main

import "fmt"

var (
	area float64
	base float64
	altura float64
)

func main() {

	fmt.Print("Digite a base do triangulo: ")
	fmt.Scanln(&base)
	fmt.Print("Digite a altura do triangulo: ")
	fmt.Scanln(&altura)
	area = (base * altura) / 2
	fmt.Println("Área do triângulo:", area, "cm³")

}

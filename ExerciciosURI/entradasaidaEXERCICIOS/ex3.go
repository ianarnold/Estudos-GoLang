/*

Calcular e exibir a área de um quadrado a partir do
valor de sua diagonal que será digitado.

*/

package main

import "fmt"

var (
	diagonal float64
	area float64
)

func main() {

	fmt.Print("Digite o valor da diagonal: ")
	fmt.Scanln(&diagonal)
	area = ((diagonal / 2) * (diagonal / 2)) * 2
	fmt.Println("Área: ", area)

}
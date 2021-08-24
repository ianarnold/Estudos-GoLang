/*

Calcular e exibir a área de um retângulo, a partir
dos valores da base e altura que serão digitados. Se
a área for maior que 100, exibir a mensagem
“Terreno grande”. 

*/

package main

import "fmt"

var (
	base int
	altura int
	area int
)

func main() {

	fmt.Print("Digite o tamanho da base: ")
	fmt.Scanln(&base)
	fmt.Print("Digite o tamanho da altura: ")
	fmt.Scanln(&altura)
	area = base * altura

	if area >= 100 {
		fmt.Println("Terreno grande")
	} else {
		fmt.Println("Terreno pequeno")
	}
	fmt.Printf("Área do retângulo: %vm²", area)

}
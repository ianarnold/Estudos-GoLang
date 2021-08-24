/*

Calcular e exibir a área de um quadrado, a partir
do valor de sua aresta que será digitado. 

*/

package main

import(
	"fmt"
	"math"
)

var (
	area float64
	lado float64
)

func main() {

	fmt.Print("Digite o lado do quadrado: ")
	fmt.Scanln(&lado)
	area = math.Pow(lado, 2)
	fmt.Println("Área:", area, "cm²")


}
/*

A partir do diâmetro de um círculo que será
digitado, calcular e exibir sua área

*/

package main

import (
	"fmt"
	"math"
)

var (
	diametro float64
	area float64
	raio float64
)

func main() {

	fmt.Print("Digite o diametro da esfera: ")
	fmt.Scanln(&diametro)
	raio = diametro / 2
	area = (4 * 3.14) * math.Pow(raio, 2)
	fmt.Printf("Área da esfera: %vm²", area)

}
/*

Calcular e exibir o volume de uma esfera a partir do
valor de seu diâmetro que será digitado.

*/

package main

import (
	"fmt"
	"math"
)

var (
	diametro float64
	volume float64
	raio float64
)

func main() {

	fmt.Print("Digite o diametro da esfera: ")
	fmt.Scanln(&diametro)
	raio = diametro / 2
	volume = 4  * 3.14 * (math.Pow(raio, 3) / 3)
	fmt.Println("Volume da esfera: ", volume, "cm³")

}
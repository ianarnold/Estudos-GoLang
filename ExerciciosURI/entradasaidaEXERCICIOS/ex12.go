/*

Calcular e exibir o volume de um cone a partir dos
valores da altura e do raio da base que serão
digitados.

*/

package main

import (
	"fmt"
	"math"
)
	
var (
	volume float64
	altura float64
	raio float64
)

func main() {

	fmt.Print("Digite a altura: ")
	fmt.Scanln(&altura)
	fmt.Print("Digite o raio: ")
	fmt.Scanln(&raio)
	volume = (3.14 * math.Pow(raio, 2) * altura) / 3
	fmt.Println("Volume: ", volume)

}
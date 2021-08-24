/*

Calcular e exibir a velocidade final (em km/h) de
um automóvel, a partir dos valores da velocidade
inicial (em m/s), da aceleração (m/s²) e do tempo
de percurso (em s) que serão digitados. 

*/

package main

import (
	"fmt"
	"math"
)

var (
	vInicial float64
	acel float64
	tempo float64
	vFinal float64
)

func main() {

	fmt.Print("Digite a velocidade inicial em m/s: ")
	fmt.Scanln(&vInicial)
	fmt.Print("Digite a aceleração em m/s²: ")
	fmt.Scanln(&acel)
	fmt.Print("Digite o tempo em s: ")
	fmt.Scanln(&tempo)
	vFinal = vInicial + (acel * tempo)
	vFinal = vFinal * 3.6
	fmt.Printf("Velocidade final: %vkm/h", math.Round(vFinal))

}
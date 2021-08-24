/*

Calcular e exibir a tensão de um determinado
circuito eletrônico a partir dos valores da resistência
e corrente elétrica que serão digitados. Utilize a lei
de Ohm. 

*/

package main

import "fmt"

var (
	tensao float64
	resistencia float64
	corrente float64
)

func main() {

	fmt.Print("Digite o valor da resistência: ")
	fmt.Scanln(&resistencia)
	fmt.Print("Digite o valor da corrente elétrica: ")
	fmt.Scanln(&corrente)	
	tensao = resistencia * corrente
	fmt.Printf("Valor da tensão elétrica: %vV", tensao)

}
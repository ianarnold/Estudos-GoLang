/*

Sabendo que uma milha marítima equivale a um
mil, oitocentos e cinquenta e dois metros e que um
quilômetro possui mil metros, fazer um programa
para converter milhas marítimas em quilômetros

*/

package main

import "fmt"

var (
	milhas float64
	metros float64
	km float64
)

func main() {

	fmt.Print("Digite o valor em milhas maritimas: ")
	fmt.Scanln(&milhas)
	metros =  milhas * 1852
	km = metros / 1000
	fmt.Printf("Valor em metros: %v \nValor em kilometros: %v", metros, km)

}
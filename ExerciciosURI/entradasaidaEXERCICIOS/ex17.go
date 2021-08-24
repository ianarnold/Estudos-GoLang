/*

Entrar via teclado com dois valores quaisquer “X” e
“X”. Calcular e exibir o cálculo XY (“X” elevado a
“Y”). Pesquisar as funções Exp e Ln. 

*/

package main

import (
	"fmt"
	"math"
)

var (
	x float64 
	y float64
	total float64  
)

func main() {

	fmt.Print("Digite o valor de X: ")
	fmt.Scanln(&x)
	fmt.Print("Digite o valor de Y: ")
	fmt.Scanln(&y)
	total = math.Pow(x, y)
	fmt.Println("Total: ", total)
	
}
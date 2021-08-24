/*

Entrar com o peso e a altura de uma determinada pessoa.
Após a digitação, exibir se esta pessoa está ou não com seu
peso ideal. Veja tabela da relação peso/altura².

Relação peso/altura2 		Mensagem
	R< 20 				Abaixo do peso
20 <= R < 25		 	Peso ideal
	R >= 25 		Acima do peso

*/

package main

import (
	"fmt"
	"math"
)

var (
	peso float64
	altura float64
	pesoideal float64
)

func main() {

	recebeValores()
	relacaoPesoAltura()

}

func recebeValores() {

	fmt.Print("Digite o seu peso: ")
	fmt.Scanln(&peso)
	fmt.Print("Digite a sua altura: ")
	fmt.Scanln(&altura)

}

func relacaoPesoAltura() {
	
	pesoideal = peso / (math.Pow(altura, 2))
	fmt.Printf("Seu IMC é: %.2f \n", pesoideal)
	if pesoideal < 20 {
		fmt.Println("Abaixo do peso")
	} else if pesoideal >= 20 && pesoideal < 25 {
		fmt.Println("Peso ideal")
	} else if pesoideal >= 25 {
		fmt.Println("Acima do peso")
	} else {
		fmt.Println("ERROR")
	}

}
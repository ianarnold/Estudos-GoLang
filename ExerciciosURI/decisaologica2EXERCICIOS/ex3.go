/*

Entrar com o peso, o sexo e a altura de uma determinada
pessoa. Após a digitação, exibir se esta pessoa está ou não
com seu peso ideal. Veja tabela da relação peso/altura2
.

Relação peso/altura2                Mensagem
	Feminino
								
	R< 19 						Abaixo do peso
19 <= R < 24 					Peso ideal
R >= 24 						Acima do peso


Relação peso/altura2				Mensagem
	Masculino

	R< 20 						Abaixo do peso
20 <= R < 25 					Peso ideal
R >= 25 						Acima do peso

*/

package main

import (
	"fmt"
	"math"
	"strings"
)

var (
	peso float64 
	altura float64
	imc float64
	sexo string
)

func main() {

	fmt.Print("Digite o peso: ")
	fmt.Scanln(&peso)
	fmt.Print("Digite a altura em m: ")
	fmt.Scanln(&altura)
	fmt.Print("MASCULINO / FEMININO \nDigite o sexo: ")
	fmt.Scanln(&sexo)
	sexo = strings.ToLower(sexo)
	verificaSexo()

}

func verificaSexo() {

	switch sexo {
	case "masculino":
		verificaImcMasc()
	case "feminino":
		verificaImcFem()
	default:
		fmt.Println("Digite um sexo válido.")
	}

}

func verificaImcMasc() {

	imc = peso / (math.Pow(altura, 2))
	fmt.Printf("Seu IMC é: %.2f \n", imc)
	if imc < 19 {
		fmt.Println("Abaixo do peso")
	} else if imc >= 20 && imc < 25 {
		fmt.Println("Peso ideal")
	} else if imc >= 25 {
		fmt.Println("Acima do peso")
	} else {
		fmt.Println("ERROR")
	}

}

func verificaImcFem() {

	imc = peso / (math.Pow(altura, 2))
	fmt.Printf("Seu IMC é: %.2f \n", imc)
	if imc < 20 {
		fmt.Println("Abaixo do peso")
	} else if imc >= 19 && imc < 24 {
		fmt.Println("Peso ideal")
	} else if imc >= 24 {
		fmt.Println("Acima do peso")
	} else {
		fmt.Println("ERROR")
	}

}


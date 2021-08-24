/*

A partir de 3 (três) valores que serão digitados, verificar se
formam ou não um triângulo. Em caso positivo, exibir sua
classificação: “Isósceles, escaleno ou equilátero”. Um
triângulo escaleno possui todos os lados diferentes, o
triângulo isósceles, dois lados iguais e o equilátero, todos os
lados iguais. Para existir triângulo é necessário que a soma
de dois lados quaisquer seja maior que o outro, isto, para os
três lados. 

*/

package main

import "fmt"

var (
	lado1 float64
	lado2 float64
	lado3 float64
)

func main() {

	fmt.Print("Digite o tamanho do 1º lado: ")
	fmt.Scanln(&lado1)
	fmt.Print("Digite o tamanho do 2º lado: ")
	fmt.Scanln(&lado2)
	fmt.Print("Digite o tamanho do 3º lado: ")
	fmt.Scanln(&lado3)
	verificaTriangulo()

}

func verificaTriangulo() {

	somal1l2 := lado1 + lado2
	somal1l3 := lado1 + lado3
	somal2l3 := lado2 + lado3

	if  somal1l2 > lado3 && somal1l3 > lado2 && somal2l3 > lado1 {
		fmt.Println("Forma triângulo")
		if lado1 != lado2 && lado1 != lado3 && lado2 != lado3 {
			fmt.Println("Triângulo escaleno")
		} else if lado1 == lado2 && lado2 == lado3 {
			fmt.Println("Triângulo equilátero")
		} else {
			fmt.Println("Triângulo isósceles")
		}
	} else {
		fmt.Println("Não forma triângulo")
	}

}


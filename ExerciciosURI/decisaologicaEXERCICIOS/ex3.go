/*

Entrar com dois valores quaisquer. Exibir o maior
deles, se existir, caso contrário, enviar mensagem
avisando que os números são idênticos. 

*/

package main

import "fmt"

var n1 int
var n2 int

func main() {

	fmt.Print("Digite o 1º valor: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite o 2º valor: ")
	fmt.Scanln(&n2)
	if n1 > n2 {
		fmt.Println(n1, "maior que", n2)
	} else if n2 > n1 {
		fmt.Println(n2, "maior que", n1)
	} else if n1 == n2 {
		fmt.Println("Valores iguais")
	}

}
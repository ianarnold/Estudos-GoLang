/*

Entrar via teclado, com dois valores distintos. Exibir o
maior deles. 

*/

package main

import "fmt"

var (
	n1 int
	n2 int
)

func main() {

	fmt.Print("Digite o 1º valor: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite o 2º valor: ")
	fmt.Scanln(&n2)

	if n1 > n2 {

		fmt.Printf("%v é maior que %v", n1, n2)
		
	} else if n2 > n1 {

		fmt.Printf("%v é maior que %v", n2, n1)

	} else if n1 == n2 {

		fmt.Println("Os valores digitados são iguais")

	}

}
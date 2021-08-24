/*

1. Calcular e exibir a soma dos “N” primeiros valores da
sequência abaixo. O valor “N” será digitado, deverá ser
positivo, mas menor ou igual a 20 (vinte). Caso o valor
não satisfaça a restrição, enviar mensagem de erro e
solicitar o valor novamente.



1				2				3				4
2				3				4				5

*/

package main

import "fmt"

var (
	n1 int
	total int
	totalSoma int
)

func main() {

	doWhile()
	somarValores()

}

func doWhile() {

	fmt.Print("Digite o tamanho: ")
	fmt.Scanln(&n1)
	for {
		if n1 <= 0 || n1 >= 20 {
			fmt.Println("ERRO DIGITE NOVAMENTE!!!")
			fmt.Print("Digite o tamanho: ")
			fmt.Scanln(&n1)
		} else {
			break
		}
	}

}

func somarValores() {

	for cont := 1; cont <= n1; cont++ {
		total = cont + (cont + 1)
		totalSoma += total
		fmt.Printf("%vº: %v + %v = %v\n", cont, cont, cont + 1, total)
	}
	fmt.Println("Total das somas: ", totalSoma)

}


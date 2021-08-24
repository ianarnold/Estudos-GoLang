/*

Calcular e exibir a soma dos “N” primeiros valores da
sequência abaixo. O valor “N” será digitado, deverá ser
positivo, mas menor ou igual a 50 (cinquenta). Caso o
valor não satisfaça a restrição, enviar mensagem de
erro e solicitar o valor novamente.


2		5		10		17
1		8		27		64

*/

package main

import "fmt"

var (
	tamanho int
	total int
)

func main() {

	fmt.Print("Digite o tamanho: ")
	fmt.Scanln(&tamanho)
	for {
		if tamanho <= 0 || tamanho > 50 {
			fmt.Print("Digite o tamanho: ")
			fmt.Scanln(&tamanho)
		} else {
			break
		}
	}
	for cont := 0; cont <= tamanho; cont++ {
		
	}

}
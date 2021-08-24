/*

Entrar via teclado com o valor de cinco produtos.
Após as entradas, digitar um valor referente ao
pagamento da somatória destes valores. Calcular e
exibir o troco que deverá ser devolvido

*/

package main

import "fmt"

var (
	p1 float64
	p2 float64
	p3 float64
	p4 float64
	p5 float64
	valorpago float64
	total float64
	troco float64
)

func main() {

	fmt.Print("Digite o valor do produto: ")
	fmt.Scanln(&p1)
	fmt.Print("Digite o valor do produto: ")
	fmt.Scanln(&p2)
	fmt.Print("Digite o valor do produto: ")
	fmt.Scanln(&p3)
	fmt.Print("Digite o valor do produto: ")
	fmt.Scanln(&p4)
	fmt.Print("Digite o valor do produto: ")
	fmt.Scanln(&p5)
	fmt.Print("Digite o total pago: ")
	fmt.Scanln(&valorpago)
	total = p1 + p2 + p3 + p4 + p5
	
	if total > valorpago {
		fmt.Println("Valor pago insuficiente")
	} else {
	troco = valorpago - total
	fmt.Printf("Troco: R$%v", troco)
	}
	
}
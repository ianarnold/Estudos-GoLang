/*

Uma escola com cursos em regime semestral,
realiza 2 (duas) avaliações durante o semestre e
calcula a média do aluno, da seguinte maneira:
media = (p1 + 2p2) / 3. Fazer um programa para
entrar via teclado com os valores das notas (P1 e
P2) e calcular a média. Exibir a situação final do
aluno (“Aprovado ou Reprovado”), sabendo que a
média de aprovação é igual a 5 (cinco).

*/

package main

import (
	"fmt"
)

var (
	p1 float64
	p2 float64
	media float64
)

func main() {

	receberNotas()
	gerarMedia()
	situacaoFinal()

}

func receberNotas() {

	fmt.Print("Digite o valor da p1: ")
	fmt.Scanln(&p1)
	fmt.Print("Digite o valor da p2: ")
	fmt.Scanln(&p2)

}

func gerarMedia() {

	media = (p1 + (p2 + p2)) / 3

}

func situacaoFinal() {

	// %.2f para delimitar o float a 2 casas apos a virgula
	if media >= 5 {
		fmt.Printf("Aprovado com a média: %.2f", media)
	} else {
		fmt.Printf("Reprovado com a média: %.2f", media)
	}

}
/*

A prefeitura de uma cidade fez uma pesquisa entre
seus 500 (quinhentos) habitantes, coletando dados
sobre o salário e o número de filhos. A prefeitura deseja
saber:

a) a média dos salários da população;
b) a média dos números de filhos;
c) o maior salário;
d) a porcentagem de pessoas com salários até
R$ 1.000,00.

*/

package main

import "fmt"

var (
	salario float64
	filhos float64
	totalSal float64
	totalFil float64
	mediaSal float64
	mediaFil float64
	maiorSal float64
	porcenMil float64
	maior float64
	cont2 float64
)

func main() {

	recebeValores()

}

func recebeValores() {

	maior = 0
	for cont := 0; cont < 500; cont++ {
		fmt.Print("Digite o seu salário: ")
		fmt.Scanln(&salario)
		fmt.Print("Digite a quantidade de filhos que você possuí : ")
		fmt.Scanln(&filhos)
		totalSal += salario
		totalFil += filhos
		if salario > maior {
			maior = salario
		} else if salario <= 1000 {
			cont2++
		}
	}
	mediaFil = totalFil / 500
	mediaSal = totalSal / 500
	porcenMil = (cont2 * 100) / 500
	fmt.Println("A média dos salários é:", mediaSal)
	fmt.Println("A média dos filhos é:", mediaFil)
	fmt.Println("O maior salário é:", maior)
	fmt.Println("A porcentagem de pessoas com salário até R$1000,00 é:", porcenMil)

}
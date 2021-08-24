package main

import "fmt"

//SOMA DOIS VALORES
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// REALIZA DIVERSOS CALCULOS
func calculos(n1, n2 int8) (int8, int8, int8) {
	
	// FUNÇÃO COM VARIOS RETORNOS
	soma := n1 + n2
	subtracao := n1 - n2
	multi := n1 * n2
	return soma, subtracao, multi

}

func main() {

	// Atribui dois valores a função soma e exibe
	valor := somar(2, 2)
	fmt.Println(valor)


	//EXIBINDO TODOS RETORNOS DA FUNÇÃO
	resultadoSoma, resultadoSub, resultadoMulti := calculos(10, 10)
	fmt.Println(resultadoSoma, resultadoSub, resultadoMulti)

	// RETORNA APENAS UMA SUBFUNÇÃO DAS 3, O UNDERLINE "_" SERVE PARA "IGNORAR" AQUELA PARTE DO CODIGO.
	_, _, multiplicacao := calculos(5, 3)
	fmt.Println(multiplicacao)

}
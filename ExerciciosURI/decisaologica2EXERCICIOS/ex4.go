/*

 A partir dos valores da aceleração (a em m/s²), da
velocidade inicial (v0 em m/s) e do tempo de percurso (t em
s). Calcular e exibir a velocidade final de automóvel em
km/h. Exibir mensagem de acordo com a tabela abaixo.
Fórmula para o cálculo da velocidade em m/s: V = v0 + a. t.

Velocidade em Km/h(V)			Mensagem

V < = 40 					Veículo muito lento
40 < V <= 60 				Velocidade permitida
60 < V <= 80 				Velocidade de cruzeiro
80 < V <= 120 				Veículo rápido
R >= 120 					Veículo muito rápido


*/

package main

import (
	"fmt"
)

var (
	vInicial float64
	acel float64
	vFinal float64
	tempo float64
)

func main() {

	fmt.Print("Digite a aceleração:")
	fmt.Scanln(&acel)
	fmt.Print("Digite a velocidade inicial:")
	fmt.Scanln(&vInicial)
	fmt.Print("Digite o tempo do percurso:")
	fmt.Scanln(&tempo)
	calculoVel()

}

func calculoVel() {

	vFinal = vInicial + (acel * tempo)
	vFinal = vFinal * 3.6
	
	if vFinal <= 40 {
		fmt.Println("Veículo muito lento")
	} else if vFinal > 40 && vFinal <= 60 {
		fmt.Println("Velocidade permitida")
	} else if vFinal > 60 && vFinal <= 80 {
		fmt.Println("Velocidade de cruzeiro")
	} else if vFinal > 80 && vFinal <= 120 {
		fmt.Println("Veículo rápido")
	} else if vFinal > 120 {
		fmt.Println("Veículo muito rápido")
	} else {
		fmt.Println("ERRO")
	}

}
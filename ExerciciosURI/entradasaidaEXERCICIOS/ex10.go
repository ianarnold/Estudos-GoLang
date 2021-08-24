/*

Entrar via teclado com o valor de uma temperatura
em graus Celsius, calcular e exibir sua temperatura
equivalente em Fahrenheit e em Kelvin

*/

package main

import "fmt"

var (
	celsius float64
	fahrenheit float64
	kelvin float64
)

func main() {

	fmt.Print("Digite o valor em Cº: ")
	fmt.Scanln(&celsius)
	fahrenheit = (celsius * (9 / 5)) + 32
	kelvin = celsius + 273.15
	fmt.Printf("Valor em Celsius: %vCº \nValor em Fahrenheit: %vFº \nValor em Kelvin: %vKº", celsius, fahrenheit, kelvin)

}
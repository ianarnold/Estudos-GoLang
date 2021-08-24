/*

- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true
- Agora demonstre os valores nestas variáveis, com:
    1. Uma única declaração print
    2. Múltiplas declarações print

*/


package main

import (
	"fmt"
)

func main() {

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T \n", y, y)
	fmt.Printf("%v %T", z, z)

	// %v retorna o valor guardado na variavel
	// %T retorna o tipo

}

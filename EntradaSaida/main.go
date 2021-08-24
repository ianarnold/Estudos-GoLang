package main

import (
	"fmt"
)

var x int

func main() {
    
	fmt.Print("Digite sua idade: ")
    fmt.Scanln(&x)

	if x > 114 {

		fmt.Println("IMPOSSIVEL MARRECO")
		
	} else if x >= 18 {

		fmt.Printf("Você é maior de idade com %d anos de idade \n",x)

	} else {

		fmt.Printf("Você é maior de idade com %v anos de idade",x)

	}
	
	
}

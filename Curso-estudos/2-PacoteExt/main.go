package main

import (
	"fmt"
	"github.com/badoux/checkmail"
)

var email string

func main() {

	fmt.Print("Digite seu email: ")
	fmt.Scanln(&email)

	err := checkmail.ValidateFormat(email)
	if err != nil {
        fmt.Println("Formato inválido")
    } else {
		fmt.Println("Formato valido")
	}
	
}
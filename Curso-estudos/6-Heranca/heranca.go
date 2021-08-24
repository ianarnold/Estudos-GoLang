// Herança entre aspas, o Go não possui algo sobre heranças mas algo parecido

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // Atribuindo ao struct "estudante" todos campos da struct "pessoa"
	curso     string
	faculdade string
}

func main() {
	fmt.Print("\n Heranças no Go? Quase isso. \n")

	var usuarioPessoa pessoa
	usuarioPessoa.nome = "Ian"
	usuarioPessoa.sobrenome = "Sales"
	usuarioPessoa.idade = 18
	usuarioPessoa.altura = 180
	fmt.Println(usuarioPessoa)
	///////////////////////////////////////////////////////

	var usuarioEstudante estudante
	usuarioEstudante.pessoa = usuarioPessoa
	usuarioEstudante.curso = "Sistemas de informação"
	usuarioEstudante.faculdade = "Unisanta"
	fmt.Println(usuarioEstudante.nome)

}

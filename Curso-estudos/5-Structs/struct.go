package main

import "fmt"

//Estruturando o tipo "pessoa"
type pessoa struct {
	nome     string
	idade    uint8
	endereco endereco //Uma struct dentro de outra
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Print("\nAula sobre struct \n \n")
	////////////////////////////////

	enderecoExemplo := endereco{"Rua João Ramalho", 192} // Definindo um endereço de exemplo
	var user pessoa                                      // Declarando variavel do tipo "PESSOA"
	user.nome = "Ian"                                    // Atribuindo nome ao user do tipo "pessoa"
	user.idade = 18                                      // Atribuindo idade ao user do tipo "pessoa"
	user.endereco = enderecoExemplo
	fmt.Println(user) // Exbindo na tela os valores de "user"
	/////////////////////////////////////////////////////////////////////////////////////

	user2 := pessoa{
		"adriano", // Atribuindo valores ao "user2" do tipo "pessoa" usando inferencia
		21,        // de tipos
		enderecoExemplo,
	}
	fmt.Println(user2) // Exibindo os valores de "user2"
	//////////////////////////////////////////////////////////////////////////////////////////

	user3 := pessoa{nome: "Marta"} // Usando inferencia de tipos e apenas inserindo um valor no usuário. A idade vai como 0 (zero value) e endereco tb
	fmt.Println(user3)             // Exibindo os valores de "user3"

}

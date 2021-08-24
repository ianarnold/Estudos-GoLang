package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS)
	//Retorna o SO da maquina = WINDOWS


	fmt.Println(runtime.GOARCH)
	// Retorna o proc

}

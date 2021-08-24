package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	

	// MOSTRAR A SEQUENCIA DESEJADA ATRAVES DO FOR
	for cont := 0; cont <= 300; cont++ {
		fmt.Println(f()) //AQUI PRINTA
	}

}

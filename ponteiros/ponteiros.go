package main

import "fmt"

func main() {
	fmt.Println("\tPonteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	//Ponteiro é um referência de memória

	var variavel3 int = 100
	var ponteiro *int

	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	variavel3 = 12351251

	//Isso é uma desreferenciação
	fmt.Println(*ponteiro)

}

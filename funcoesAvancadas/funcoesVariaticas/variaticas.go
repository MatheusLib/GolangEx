package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)

	}
}

func main() {
	totalsoma := soma(1, 2, 3, 4, 5, 6, 7, 89, 123, 615254, 12023)
	fmt.Println(totalsoma)

	escrever("os numeros da função são: ", 10, 2, 3, 5, 7, 123, 651)
}

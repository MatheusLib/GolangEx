package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá Mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando Parâmetros")

	funcComRetorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("parametro da função")

	fmt.Println(funcComRetorno)
}

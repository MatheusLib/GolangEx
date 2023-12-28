package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

// Voce está pegando todos os campos que estão dentro de pessoa e jogando no type de estudante
func main() {
	fmt.Println("Aula de Herança")
	p1 := pessoa{"João", "Pedro", 26, 173}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
}

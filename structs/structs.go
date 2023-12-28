package main

import "fmt"

type usuario struct {
	nome        string
	idade       uint8
	localizacao endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u)
	u.nome = "Matheus"
	u.idade = 18
	fmt.Println(u)

	enderecoExemplo := endereco{"rua blablabla", 123}

	// Outra maneira de declarar
	u2 := usuario{"Libanio", 26, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{idade: 12}

	fmt.Println(u3)

}

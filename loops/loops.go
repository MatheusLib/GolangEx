package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		time.Sleep(time.Second)
		fmt.Println(i)
		i += 5
	}

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando ", j)
	}

	nomes := []string{"Joao", "Davi", "cleyson"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice2, letra := range "PalaVRA" {
		fmt.Println(indice2, string(letra))
	}

	usario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usario {
		fmt.Println(chave, valor)

	}
}

package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		//Dentro do [] é o tipo das chaves e fora é o tipo dos valores
		"nome":      "Matheus",
		"sobrenome": "Libanio",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	// também conseguimos fazer um map dentro de map
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Matheus",
			"sobrenome": "Libanio",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "campus1",
		},
	}
	fmt.Println(usuario2)

	//deletando chaves

	delete(usuario2, "nome")
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string{
		"nome": "Libra",
	}
	fmt.Println(usuario2)
}

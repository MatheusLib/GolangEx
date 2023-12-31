package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) Salvar() {
	fmt.Printf("Salvando o %s em nosso banco de dados\n", u.nome)
}

func (u *usuario) FazAniversario() {
	u.idade++
}
func (u usuario) MaiorDeIdade() bool {
	return u.idade >= 18
}

func main() {

	usuario1 := usuario{"Matheus", 26}

	usuario1.FazAniversario()
	usuario1.Salvar()
	fmt.Println(usuario1.idade)
	fmt.Println(usuario1.MaiorDeIdade())

}

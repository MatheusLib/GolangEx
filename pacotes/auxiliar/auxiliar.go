package auxiliar

import (
	"fmt"
)

// Escrever registra um mensagem na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	//Dentro do mesmo pacote (Nesse caso pacote auxiliar) consigo utilizar função que começam em minúsculo porém fora do pacote tem que começar com maiúsculo
	escrever2()
}

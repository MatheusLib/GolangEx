package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}

}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6!")
}

func main() {

	fmt.Println(alunoEstaAprovado(6, 8))
	fmt.Println("Pós execução")

}

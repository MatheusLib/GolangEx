package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {

	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func main() {
	fmt.Println("Aula de Funções")
	soma := somar(10, 20)
	fmt.Println(soma)

	soma2 := somar(10, somar(5, 15))

	fmt.Println(soma2)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("texto da função 1")

	var g = func(txt string) string {
		return txt
	}

	resultado := g("texto da função 2")

	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	//Como utilizar uma função que tem 2 retornos porém só preciso de 1 deles?

	novaSoma, _ := calculosMatematicos(83, 25)
	fmt.Println(novaSoma)
}

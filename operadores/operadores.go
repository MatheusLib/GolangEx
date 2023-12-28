package main

import "fmt"

func main() {

	// Aritmeticos
	soma := 1 + 2
	subtracao := 2 - 1
	multiplicacao := 3 * 2
	divisao := 10 / 2
	resto := 10 % 5

	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2

	fmt.Println(soma2)

	//Atribuição

	var texto string = "Txt1"
	texto2 := "txt2"

	fmt.Println(texto, texto2)

	//Fim dos operadores de atribuição

	//Operadores relacionais , >,<, ==, !=, >=, <=

	fmt.Println(1 > 2)

	// Operadores Logicos

	fmt.Println(!(true && false))

	fmt.Println(true && false || true)

	//Operadores Unarios
	numeroA := 10
	numeroA++
	println(numeroA)
	numeroA += 15
	numeroA--
	println(numeroA)
	numeroA *= 3
	numeroA -= 20
	println(numeroA)

	//Operadores Ternarios

	var textoTeste string

	if numeroA > 5 {
		textoTeste = "Maior que 5"
	} else {
		textoTeste = "Menor que 5"
	}

	fmt.Println(textoTeste)

}

package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64
	var numero int16 = 100
	var numero32 int32 = 400000
	fmt.Println(numero, numero32)

	//uint = unsigned int, ou seja não tem sinal

	var numeroA uint32 = 100000
	fmt.Println(numeroA)

	//alias
	// int32 =rune, int32 = byte

	var numeroB rune = 12345
	var numeroC byte = 123
	fmt.Println(numeroB, numeroC)

	// Para número reais existem float32 e float64

	var numeroReal1 float32 = 1234.546

	fmt.Println(numeroReal1)

	numeroReal2 := 1234.656544
	fmt.Println(numeroReal2)

	//Strings

	var str string = "texto"

	// Não existe char em GO

	str2 := "\n\n\ttexto2"

	fmt.Println(str, str2)

	char := 'B'

	fmt.Println(char)

	// FIM de strings

	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	var booleano2 bool = false
	var booleano3 bool

	var erro error = errors.New("Erro interno")
	// erro é o nome da variavel, error é o tipo de variavel e errors é o nome do pacote

	fmt.Println(booleano1, booleano2, booleano3, erro)

}

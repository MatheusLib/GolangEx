package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "var3"
		variavel4 string = "var4"
	)

	variavel5, variavel6 := "var5", "var6"

	const constante1 string = "constante1"

	fmt.Println(variavel3, variavel4, variavel5, variavel6, constante1)
}

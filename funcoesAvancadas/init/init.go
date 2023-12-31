package main

import "fmt"

var n int

func init() {
	fmt.Println("executando a função init")
	n = 10
}

func main() {
	fmt.Println("executando a função main")
	fmt.Println(n)
}

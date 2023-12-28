package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("\tArrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//arrays tem tamanhos fixos, slices tem tamanhos dinâmicos

	slice := []int{10, 11, 21, 13, 154, 16, 17}

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)

	fmt.Println(slice)

	slice2 := array2[1:3] // Índice 1 é Inclusivo e o Índice 2 é exclusivo então nesse caso ele pega os índices 1 e 2

	//Ele serve como um ponteiro então se for alterado no array, altera no slice tbm
	fmt.Println(slice2)

	array2[1] = "Posição alterada"

	fmt.Println(slice2)

	//ARRAYS INTERNOS

	slice3 := make([]float32, 10, 13)
	fmt.Println(len(slice3), cap(slice3))

	slice3 = append(slice3, 18)
	slice3 = append(slice3, 22)
	slice3 = append(slice3, 14)
	slice3 = append(slice3, 67)

	fmt.Println(len(slice3), cap(slice3)) // lenght e capacity
	fmt.Println(slice3)

	slice4 := make([]float32, 5)
	fmt.Println(slice4)

}

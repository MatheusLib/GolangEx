package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Domingo"
	default:
		return "Número Inválido"
	}

}

func diaDaSemana2(numero int) string {
	//	Quando não se quer avaliar a mesma variável

	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
		//Clausula fallthrough faz cair para a proxima condição sem avaliar a condição numero então nesse caso diaDaSemana iria retornar segunda-Feira
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Número Inválido"

	}
	return diaDaSemana
}

func main() {
	fmt.Println("Estrutura de Controle")

	numero := 10
	fmt.Println(numero)

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//Variavel dentro somente do IF
	if outronumero := numero; outronumero > 0 {
		fmt.Println("número é maior que 0")
	}
	//Switch

	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana2(3))
}

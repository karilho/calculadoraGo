package main

import "fmt"
import "calculadoraGo/calculadoraSoma"

func main() {
	fmt.Println("Iniciando a aplicação, escolha uma das opções abaixo: ")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Multiplicação")
	fmt.Println("4 - Divisão")
	fmt.Println("5 - Sair")
	var opcao int
	fmt.Scanln(&opcao)

	switch opcao {
	case 1:
		calculadoraSoma.CalculadoraSoma()
	case 5:
		fmt.Println("Saindo da aplicação")
		break
	}

}
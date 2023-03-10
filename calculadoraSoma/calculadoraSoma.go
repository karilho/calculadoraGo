package calculadoraSoma

import "fmt"

func CalculadoraSoma() int {
	fmt.Println("xamou a função")
	fmt.Println("Digite o primeiro número: ")
	var numero1 int
	var numero2 int

	fmt.Scanln(&numero1)
	fmt.Println(numero1)
	fmt.Println("digite o segundo numero: ")

	fmt.Scanln(&numero2)

	fmt.Println(numero2)
	somaDosNumeros := numero2 + numero1
	fmt.Println("A soma dos números é: ", somaDosNumeros)
	return somaDosNumeros
}

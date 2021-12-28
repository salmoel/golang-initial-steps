package main

import "fmt"

func main() {
	// aritméticos
	soma := 1 + 2
	substracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, substracao, divisao, multiplicacao, restoDaDivisao)

	var numeroone int16 = 10
	var numerotwo int16 = 25
	somatwo := numeroone + numerotwo

	fmt.Println(somatwo)

	// atribuição
	var variavel1 string = "String"
	variavel2 := "String"

	fmt.Println(variavel1, variavel2)

	// operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 != 2)

	// operadores lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// operadores unários
	numero := 10
	numero++
	fmt.Println(numero)

	numeroplus := 10
	numeroplus += 15
	fmt.Println(numeroplus)

	numero--
	numero -= 20

	numero *= 3

	numero /= 10

	numero %= 30

	fmt.Println(numero)

	// operador ternário
	var texto string
	if numero > 5 {
		texto = "Maior que cinco!"
	} else {
		texto = "Menor que cinco!"
	}

	fmt.Println(texto)
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	var numeroint8 int8 = 10
	fmt.Println(numeroint8)

	var numeroint16 int16 = 100
	fmt.Println(numeroint16)

	var numeroint32 int32 = 10000
	fmt.Println(numeroint32)

	var numeroint64 int64 = 1000000
	fmt.Println(numeroint64)

	var numeroint int = 100000000000000
	fmt.Println(numeroint)

	numero := 1000000000000000000
	fmt.Println(numero)

	var numeronegativo uint32 = 100000
	fmt.Println(numeronegativo)

	// alias
	// int32 = rune
	var numeronegativorune rune = 100000
	fmt.Println(numeronegativorune)

	// byte = uint
	var numeronegativobyte = 10000000
	fmt.Println(numeronegativobyte)

	var numeroRealum float32 = 123.3
	fmt.Println(numeroRealum)

	var numeroRealdois float32 = 12333.300
	fmt.Println(numeroRealdois)

	nuemroRealtres := 231601.2
	fmt.Println(nuemroRealtres)

	var str string = "golang, go!"
	fmt.Println(str)

	strdois := "go, golang!"
	fmt.Println(strdois)

	char := 'g'
	fmt.Println(char)

	var textoum string
	fmt.Println(textoum)

	textodois := 15
	fmt.Println(textodois)

	var boolean bool = true
	fmt.Println(boolean)

	var booleanone bool
	fmt.Println(booleanone)

	var erroconvencional error
	fmt.Println(erroconvencional)

	var erro error = errors.New("Aqui bateu um erro, pegou?!")
	fmt.Println(erro)
}

package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello world, pacote!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("salmoel@brodevs.com.br")
	fmt.Println(erro)
}

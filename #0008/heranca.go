package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança!")

	p1 := pessoa{"salmoel", "verneck", 24, 175}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)

	e1nome := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1nome.nome)
}

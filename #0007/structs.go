package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {
	fmt.Println("Arquivo Struct!")

	var u usuario
	u.nome = "salmoel"
	u.idade = 24
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Lá", 8}

	usuario2 := usuario{"salmoel", 24, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "salmoel"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade: 24}
	fmt.Println(usuario4)
}

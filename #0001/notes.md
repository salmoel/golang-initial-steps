mod.go funciona como package.json no NodeJS

pela primeira letra se define se será público ou privado, não é como na POO que definimos private, public, protected, se define pela primeira letra, veja:

func main() {
fmt.Print("Hello world, pacote!")
auxiliar.Escrever()
}

Essa primeira letra indica que é pública essa função que foi criada em outro modulo.

Quando essa função começa com letra minuscula ela só pode ser usada no pacote que ela está.

go build

go install

go run ./main.go

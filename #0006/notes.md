na linguagem golang não posso comparar variaveis de tipos diferentes.
um exemplo, veja:

    var numeroone int16 = 10
    var numerotwo int32 = 25
    soma := numeroone + numerotwo

    fmt.Println(soma)

isso não acontece no golang, devido sua forte tipagem.

operadores relacionais, veja alguns:

    // operadores relacionais
    fmt.Println(1 > 2)
    fmt.Println(1 >= 2)
    fmt.Println(1 == 2) -> comparação
    fmt.Println(1 <= 2)
    fmt.Println(1 > 2)
    fmt.Println(1 != 2) -> diferente

operadores lógicos, veja alguns:

    // operadores lógicos
    verdadeiro, falso := true, false
    fmt.Println(verdadeiro && falso) -> e &&
    fmt.Println(verdadeiro || falso) -> ou ||
    fmt.Println(!verdadeiro) -> ! = negação
    fmt.Println(!falso) -> ! = negação

operadores unários, veja alguns:

    [incrementando]
    numeroplus := 10
    numeroplus += 15
    fmt.Println(numeroplus)

    [decrementando]
    numero--
    numero -= 20

veja o exemplo completa:

    numero := 10
    numero++ -> acrescenta mais 1
    fmt.Println(numero)

    numeroplus := 10
    numeroplus += 15 -> acrescenta mais 15
    fmt.Println(numeroplus)

    numero--
    numero -= 20

    numero *= 3 -> multiplica por 3

    numero /= 10 -> divide por 10

    numero %= 30 -> faz o resto de 30

    fmt.Println(numero)

não existe em golang o operador ternário, porque a premissa em golang, é que exista apenas uma maneira de fazer ou executar cada coisa. uma solução é fazermos if e else, xD!

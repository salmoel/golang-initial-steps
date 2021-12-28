tipos de inteiros que existem no go, veja: int8, int16, int32, int64; cada um deles ocupada um espaço diferentes de bits. e também tem o int sozinho, onde não especifico o tamanho.

para números negativos podemos usar o uint ou rune, no caso do rune ele funciona da mesma maneiro que uint32.

8 = byte (uint8)
32 = rune (uint32)

uint é um int sem sinal
por exemplo de uint, veja:

var numeronegativo uint = -10
fmt.Print(numeronegativo)

nesse exemplo acima não posso passar um sinal como fiz, o uint não permite isso, porque ele é sem sinal.

agora veja neste cenário:

var numeronegativoint int64 = -1000000
fmt.Print(numeronegativoint)

quando é nesse cenário eu conseigo fazer e passar com o sinal de negativo "-".

agora com os números reais é mais fácil de se utlizar porque eles usam apenas float32 e float64

núemros reais sempre devem ser separados com ponto e não vírgula.

no caso das string tipamos ela como string como é feito comumente e não existe char em golang. é importante sempre usarmos aspas duplas para string porque se usarmos aspas simples ele identifica como se fosse um char.

o mais próximo que no golang chegamos do char é usando aspas simples. veja:

char := 'g'
fmt.Println(char)

com isso ele retorna de acordo com tabela ascii

    var erro error
    fmt.Println(erro) => desta forma tratamos erro

Para podermos tratar algo com erro usando uma string como fazemos nos console.log, precisamos importar um pacote chamado errors e fazemos da seguinte forma, veja:

    var erro error = errors.New("Aqui bateu um erro, pegou?!")
    fmt.Println(erro)

funções são uma sequencia de dados que o meu programa irá executar ações.

em uma função golang caso ela tenha retorno eu especifico depois do parentese, veja:

func somar(n1 int8, n2 int8) int8 {
}

em GoLang funções são tipos.

como podemos declarar um variavel e jogar uma função dentro dela?
veja:

    var f = func() {
    	fmt.Println("função f!")
    }

    f()

nesta parte f() estou chamando essa função.

quando fazemos desta forma seguinte, veja:

    var n = func(txt string) string {
    	fmt.Println(txt)
    	return txt
    }

    resultado := n("função n sendo passada tanto com parametros e também com um return!")
    fmt.Println(resultado)

quando fazemos assim ele me retorna dois valores, o do return txt e do println

as funções no golang podem possuir mais de um retorno

caso queiramos apenas um resultado no golang passamos um \_ para que ele retorne apenas um resultado da função. veja o exemplo:

    resultadoSoma, _ := calculosMatematicos(10, 15)
    fmt.Println(resultadoSoma, resultadoSubstracao)

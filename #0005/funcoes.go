package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	substracao := n1 - n2
	return soma, substracao
}

func main() {
	soma := somar(10, 10)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("função f!")
	}

	f()

	var s = func(txt string) {
		fmt.Println(txt)
	}

	s("função s com parametros!")

	var n = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := n("função n sendo passada tanto com parametros e também com um return!")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubstracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubstracao)

	// resultadoSoma, _ := calculosMatematicos(10, 15)
	// fmt.Println(resultadoSoma) // pegando apenas um resultado

}

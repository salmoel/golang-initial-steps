para declararmos variaveis podemos usar o var ou := que também funciona como um representação do tipo que estamos definindo nesta variável.

:= => declaração implicitamente
var => declaração explicita

var (
variavel3 string= "lalala"
variavel4 string= "lalala"
) => desta forma definimos várias variáveis de maneira explicita

variavel5, variavel6 := "la", "la"
fmt.Println(variavel5, variavel6) => desta maneira defino várias variáveis de forma implícita.

uma constante declaro da mesma forma que as variáveis, a diferença é que, na constante não se pode mudar o valor.

variavel5, variavel6 = variavel6, variavel5
fmt.Println(variavel5, variavel6) => desta forma eu consigo inverter os valores das variaveis.

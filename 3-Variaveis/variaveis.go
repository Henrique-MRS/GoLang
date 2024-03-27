package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1" //forma 1 = Variavel do tipo Explicita

	variavel2 := "Variavel 2" //forma 2 = Variavel do tipo Implicito (inferencia de tipos)

	var ( //forma 3 = outra forma de declarar variaveis

		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)
	variavel5, variavel6 := "Variavel 5", "Variavel 6" //forma 4

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)
}

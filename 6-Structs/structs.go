package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {

	fmt.Println("Arquivo Structs") // Uma struct é um tipo de dado fundamental que agrupa “variáveis” (atributos) sob um único nome, similar a como uma classe agrupa propriedades e métodos em linguagens orientadas a objetos como Java ou C#

	var usuario1 usuario // primeira forma de fazer
	usuario1.nome = "Davi"
	usuario1.idade = 21
	fmt.Println(usuario1)

	usuario2 := usuario{"Davi", 21} // Segunda forma de fazer
	fmt.Println(usuario2)
}

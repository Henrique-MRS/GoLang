package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail" //Repositório previamente importado no go.mod
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("henrique.rocha@automapower.com") //utilizando a validação de e-mail, fazendo reportar um erro caso não seja válido
	fmt.Println(erro)
}

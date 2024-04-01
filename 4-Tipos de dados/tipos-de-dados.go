package main

import "fmt"

func main() {
	//
	// INTEIROS:
	//
	var numero int = 1 // Número inteiro no Go, pode ser definido como int8, int16, int32 e int64 (bits), caso você deixe apenas int, ele vai pegar da arquitetura do seu computador.
	fmt.Println(numero)

	var numero2 uint32 = 10 // unsygned int = inteiro sem sinal
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 100
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//
	// REAIS (float32, float64)
	//
	var numeroReal float32 = 123.54
	fmt.Println(numeroReal)

	var numeroReal1 float64 = 123000.54
	fmt.Println(numeroReal1)

	numeroReal2 := 12300.97
	fmt.Println(numeroReal2)

	//
	// STRING
	//
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

}

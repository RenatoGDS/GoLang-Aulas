package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 1000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINTS
	var numero4 byte = 123
	fmt.Println(numero4)
	//FIM NUMEROS INTEIROS

	//NUMEROS REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)
	//FIM NUMEROS REAIS

	//STRINGS
	var str string = "texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)
	//FIM STRINGS

	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}

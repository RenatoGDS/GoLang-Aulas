package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Engenharia", "USP"}
	fmt.Println(estudante1)
	fmt.Println(pessoa1.altura)
}

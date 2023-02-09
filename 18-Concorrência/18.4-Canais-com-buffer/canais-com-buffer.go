package main

import "fmt"

func main() {
	canal := make(chan string, 3)
	canal <- "Olá Mundo"
	canal <- "Programando em go!"
	canal <- "Programando em go novamente"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}

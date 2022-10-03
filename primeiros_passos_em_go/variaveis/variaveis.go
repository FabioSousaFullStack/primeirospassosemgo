package main

import 	"fmt"


func main(){
	nome := "Fabio" 
	versao := 1.1

	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão: ", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Sair dos programas")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da variável comando é",&comando)
	fmt.Println("O comando escolhido foi", comando)

	
}
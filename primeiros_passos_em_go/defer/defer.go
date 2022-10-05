package main

import "fmt"

func main(){
	texto := [15]string{
		"Não te amo mais.",
		"Estarei mentindo dizendo que",
		"Ainda te quero como sempre quis.",
		"Tenho certeza que",
		"Nada foi em vão.",
		"Sinto dentro de mim que",
		"Você não significa nada.",
		"Não poderia dizer jamais que",
		"Alimento um grande amor.",
		"Sinto cada vez mais que",
		"Já te esqueci !",
		"E jamais usarei a frase:",
		"EU TE AMO!",
		"Sinto, mas tenho que dizer a verdade",	
		"É tarde demais..."}

	var escolha int

	fmt.Println("Escolha uma opção para ler um dos poemas mais famosos de Claric1e Lispector")
	fmt.Println("1. Sem uso da declaração defer")
	fmt.Println("2. Com uso declaração defer")
	fmt.Print("Digite sua opção: ")
	fmt.Scan(&escolha)


	
	if escolha == 2{
		for i := 0; i < 15 ; i++ {
			defer fmt.Println(texto[i])
		}}else{
			for i := 0; i < 15 ; i++ {
			fmt.Println(texto[i])
		}} 

		fmt.Println("")
		fmt.Println("********************************")
		fmt.Println("Poema de Clarice Lispector")
		fmt.Println("Não te amo mais")
		fmt.Println("*********************************")
		fmt.Println("")
}
































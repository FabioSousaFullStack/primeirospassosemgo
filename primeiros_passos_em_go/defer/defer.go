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

	for i := 0; i < 15 ; i++ {
		defer fmt.Println(texto[i])
	}
	 
}
































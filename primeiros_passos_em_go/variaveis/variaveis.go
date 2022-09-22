package main

import 	"fmt"
import	"reflect"

func main(){
	nome := "Fabio" 
	versao := 1.1

	fmt.Println("Olá sr.", nome)
	fmt.Println("Versão", versao)

	fmt.Println("O tipo da variável nome é :", reflect.TypeOf(nome))


	
}
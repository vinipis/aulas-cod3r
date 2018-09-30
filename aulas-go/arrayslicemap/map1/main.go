package main

import (
	"fmt"
)

func main() {
	//var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678971] = "Pedro"
	aprovados[12345678972] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678972])
	delete(aprovados, 12345678972)
	fmt.Println(aprovados[12345678972])
}

package main

import (
	"fmt"
)

func inc1(n int) {
	n++
}

//Revisão: um ponteiro é representado por um *
func inc2(n *int) {
	//Revisão: * é usado para dexreferenciar, ou seja,
	//ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	//revisão: & usado para obter o endereço da variavel
	inc2(&n) // por referencia
	fmt.Println(n)
}

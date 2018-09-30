package main

import (
	"fmt"
)

func main() {
	funcsESalarios := map[string]float64{
		"José João":       11325.45,
		"Gabriela Santos": 15456.78,
		"Pedro Junior":    1200.0,
	}
	funcsESalarios["Rafael filho"] = 1340.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}

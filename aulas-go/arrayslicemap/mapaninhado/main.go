package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabrliela Santos": 15456.78,
			"Gabriel Silva":    1567.9,
		},
		"J": {
			"José": 11325.43,
			"João": 23413.54,
		},
		"P": {
			"Pedro Junior": 1200.9,
		},
	}
	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}

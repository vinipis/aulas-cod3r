package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"Nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct para json
	p1 := produto{1, "Notbook", 1899.00, []string{"Promoção", "Eletronico"}}
	p1JSON, _ := json.Marshal(p1)
	fmt.Println(string(p1JSON))

	//json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}

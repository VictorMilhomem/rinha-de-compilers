package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/VictorMilhomem/rinha-de-compilers/cmd/interpreter"
)

func main() {
	data, err := ioutil.ReadFile("examples\\hello.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Decodificar o JSON para a estrutura File
	var file interpreter.File
	if err := json.Unmarshal(data, &file); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	fmt.Printf("%v", file)
}

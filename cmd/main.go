package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/VictorMilhomem/rinha-de-compilers/cmd/interpreter"
)

func main() {
	data, err := ioutil.ReadFile("examples\\sumNumStr.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	var file interpreter.File
	if err := json.Unmarshal(data, &file); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}
	interpreter.Eval(file.Expression)
}

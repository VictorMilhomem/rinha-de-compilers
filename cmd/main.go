package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/VictorMilhomem/rinha-de-compilers/cmd/interpreter"
)

var env = interpreter.NewEnvironment()

func main() {
	data, err := ioutil.ReadFile("examples\\comb.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	var file interpreter.Ast
	if err := json.Unmarshal(data, &file); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	interpreter.Eval(file.Expression, env)
}

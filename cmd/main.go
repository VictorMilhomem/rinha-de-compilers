package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/VictorMilhomem/rinha-de-compilers/cmd/interpreter"
)

var env = interpreter.NewEnvironment()

func main() {
	jsonFilePath := flag.String("json", "", "Path to the JSON file")
	flag.Parse()

	if *jsonFilePath == "" {
		fmt.Println("Usage: programName -json <jsonFilePath>")
		return
	}

	data, err := ioutil.ReadFile(*jsonFilePath)
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

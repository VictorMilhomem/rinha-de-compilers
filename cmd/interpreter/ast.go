package interpreter

import "fmt"

type Ast struct {
	Expression interface{} `json:"expression"`
}

func node(ast interface{}, key string) interface{} {
	return ast.(map[string]interface{})[key]
}

type Tuple struct {
	first, second interface{}
}

func (t *Tuple) printTuple() {
	fmt.Print("(", t.first, ",", t.second, ")")
}

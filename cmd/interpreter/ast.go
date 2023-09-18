package interpreter

import "fmt"

type Ast struct {
	Expression interface{} `json:"expression"`
}

func node(ast interface{}, key string) interface{} {
	return ast.(map[string]interface{})[key]
}

type Closure struct {
	parameters []interface{}
	body       map[string]interface{}
	location   interface{}
}

type Tuple struct {
	first, second interface{}
}

func (t *Tuple) printTuple(indent int) {
	fmt.Print("(")
	printValue(t.first, indent)
	fmt.Print(",")
	printValue(t.second, indent)
	fmt.Print(")")
}

func printValue(value interface{}, indent int) {
	switch v := value.(type) {
	case Tuple:
		v.printTuple(indent + 1)
	default:
		fmt.Print(v)
	}
}

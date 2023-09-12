package interpreter

import "fmt"

// Simple Tree-Walk Interpreter

func evalPrint(expression map[string]interface{}) {
	switch expression["kind"].(string) {
	case "Str":
		fmt.Printf("%s", expression["value"].(string))
	case "Int":
		fmt.Printf("%d", int(expression["value"].(float64)))
	case "Bool":
		fmt.Printf("%v", expression["value"].(bool))
	}
}

func Eval(expression Term) interface{} {
	expressionKind, ok := expression.(map[string]interface{})
	if !ok {
		fmt.Println("Error parsing expression kind")
		return nil
	}

	switch expressionKind["kind"].(string) {
	case "Print":
		evalPrint(expressionKind["value"].(map[string]interface{}))
		return nil
	case "Str":
		return NewStr(expressionKind["value"].(string), expressionKind["location"].(Location))
	case "Int":
		return NewInt(expressionKind["value"].(float64), expressionKind["location"].(Location))
	case "Bool":
		return NewBool(expressionKind["value"].(bool), expressionKind["location"].(Location))
	default:
		return nil
	}
}

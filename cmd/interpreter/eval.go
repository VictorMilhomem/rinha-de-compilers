package interpreter

import "fmt"

func Eval(expression Term) interface{} {
	expressionKind, ok := expression.(map[string]interface{})
	if !ok {
		fmt.Println("Error parsing expression kind")
		return nil
	}

	switch expressionKind["kind"].(string) {
	case "Print":
		printValue := expressionKind["value"].(map[string]interface{})
		switch printValue["kind"].(string) {
		case "Str":
			fmt.Printf("%s", printValue["value"].(string))
		}
		return nil
	case "Str":
		return NewStr("Str", expressionKind["value"].(string), expressionKind["location"].(Location))
	default:
		return nil
	}
}

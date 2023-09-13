package interpreter

import "fmt"

func Eval(node Expression) interface{} {
	switch node.Kind {
	case "Int":
		return node.Value.(float64)
	case "Str":
		return node.Value.(string)
	case "Print":
		kind := node.Value.(map[string]interface{})["kind"]
		var printValue interface{}
		if kind == "Binary" {
			printValue = evaluateBinary(node.Value.(map[string]interface{}))
		} else {
			printValue = node.Value.(map[string]interface{})["value"]
		}
		fmt.Printf("%v", printValue)
		return nil
	default:
		panic("Unsupported expression kind: " + node.Kind)
	}
}

func evaluateBinary(binaryNode map[string]interface{}) interface{} {
	lhsExpr, lhsExists := binaryNode["lhs"]
	rhsExpr, rhsExists := binaryNode["rhs"]
	op, opExists := binaryNode["op"]

	if !lhsExists || !rhsExists || !opExists {
		panic("Invalid binary expression structure")
	}

	lhs := Eval(Expression{
		Kind:     lhsExpr.(map[string]interface{})["kind"].(string),
		Value:    lhsExpr.(map[string]interface{})["value"],
		Location: parseLocation(lhsExpr.(map[string]interface{})["location"].(map[string]interface{})),
	})
	rhs := Eval(Expression{
		Kind:     rhsExpr.(map[string]interface{})["kind"].(string),
		Value:    rhsExpr.(map[string]interface{})["value"],
		Location: parseLocation(rhsExpr.(map[string]interface{})["location"].(map[string]interface{})),
	})
	operator := op.(string)

	switch operator {
	case "Add":
		// TODO: Sum number and string
		return lhs.(float64) + rhs.(float64)
	case "Sub":
		return lhs.(float64) - rhs.(float64)
	case "Mul":
		return lhs.(float64) * rhs.(float64)
	case "Div":
		return lhs.(float64) / rhs.(float64)
	default:
		panic("Unsupported binary operation: " + operator)
	}
}

func parseLocation(locationData map[string]interface{}) Location {
	return Location{
		Start:    int(locationData["start"].(float64)),
		End:      int(locationData["end"].(float64)),
		Filename: locationData["filename"].(string),
	}
}

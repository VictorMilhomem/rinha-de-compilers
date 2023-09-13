package interpreter

import (
	"fmt"
	"strings"
)

func Eval(node Expression) interface{} {
	switch node.Kind {
	case "If":
		return evalueteIf(node.Value.(map[string]interface{}))
	case "Int":
		return node.Value.(float64)
	case "Str":
		return node.Value.(string)
	case "Bool":
		return node.Value.(bool)
	case "Binary":
		return evaluateBinary(node.Value.(map[string]interface{}))
	case "Print":
		kind := node.Value.(map[string]interface{})["kind"]
		var printValue interface{}
		switch kind {
		case "Binary":
			printValue = evaluateBinary(node.Value.(map[string]interface{}))
		case "If":
			printValue = evalueteIf(node.Value.(map[string]interface{}))
		default:
			printValue = node.Value.(map[string]interface{})["value"]
		}

		fmt.Printf("%v", printValue)
		return nil
	default:
		panic("Unsupported expression kind: " + node.Kind)
	}
}

func evalueteIf(ifNode map[string]interface{}) interface{} {
	condition, conditionExists := ifNode["condition"]
	then, theExistis := ifNode["then"]
	otherwise, otherwiseExists := ifNode["otherwise"]

	if !conditionExists || !theExistis || !otherwiseExists {
		panic("Invalid if expression structure")
	}
	switch val := Eval(Expression{
		Kind:     condition.(map[string]interface{})["kind"].(string),
		Value:    condition.(map[string]interface{})["value"],
		Location: parseLocation(condition.(map[string]interface{})["location"].(map[string]interface{})),
	}); val {
	case true:
		return Eval(Expression{
			Kind:     then.(map[string]interface{})["kind"].(string),
			Value:    then.(map[string]interface{})["value"],
			Location: parseLocation(then.(map[string]interface{})["location"].(map[string]interface{})),
		})
	case false:
		return Eval(Expression{
			Kind:     otherwise.(map[string]interface{})["kind"].(string),
			Value:    otherwise.(map[string]interface{})["value"],
			Location: parseLocation(otherwise.(map[string]interface{})["location"].(map[string]interface{})),
		})
	default:
		panic("Error in if expression")
	}
}

func evaluateBinary(binaryNode map[string]interface{}) interface{} {
	lhsExpr, lhsExists := binaryNode["lhs"]
	rhsExpr, rhsExists := binaryNode["rhs"]
	op, opExists := binaryNode["op"]

	if !lhsExists || !rhsExists || !opExists {
		panic("Invalid binary expression structure")
	}
	lkind := lhsExpr.(map[string]interface{})["kind"].(string)
	var lhs interface{}
	switch lkind {
	case "Binary":
		lhs = evaluateBinary(lhsExpr.(map[string]interface{}))
	default:
		lhs = Eval(Expression{
			Kind:     lhsExpr.(map[string]interface{})["kind"].(string),
			Value:    lhsExpr.(map[string]interface{})["value"],
			Location: parseLocation(lhsExpr.(map[string]interface{})["location"].(map[string]interface{})),
		})
	}

	rkind := rhsExpr.(map[string]interface{})["kind"].(string)
	var rhs interface{}
	switch rkind {
	case "Binary":
		rhs = evaluateBinary(rhsExpr.(map[string]interface{}))
	default:
		rhs = Eval(Expression{
			Kind:     rhsExpr.(map[string]interface{})["kind"].(string),
			Value:    rhsExpr.(map[string]interface{})["value"],
			Location: parseLocation(rhsExpr.(map[string]interface{})["location"].(map[string]interface{})),
		})
	}

	operator := op.(string)
	// Equality on string operations
	switch operator {
	case "Add":
		switch lhs.(type) {
		case string:
			lnode, _ := lhs.(string)
			switch rhs.(type) {
			case string:
				rnode, _ := rhs.(string)
				return lnode + rnode
			case float64:
				rnode := fmt.Sprintf("%v", rhs.(float64))
				return lnode + rnode
			default:
				panic("Error: could not add")
			}
		case float64:
			lnode, _ := lhs.(float64)
			switch rhs.(type) {
			case string:
				lnode := fmt.Sprintf("%v", lhs.(float64))
				rnode, _ := rhs.(string)
				return lnode + rnode
			case float64:
				rnode := rhs.(float64)
				return lnode + rnode
			default:
				panic("Error: could not add")
			}
		default:
			panic("Error: could not add")
		}
	case "Sub":
		return lhs.(float64) - rhs.(float64)
	case "Mul":
		return lhs.(float64) * rhs.(float64)
	case "Div":
		return lhs.(float64) / rhs.(float64)
	case "Rem":
		return lhs.(float64) / rhs.(float64)
	case "Eq":
		switch lhs.(type) {
		case string:
			lnode, _ := lhs.(string)
			rnode, okr := rhs.(string)
			if okr {
				return strings.Compare(lnode, rnode)
			}
		case float64:
			lnode, _ := lhs.(float64)
			rnode, okr := rhs.(float64)
			if okr {
				return lnode == rnode
			}
		case bool:
			lnode, _ := lhs.(bool)
			rnode, okr := rhs.(bool)
			if okr {
				return lnode == rnode
			}
		}
		// TODO: Create Error Type
		panic("Unsupported binary operation: " + operator)
	case "Neq":
		switch lhs.(type) {
		case string:
			lnode, _ := lhs.(string)
			rnode, okr := rhs.(string)
			if okr {
				return strings.Compare(lnode, rnode)
			}
		case float64:
			lnode, _ := lhs.(float64)
			rnode, okr := rhs.(float64)
			if okr {
				return lnode != rnode
			}
		case bool:
			lnode, _ := lhs.(bool)
			rnode, okr := rhs.(bool)
			if okr {
				return lnode != rnode
			}
		}
		// TODO: Create Error Type
		panic("Unsupported binary operation: " + operator)
	case "Lt":
		return lhs.(float64) < rhs.(float64)
	case "Gt":
		return lhs.(float64) > rhs.(float64)
	case "Lte":
		return lhs.(float64) <= rhs.(float64)
	case "Gte":
		return lhs.(float64) >= rhs.(float64)
	case "And":
		return lhs.(bool) && rhs.(bool)
	case "Or":
		return lhs.(bool) || rhs.(bool)
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

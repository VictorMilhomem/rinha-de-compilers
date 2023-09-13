package interpreter

import (
	"fmt"
	"strings"
)

func Eval(node Expression) interface{} {
	switch node.Kind {
	case "Int":
		return node.Value.(float64)
	case "Str":
		return node.Value.(string)
	case "Binary":
		return evaluateBinary(node.Value.(map[string]interface{}))
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
		// TODO: Sum number and string
		return lhs.(float64) + rhs.(float64)
	case "Sub":
		return lhs.(float64) - rhs.(float64)
	case "Mul":
		return lhs.(float64) * rhs.(float64)
	case "Div":
		return lhs.(float64) / rhs.(float64)
	case "Rem":
		return lhs.(float64) / rhs.(float64)
	case "Eq":
		lnode, okl := lhs.(string)
		rnode, okr := rhs.(string)
		if okl && okr {
			return strings.Compare(lnode, rnode)
		}

		l, okl := lhs.(float64)
		r, okr := rhs.(float64)
		if okl && okr {
			return l == r
		}

		lb, okl := lhs.(bool)
		rb, okr := rhs.(bool)
		if okl && okr {
			return lb == rb
		}
		// TODO: Create Error Type
		panic("Unsupported binary operation: " + operator)
	case "Neq":
		lnode, okl := lhs.(string)
		rnode, okr := rhs.(string)
		if okl && okr {
			return strings.Compare(lnode, rnode)
		}

		l, okl := lhs.(float64)
		r, okr := rhs.(float64)
		if okl && okr {
			return l != r
		}

		lb, okl := lhs.(bool)
		rb, okr := rhs.(bool)
		if okl && okr {
			return lb != rb
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

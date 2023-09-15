package interpreter

import (
	"fmt"
	"strings"
)

type Environment struct {
	Values map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{Values: make(map[string]interface{})}
}

func (e *Environment) Get(name string) (interface{}, bool) {
	val, ok := e.Values[name]
	return val, ok
}

func (e *Environment) Set(name string, val interface{}) interface{} {
	e.Values[name] = val
	return val
}

// TODO: Tail Call Optimization
func Eval(node Expression, env *Environment) interface{} {
	switch node.Kind {
	case "Let":
		nameNode := node.Let.(map[string]interface{})["text"].(string)
		var val interface{}
		switch kind := node.Value.(map[string]interface{})["kind"].(string); kind {
		case "Binary":
			val = evaluateBinary(node.Value.(map[string]interface{}), env)
		case "Tuple":
			val = evaluateTuple(node.Value.(map[string]interface{}), env)
		default:
			val = Eval(Expression{
				Kind:     kind,
				Value:    node.Value.(map[string]interface{})["value"],
				Location: parseLocation(node.Value.(map[string]interface{})["location"].(map[string]interface{})),
			}, env)
		}
		env.Set(nameNode, val)
		nextKind := node.Next.(map[string]interface{})["kind"].(string)
		switch nextKind {
		case "Binary":
			return evaluateBinary(node.Next.(map[string]interface{}), env)
		case "Tuple":
			return evaluateTuple(node.Next.(map[string]interface{}), env)
		case "Let":
			return Eval(Expression{
				Kind:     node.Next.(map[string]interface{})["kind"].(string),
				Let:      node.Next.(map[string]interface{})["name"],
				Next:     node.Next.(map[string]interface{})["next"],
				Value:    node.Next.(map[string]interface{})["value"],
				Location: parseLocation(node.Next.(map[string]interface{})["location"].(map[string]interface{})),
			}, env)
		default:
			return Eval(Expression{
				Kind:     node.Next.(map[string]interface{})["kind"].(string),
				Value:    node.Next.(map[string]interface{})["value"],
				Location: parseLocation(node.Next.(map[string]interface{})["location"].(map[string]interface{})),
			}, env)
		}
	case "Tuple":
		return evaluateTuple(node.Value.(map[string]interface{}), env)
	case "Var":
		return evaluateVar(node.Value.(map[string]interface{}), env)
	case "Function":
		return nil
	case "If":
		return evaluateIf(node.Value.(map[string]interface{}), env)
	case "Int":
		return node.Value.(float64)
	case "Str":
		return node.Value.(string)
	case "Bool":
		return node.Value.(bool)
	case "Binary":
		return evaluateBinary(node.Value.(map[string]interface{}), env)
	case "Print":
		kind := node.Value.(map[string]interface{})["kind"]
		var printValue interface{}
		switch kind {
		case "Binary":
			printValue = evaluateBinary(node.Value.(map[string]interface{}), env)
		case "If":
			printValue = evaluateIf(node.Value.(map[string]interface{}), env)
		case "Var":
			printValue = evaluateVar(node.Value.(map[string]interface{}), env)
		case "Tuple":
			printValue = evaluateTuple(node.Value.(map[string]interface{}), env)
			// fmt.Printf("(%v, %v)", printValue.(Tuple).first, printValue.(Tuple).second)
			// return nil
		default:
			printValue = node.Value.(map[string]interface{})["value"]
		}

		fmt.Printf("%v", printValue)
		return nil
	default:
		panic("Unsupported expression kind: " + node.Kind)
	}
}

func evaluateTuple(tupleNode map[string]interface{}, env *Environment) interface{} {
	firstExpr := tupleNode["first"]
	secondExpr := tupleNode["second"]
	var first interface{}
	switch kind := firstExpr.(map[string]interface{})["kind"].(string); kind {
	case "Binary":
		first = evaluateBinary(firstExpr.(map[string]interface{}), env)
	case "Var":
		first = evaluateVar(firstExpr.(map[string]interface{}), env)
	default:
		first = Eval(Expression{
			Kind:     kind,
			Value:    firstExpr.(map[string]interface{})["value"],
			Location: parseLocation(firstExpr.(map[string]interface{})["location"].(map[string]interface{})),
		}, env)
	}

	var second interface{}
	switch kind := secondExpr.(map[string]interface{})["kind"].(string); kind {
	case "Binary":
		second = evaluateBinary(secondExpr.(map[string]interface{}), env)
	case "Var":
		second = evaluateVar(secondExpr.(map[string]interface{}), env)
	default:
		second = Eval(Expression{
			Kind:     secondExpr.(map[string]interface{})["kind"].(string),
			Value:    secondExpr.(map[string]interface{})["value"],
			Location: parseLocation(secondExpr.(map[string]interface{})["location"].(map[string]interface{})),
		}, env)
	}

	return Tuple{first: first, second: second}
}

func evaluateVar(varNode map[string]interface{}, env *Environment) interface{} {
	text := varNode["text"].(string)

	val, ok := env.Get(text)
	if !ok {
		panic("could not get env value")
	}
	return val
}

func evaluateIf(ifNode map[string]interface{}, env *Environment) interface{} {
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
	}, env); val {
	case true:
		return Eval(Expression{
			Kind:     then.(map[string]interface{})["kind"].(string),
			Value:    then.(map[string]interface{})["value"],
			Location: parseLocation(then.(map[string]interface{})["location"].(map[string]interface{})),
		}, env)
	case false:
		return Eval(Expression{
			Kind:     otherwise.(map[string]interface{})["kind"].(string),
			Value:    otherwise.(map[string]interface{})["value"],
			Location: parseLocation(otherwise.(map[string]interface{})["location"].(map[string]interface{})),
		}, env)
	default:
		panic("Error in if expression")
	}
}

func evaluateBinary(binaryNode map[string]interface{}, env *Environment) interface{} {
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
		lhs = evaluateBinary(lhsExpr.(map[string]interface{}), env)
	case "Var":
		lhs = evaluateVar(lhsExpr.(map[string]interface{}), env)
	default:
		lhs = Eval(Expression{
			Kind:     lhsExpr.(map[string]interface{})["kind"].(string),
			Value:    lhsExpr.(map[string]interface{})["value"],
			Location: parseLocation(lhsExpr.(map[string]interface{})["location"].(map[string]interface{})),
		}, env)
	}

	rkind := rhsExpr.(map[string]interface{})["kind"].(string)
	var rhs interface{}
	switch rkind {
	case "Binary":
		rhs = evaluateBinary(rhsExpr.(map[string]interface{}), env)
	case "Var":
		rhs = evaluateVar(rhsExpr.(map[string]interface{}), env)
	default:
		rhs = Eval(Expression{
			Kind:     rhsExpr.(map[string]interface{})["kind"].(string),
			Value:    rhsExpr.(map[string]interface{})["value"],
			Location: parseLocation(rhsExpr.(map[string]interface{})["location"].(map[string]interface{})),
		}, env)
	}

	operator := op.(string)
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

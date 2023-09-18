package interpreter

import (
	"fmt"
	"math"
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

func CopyEnv(env *Environment) *Environment {
	newEnv := NewEnvironment()
	for k, v := range env.Values {
		newEnv.Values[k] = v
	}
	return newEnv
}

func Eval(expression interface{}, env *Environment) interface{} {
	switch kind := node(expression, "kind"); kind {
	case "Str":
		return node(expression, "value").(string)
	case "Int":
		return node(expression, "value").(float64)
	case "Bool":
		return node(expression, "value").(bool)
	case "Function":
		return evalFunction(expression, env)
	case "Call":
		return evalCall(expression, env)
	case "Tuple":
		return evalTuple(expression, env)
	case "First", "Second":
		return getTupleElement(expression, env, kind.(string))
	case "Let":
		return evalLet(expression, env)
	case "Var":
		return evalVar(expression, env)
	case "Binary":
		return evalBinary(expression, env)
	case "If":
		return evalIf(expression, env)
	case "Print":
		val := Eval(node(expression, "value"), env)
		if val, ok := val.(Tuple); ok {
			val.printTuple(0)
			return nil
		}

		fmt.Printf("%v", val)

	}
	return nil
}

func checkFloat(expression interface{}) bool {
	switch expression.(type) {
	case float64:
		return true
	default:
		return false
	}
}

func checkStr(expression interface{}) bool {
	switch expression.(type) {
	case string:
		return true
	default:
		return false
	}
}

func CompareStr(lhs, rhs string) bool {
	switch res := strings.Compare(lhs, rhs); res {
	case 0:
		return true
	default:
		return false
	}
}

func evalFunction(expression interface{}, env *Environment) interface{} {
	params := node(expression, "parameters").([]interface{})
	body := node(expression, "value").(map[string]interface{})
	loc := node(expression, "location")
	return Closure{
		parameters: params,
		body:       body,
		location:   loc,
	}
}

func evalCall(expression interface{}, env *Environment) interface{} {
	callee := node(expression, "callee")
	args := node(expression, "arguments").([]interface{})
	fn := Eval(callee, env)

	fnScope := CopyEnv(env)
	for i, param := range fn.(Closure).parameters {
		val := Eval(args[i], env)
		fnScope.Set(node(param, "text").(string), val)

	}

	body := fn.(Closure).body
	return Eval(body, fnScope)
}

func getTupleElement(expression interface{}, env *Environment, kind string) interface{} {
	val := Eval(node(expression, "value"), env)
	switch kind {
	case "First":
		return val.(Tuple).first
	case "Second":
		return val.(Tuple).second
	default:
		panic("could not get tuple value")
	}
}

func evalTuple(expression interface{}, env *Environment) interface{} {
	first := Eval(node(expression, "first"), env)
	second := Eval(node(expression, "second"), env)
	return Tuple{first: first, second: second}
}

func evalVar(expression interface{}, env *Environment) interface{} {
	text := node(expression, "text")
	val, ok := env.Get(text.(string))
	if !ok {
		panic("could not get variable")
	}

	return val
}

func evalLet(expression interface{}, env *Environment) interface{} {
	nameNode := node(expression, "name")
	text := node(nameNode, "text")

	val := Eval(node(expression, "value"), env)

	env.Set(text.(string), val)
	return Eval(node(expression, "next"), env)
}

func evalIf(expression interface{}, env *Environment) interface{} {
	condition := node(expression, "condition")
	then := node(expression, "then")
	otherwise := node(expression, "otherwise")

	switch val := Eval(condition, env); val {
	case true:
		return Eval(then, env)
	case false:
		return Eval(otherwise, env)
	default:
		panic("Error in if expression")
	}
}

func evalBinary(expression interface{}, env *Environment) interface{} {
	lhs := Eval(node(expression, "lhs"), env)
	rhs := Eval(node(expression, "rhs"), env)

	switch op := node(expression, "op"); op {
	case "Add":
		switch {
		case checkFloat(lhs) && checkFloat(rhs):
			return lhs.(float64) + rhs.(float64)
		case checkFloat(lhs) && checkStr(rhs):
			return fmt.Sprintf("%v", lhs.(float64)) + rhs.(string)
		case checkStr(lhs) && checkFloat(rhs):
			return lhs.(string) + fmt.Sprintf("%v", rhs.(float64))
		case checkStr(lhs) && checkStr(rhs):
			return lhs.(string) + rhs.(string)
		default:
			panic("Unsupported operation")
		}

	case "Sub":
		return lhs.(float64) - rhs.(float64)
	case "Mul":
		return lhs.(float64) * rhs.(float64)
	case "Div":
		return lhs.(float64) / rhs.(float64)
	case "Rem":
		lhs, _ := lhs.(float64)
		rhs, _ := rhs.(float64)
		return math.Mod(lhs, rhs)
	case "Eq":
		switch {
		case checkFloat(lhs) && checkFloat(rhs):
			return lhs.(float64) == rhs.(float64)
		case checkStr(lhs) && checkStr(rhs):
			return CompareStr(lhs.(string), rhs.(string))
		default:
			panic("Unsupported operation")
		}
	case "Neq":
		switch {
		case checkFloat(lhs) && checkFloat(rhs):
			return lhs.(float64) != rhs.(float64)
		case checkStr(lhs) && checkStr(rhs):
			return !CompareStr(lhs.(string), rhs.(string))
		default:
			panic("Unsupported operation")
		}
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
		panic("Unsupported binary operation: " + op.(string))
	}
}

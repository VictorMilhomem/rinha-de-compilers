package interpreter

type Ast struct {
	Expression interface{} `json:"expression"`
}

func node(ast interface{}, key string) interface{} {
	return ast.(map[string]interface{})[key]
}

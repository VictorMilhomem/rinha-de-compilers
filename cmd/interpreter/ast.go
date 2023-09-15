package interpreter

type File struct {
	Name       string     `json:"name"`
	Expression Expression `json:"expression"`
	Location   Location   `json:"location"`
}

type Expression struct {
	Kind     string      `json:"kind"`
	Let      interface{} `json:"name"`
	Next     interface{} `json:"next"`
	Value    interface{} `json:"value"`
	Location Location    `json:"location"`
}

type Closure struct {
	body   interface{}
	params []interface{}
	env    *Environment
}

type Tuple struct {
	first, second interface{}
}

type Location struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Filename string `json:"filename"`
}

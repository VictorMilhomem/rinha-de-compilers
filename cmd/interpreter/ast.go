package interpreter

type File struct {
	Name       string     `json:"name"`
	Expression Expression `json:"expression"`
	Location   Location   `json:"location"`
}

type Expression struct {
	Kind     string      `json:"kind"`
	Value    interface{} `json:"value"`
	Location Location    `json:"location"`
}

type Location struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Filename string `json:"filename"`
}

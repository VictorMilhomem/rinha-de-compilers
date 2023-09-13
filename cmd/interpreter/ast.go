package interpreter

type File struct {
	Name       string     `json:"name"`
	Expression Expression `json:"expression"`
	Location   Location   `json:"location"`
}

type Expression struct {
	Kind     string      `json:"kind"`
	Let      Let         `json:"name"`
	Next     interface{} `json:"next"`
	Value    interface{} `json:"value"`
	Location Location    `json:"location"`
}

type Let struct {
	Name Parameter `json:"name"`
}

type Parameter struct {
	Text     string   `json:"text"`
	Location Location `json:"location"`
}

type Location struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Filename string `json:"filename"`
}

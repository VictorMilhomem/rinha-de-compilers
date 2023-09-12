package interpreter

type File struct {
	name       string   `json:"name"`
	expression Term     `json:"expression"`
	location   Location `json:"location"`
}

type Location struct {
	start    int    `json:"start"`
	end      int    `json:"end"`
	filename string `json:"filename"`
}

type Parameter struct {
	text     string   `json:"text"`
	location Location `json:"location"`
}

type Var struct {
	kind     string   `json:"kind"`
	text     string   `json:"text"`
	location Location `json:"location"`
}

type Function struct {
	kind       string      `json:"kind"`
	parameters []Parameter `json:"parameters"`
	value      Term        `json:"value"`
	location   Location    `json:"location"`
}

type Call struct {
	kind      string   `json:"kind"`
	callee    Term     `json:"callee"`
	arguments Term     `json:"arguments"`
	location  Location `json:"location"`
}
type Let struct {
	kind     string    `json:"kind"`
	name     Parameter `json:"parameter"`
	value    Term      `json:"value"`
	next     Term      `json:"next"`
	location Location  `json:"location"`
}

type Str struct {
	kind     string   `json:"kind"`
	value    string   `json:"value"`
	location Location `json:"location"`
}
type Int struct {
	kind     string   `json:"kind"`
	value    float64  `json:"value"`
	location Location `json:"location"`
}
type BinaryOp struct{}

type Bool struct {
	kind     string   `json:"kind"`
	value    bool     `json:"value"`
	location Location `json:"location"`
}
type If struct {
	kind      string   `json:"kind"`
	condition Term     `json:"condition"`
	then      Term     `json:"then"`
	otherwise Term     `json:"otherwise"`
	location  Location `json:"location"`
}

type Binary struct {
	kind     string   `json:"kind"`
	lhs      Term     `json:"lhs"`
	op       BinaryOp `json:"op"`
	rhs      Term     `json:"rhs"`
	location Location `json:"location"`
}
type Tuple struct {
	kind     string   `json:"kind"`
	first    Term     `json:"first"`
	second   Term     `json:"second"`
	location Location `json:"location"`
}
type First struct {
	kind     string   `json:"kind"`
	value    bool     `json:"value"`
	location Location `json:"location"`
}
type Second struct {
	kind     string   `json:"kind"`
	value    bool     `json:"value"`
	location Location `json:"location"`
}
type Print struct {
	kind     string   `json:"kind"`
	value    bool     `json:"value"`
	location Location `json:"location"`
}

type Term interface{}

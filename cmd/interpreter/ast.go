package interpreter

type File struct {
	Name       string   `json:"name"`
	Expression Term     `json:"expression"`
	Location   Location `json:"location"`
}

type Location struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Filename string `json:"filename"`
}

type Parameter struct {
	Text     string   `json:"text"`
	Location Location `json:"location"`
}

type Var struct {
	Kind     string   `json:"kind"`
	Text     string   `json:"text"`
	Location Location `json:"location"`
}

type Function struct {
	Kind       string      `json:"kind"`
	Parameters []Parameter `json:"parameters"`
	Value      Term        `json:"value"`
	Location   Location    `json:"location"`
}

type Call struct {
	Kind      string   `json:"kind"`
	Callee    Term     `json:"callee"`
	Arguments Term     `json:"arguments"`
	Location  Location `json:"location"`
}
type Let struct {
	Kind     string    `json:"kind"`
	Name     Parameter `json:"parameter"`
	Value    Term      `json:"value"`
	Next     Term      `json:"next"`
	Location Location  `json:"location"`
}

type Str struct {
	Kind     string   `json:"kind"`
	Value    string   `json:"value"`
	Location Location `json:"location"`
}

func NewStr(value string, loc Location) Str {
	return Str{
		Kind:     "Print",
		Value:    value,
		Location: loc,
	}
}

type Int struct {
	Kind     string   `json:"kind"`
	Value    float64  `json:"value"`
	Location Location `json:"location"`
}

func NewInt(value float64, loc Location) Int {
	return Int{
		Kind:     "Int",
		Value:    value,
		Location: loc,
	}
}

type OpType int

const (
	Add OpType = iota
	Sub
	Mul
	Div
	Rem
	Eq
	Neq
	Lt
	Gt
	Lte
	Gte
	And
	Or
)

type Op struct {
	Name string
	Op   OpType
}

type BinaryOp struct {
	Op string
}

func NewBinaryOp(optype Op) BinaryOp {
	switch optype.Op {
	case Add:
		return BinaryOp{Op: "Add"}
	case Sub:
		return BinaryOp{Op: "Sub"}
	case Mul:
		return BinaryOp{Op: "Mul"}
	case Div:
		return BinaryOp{Op: "Div"}
	case Rem:
		return BinaryOp{Op: "Rem"}
	case Eq:
		return BinaryOp{Op: "Eq"}
	case Neq:
		return BinaryOp{Op: "Neq"}
	case Lt:
		return BinaryOp{Op: "Lt"}
	case Gt:
		return BinaryOp{Op: "Gt"}
	case Lte:
		return BinaryOp{Op: "Lte"}
	case Gte:
		return BinaryOp{Op: "Gte"}
	case And:
		return BinaryOp{Op: "And"}
	case Or:
		return BinaryOp{Op: "Or"}
	default:
		return BinaryOp{Op: "Unknow binary op"}
	}
}

type Bool struct {
	Kind     string   `json:"kind"`
	Value    bool     `json:"value"`
	Location Location `json:"location"`
}

func NewBool(value bool, loc Location) Bool {
	return Bool{
		Kind:     "Bool",
		Value:    value,
		Location: loc,
	}
}

type If struct {
	Kind      string   `json:"kind"`
	Condition Term     `json:"condition"`
	Then      Term     `json:"then"`
	Otherwise Term     `json:"otherwise"`
	Location  Location `json:"location"`
}

type Binary struct {
	Kind     string   `json:"kind"`
	Lhs      Term     `json:"lhs"`
	Op       BinaryOp `json:"op"`
	Rhs      Term     `json:"rhs"`
	Location Location `json:"location"`
}
type Tuple struct {
	Kind     string   `json:"kind"`
	First    Term     `json:"first"`
	Second   Term     `json:"second"`
	Location Location `json:"location"`
}
type First struct {
	Kind     string   `json:"kind"`
	Value    bool     `json:"value"`
	Location Location `json:"location"`
}
type Second struct {
	Kind     string   `json:"kind"`
	Value    bool     `json:"value"`
	Location Location `json:"location"`
}
type Print struct {
	Kind     string   `json:"kind"`
	Value    Term     `json:"value"`
	Location Location `json:"location"`
}

func NewPrint(kind string, value Term, loc Location) Print {
	return Print{
		Kind:     kind,
		Value:    value,
		Location: loc,
	}
}

type Term interface{}

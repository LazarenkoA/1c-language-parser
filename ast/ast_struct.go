package ast

type StatementType int
type OperationType int

const (
	Undefined StatementType = iota
	NodeTypeProcedure
	NodeTypeFunction
)

const (
	OpPlus OperationType = iota
	OpMinus
	OpMul
	OpDiv
	// OpAssig // = (присвоение)
	OpEq  // =
	OpGt  // >
	OpLt  // <
	OpNe  // <>
	OpLe  // <=
	OpGe  // >=
	OpMod // % - деление по модулю
	OpOr
	OpAnd
)

type IUnary interface {
	Unary() interface{}
}

type INot interface {
	Not() interface{}
}

type Statement interface{}

type ModuleStatement struct {
	Name string
	Body []Statement
}

type VarStatement struct {
	Name  string
	unary bool
	not   bool
}

type FunctionOrProcedure struct {
	Type              StatementType
	Name              string
	Body              []Statement
	Export            bool
	Params            []ParamStatement
	Directive         string
	ExplicitVariables map[string]VarStatement
	// ImplicitVariables []VarStatement
}

type ParamStatement struct {
	Name    string
	IsValue bool      `json:"IsValue,omitempty"`
	Default Statement `json:"Default,omitempty"`
}

type ExpStatement struct {
	Operation OperationType
	Left      interface{}
	Right     interface{}
	unary     bool // для сложных выражений, таких как -(1+1)
	not       bool
}

type IfStatement struct {
	Expression  Statement
	TrueBlock   []Statement
	IfElseBlock []*IfStatement
	ElseBlock   []Statement
}

type TryStatement struct {
	Body  []Statement
	Catch []Statement
}

type ThrowStatement struct {
	Param Statement
}

type UndefinedStatement struct{}

type ReturnStatement struct {
	Param Statement
}

type NewObjectStatement struct {
	Constructor string
	Param       []Statement
}

type CallChainStatement struct {
	Unit Statement
	Call Statement
}

type MethodStatement struct {
	Name  string
	Param []Statement
}

type BreakStatement struct {
}

type ContinueStatement struct {
}

type LoopStatement struct {
	Body []Statement
	For  Statement
	To   Statement
	In   Statement
}

type TernaryStatement struct {
	Expression Statement
	TrueBlock  Statement
	ElseBlock  Statement
}

type ItemStatement struct {
	Item   Statement
	Object Statement
}

func (p *ParamStatement) Fill(valueParam *Token, identifier Token) *ParamStatement {
	p.IsValue = valueParam != nil
	p.Name = identifier.literal
	return p
}

func (p *ParamStatement) DefaultValue(value Statement) *ParamStatement {
	if value == nil {
		p.Default = UndefinedStatement{}
	} else {
		p.Default = value
	}

	return p
}

func (e ExpStatement) Unary() interface{} {
	e.unary = true
	return e
}

func (e VarStatement) Unary() interface{} {
	e.unary = true
	return e
}

func (e *ExpStatement) Not() interface{} {
	e.not = true
	return e
}

func (e *VarStatement) Not() interface{} {
	e.not = true
	return e
}

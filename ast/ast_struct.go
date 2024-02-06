package ast

type StatementType int
type OperationType int

const (
	PFTypeUndefined StatementType = iota
	PFTypeProcedure
	PFTypeFunction
)

const (
	OpPlus OperationType = iota
	OpMinus
	OpMul
	OpDiv
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

// type Statements []Statement

type ModuleStatement struct {
	Name string
	Body []Statement
}

type VarStatement struct {
	addStatementField
	Name string
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

type addStatementField struct {
	unary bool
	not   bool
}

type ExpStatement struct {
	addStatementField

	Operation OperationType
	Left      interface{}
	Right     interface{}
}

// type IfElseStatement struct {
// 	Expression Statement
// 	TrueBlock  []Statement
// }

type IfStatement struct {
	Expression  Statement
	TrueBlock   []Statement
	IfElseBlock []Statement
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
	addStatementField

	Unit Statement
	Call Statement
}

type MethodStatement struct {
	addStatementField

	Name  string
	Param []Statement
}

type BreakStatement struct {
}

type ContinueStatement struct {
}

type LoopStatement struct {
	Body      []Statement
	For       Statement `json:"For,omitempty"`
	To        Statement `json:"To,omitempty"`
	In        Statement `json:"In,omitempty"`
	WhileExpr Statement `json:"WhileExpr,omitempty"`
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

func (e *ExpStatement) Unary() interface{} {
	e.unary = true
	return e
}

func (e *ExpStatement) Not() interface{} {
	e.not = true
	return e
}

func (e VarStatement) Unary() interface{} {
	e.unary = true
	return e
}

func (e VarStatement) Not() interface{} {
	e.not = true
	return e
}

func (e CallChainStatement) Unary() interface{} {
	e.unary = true
	return e
}

func (e CallChainStatement) Not() interface{} {
	e.not = true
	return e
}

func (e MethodStatement) Not() interface{} {
	e.not = true
	return e
}

func (o OperationType) String() string {
	switch o {
	case OpPlus:
		return "+"
	case OpMinus:
		return "-"
	case OpMul:
		return "*"
	case OpDiv:
		return "/"
	case OpEq:
		return "="
	case OpGt:
		return ">"
	case OpLt:
		return "<"
	case OpNe:
		return "<>"
	case OpLe:
		return "<="
	case OpGe:
		return ">="
	case OpMod:
		return "%"
	case OpOr:
		return "ИЛИ"
	case OpAnd:
		return "И"
	default:
		return ""
	}
}

func (m ModuleStatement) Walk(callBack func(statement *Statement)) {
	walkHelper(m.Body, callBack)
}

// func (m Statements) Walk(callBack func(statement *Statement)) {
// 	walkHelper(m, callBack)
// }

func walkHelper(statements []Statement, callBack func(statement *Statement)) {
	for _, item := range statements {
		switch v := item.(type) {
		case *IfStatement:
			walkHelper(v.TrueBlock, callBack)
			walkHelper(v.ElseBlock, callBack)
			walkHelper(v.IfElseBlock, callBack)
		case TryStatement:
			walkHelper(v.Body, callBack)
			walkHelper(v.Catch, callBack)
		case *LoopStatement:
			walkHelper(v.Body, callBack)
		case *FunctionOrProcedure:
			walkHelper(v.Body, callBack)
		}

		callBack(&item)
	}

}

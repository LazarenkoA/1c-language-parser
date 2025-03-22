package ast

import "fmt"

type StatementType int
type OperationType int

const (
	PFTypeUndefined StatementType = iota
	PFTypeProcedure
	PFTypeFunction
)

const (
	OpUndefined OperationType = iota
	OpPlus
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
	UnaryMinus() interface{}
}

type INot interface {
	Not() interface{}
}

type IParams interface {
	Params() []Statement
}

type Statement interface{}

type GlobalVariables struct {
	Directive string
	Var       VarStatement
	Export    bool
}

type ModuleStatement struct {
	Name            string
	GlobalVariables map[string]GlobalVariables `json:"GlobalVariables,omitempty"`
	Body            []Statement
}

type VarStatement struct {
	Name string
	addStatementField
}

type FunctionOrProcedure struct {
	ExplicitVariables map[string]VarStatement
	Name              string
	Directive         string
	Body              []Statement
	Params            []ParamStatement
	Type              StatementType
	Export            bool
}

type ParamStatement struct {
	Default Statement `json:"Default,omitempty"`
	Name    string
	IsValue bool `json:"IsValue,omitempty"`
}

type addStatementField struct {
	unaryMinus bool
	unaryPlus  bool
	not        bool
}

type ExpStatement struct {
	Left      interface{}
	Right     interface{}
	Operation OperationType
	addStatementField
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
	Unit Statement
	Call Statement
	addStatementField
}

type MethodStatement struct {
	Name  string
	Param []Statement
	addStatementField
}

type BreakStatement struct {
}

type ContinueStatement struct {
}

type LoopStatement struct {
	For       Statement `json:"For,omitempty"`
	To        Statement `json:"To,omitempty"`
	In        Statement `json:"In,omitempty"`
	WhileExpr Statement `json:"WhileExpr,omitempty"`
	Body      []Statement
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

type GoToStatement struct {
	Label *GoToLabelStatement
}

type GoToLabelStatement struct {
	Name string
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

func (e *ExpStatement) UnaryMinus() interface{} {
	e.unaryMinus = true
	return e
}

func (e *ExpStatement) Not() interface{} {
	e.not = true
	return e
}

func (e VarStatement) UnaryMinus() interface{} {
	e.unaryMinus = true
	return e
}

func (e VarStatement) Not() interface{} {
	e.not = true
	return e
}

func (e CallChainStatement) UnaryMinus() interface{} {
	e.unaryMinus = true
	return e
}

func (e CallChainStatement) Not() interface{} {
	e.not = true
	return e
}

// IsMethod вернет true в случаях Блокировка.Заблокировать() и false для Источник.Ссылка
func (e CallChainStatement) IsMethod() bool {
	_, ok := e.Unit.(MethodStatement)
	return ok
}

func (n MethodStatement) Not() interface{} {
	n.not = true
	return n
}

func (n NewObjectStatement) Params() []Statement {
	return n.Param
}

func (n MethodStatement) Params() []Statement {
	return n.Param
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

func (m ModuleStatement) Walk(callBack func(current *FunctionOrProcedure, statement *Statement)) {
	StatementWalk(m.Body, callBack)
}

func StatementWalk(stm []Statement, callBack func(current *FunctionOrProcedure, statement *Statement)) {
	walkHelper(nil, stm, callBack)
}

func (m *ModuleStatement) Append(item Statement, yylex yyLexer) {
	switch v := item.(type) {
	case GlobalVariables:
		if len(m.Body) > 0 {
			yylex.Error("variable declarations must be placed at the beginning of the module")
			return
		}

		if m.GlobalVariables == nil {
			m.GlobalVariables = map[string]GlobalVariables{}
		}

		if _, ok := m.GlobalVariables[v.Var.Name]; ok {
			yylex.Error(fmt.Sprintf("%v: with the specified name %q", errVariableAlreadyDefined, v.Var.Name))
		} else {
			m.GlobalVariables[v.Var.Name] = v
		}
	case []GlobalVariables:
		for _, item := range v {
			m.Append(item, yylex)
		}
	case []Statement:
		m.Body = append(m.Body, v...)
	case *FunctionOrProcedure:
		// если предыдущее выражение не процедура функция, то это значит что какой-то умник вначале или в середине модуля вставил какие-то выражения, а это нельзя. 1С разрешает выражения только в конце модуля
		if len(m.Body) > 0 {
			if _, ok := m.Body[len(m.Body)-1].(*FunctionOrProcedure); !ok {
				yylex.Error("procedure and function definitions should be placed before the module body statements")
				return
			}
		}

		m.Body = append(m.Body, item)
	default:
		m.Body = append(m.Body, item)
	}
}

// func (m Statements) Walk(callBack func(statement *Statement)) {
// 	walkHelper(m, callBack)
// }

func walkHelper(parent *FunctionOrProcedure, statements []Statement, callBack func(current *FunctionOrProcedure, statement *Statement)) {
	for i, item := range statements {
		switch v := item.(type) {
		case *IfStatement:
			walkHelper(parent, []Statement{v.Expression}, callBack)
			walkHelper(parent, v.TrueBlock, callBack)
			walkHelper(parent, v.ElseBlock, callBack)
			walkHelper(parent, v.IfElseBlock, callBack)
		case TryStatement:
			walkHelper(parent, v.Body, callBack)
			walkHelper(parent, v.Catch, callBack)
		case *LoopStatement:
			walkHelper(parent, v.Body, callBack)
		case *FunctionOrProcedure:
			walkHelper(v, v.Body, callBack)
			parent = v
		case MethodStatement:
			walkHelper(parent, v.Param, callBack)
		//case CallChainStatement:
		//	walkHelper(parent, []Statement{v.Unit}, callBack)
		case *ExpStatement:
			walkHelper(parent, []Statement{v.Right}, callBack)
			walkHelper(parent, []Statement{v.Left}, callBack)
		case TernaryStatement:
			walkHelper(parent, []Statement{v.Expression}, callBack)
			walkHelper(parent, []Statement{v.TrueBlock}, callBack)
			walkHelper(parent, []Statement{v.ElseBlock}, callBack)
		case *ReturnStatement:
			walkHelper(parent, []Statement{v.Param}, callBack)
		}

		callBack(parent, &statements[i])
	}
}

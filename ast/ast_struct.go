package ast

import "fmt"

type StatementType int
type OperationType int
type fCallBack func(root *FunctionOrProcedure, parentStm, stm *Statement)

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
	Params() Statements
}

type Statement interface{}
type Statements []Statement

type AssignmentStatement struct {
	Var  Statement
	Expr ExprStatements
}

type ExprStatements struct {
	Statements
	addStatementField
}

type GlobalVariables struct {
	Directive string
	Var       VarStatement
	Export    bool
}

type ModuleStatement struct {
	Name            string
	GlobalVariables map[string]GlobalVariables `json:"GlobalVariables,omitempty"`
	Body            Statements
}

type VarStatement struct {
	Name string
	addStatementField
}

type FunctionOrProcedure struct {
	ExplicitVariables map[string]VarStatement
	Name              string
	Directive         string
	Body              Statements
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
// 	TrueBlock  Statements
// }

type IfStatement struct {
	Expression  Statement
	TrueBlock   Statements
	IfElseBlock Statements
	ElseBlock   Statements
}

type TryStatement struct {
	Body  Statements
	Catch Statements
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
	Param       ExprStatements
}

type CallChainStatement struct {
	Unit Statement
	Call Statement
	addStatementField
}

type MethodStatement struct {
	Name  string
	Param ExprStatements
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
	Body      Statements
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

func (e ExprStatements) Not() interface{} {
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

func (n NewObjectStatement) Params() ExprStatements {
	return n.Param
}

func (n MethodStatement) Params() ExprStatements {
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

func (m *ModuleStatement) Walk(callBack fCallBack) {
	StatementWalk(m.Body, m.Body, callBack)
}

func StatementWalk(parentStm Statement, stm Statements, callBack fCallBack) {
	walkHelper(nil, parentStm, stm, callBack)
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
	case Statements:
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

func walkHelper(parent *FunctionOrProcedure, parentStm Statement, statements Statements, callBack fCallBack) {
	for i, item := range statements {
		switch v := item.(type) {
		case *IfStatement:
			walkHelper(parent, v, Statements{v.Expression}, callBack)
			walkHelper(parent, v, v.TrueBlock, callBack)
			walkHelper(parent, v, v.ElseBlock, callBack)
			walkHelper(parent, v, v.IfElseBlock, callBack)
		case TryStatement:
			walkHelper(parent, v, v.Body, callBack)
			walkHelper(parent, v, v.Catch, callBack)
		case *LoopStatement:
			walkHelper(parent, v, v.Body, callBack)
		case *FunctionOrProcedure:
			walkHelper(v, v, v.Body, callBack)
			parent = v
		case MethodStatement:
			walkHelper(parent, v, v.Param.Statements, callBack)
		//case CallChainStatement:
		//	walkHelper(parent, Statements{v.Unit}, callBack)
		case *ExpStatement:
			walkHelper(parent, v, Statements{v.Right}, callBack)
			walkHelper(parent, v, Statements{v.Left}, callBack)
		case TernaryStatement:
			walkHelper(parent, v, Statements{v.Expression}, callBack)
			walkHelper(parent, v, Statements{v.TrueBlock}, callBack)
			walkHelper(parent, v, Statements{v.ElseBlock}, callBack)
		case *ReturnStatement:
			walkHelper(parent, v, Statements{v.Param}, callBack)
		}

		callBack(parent, &parentStm, &statements[i])
	}
}
